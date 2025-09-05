package service

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/cmd/config/options"
)

type RedisService struct {
	Conf   *options.RedisDBInfo
	client *redis.Client
	cancel context.CancelFunc
	ctx    context.Context
}

func (r *RedisService) Name() string {
	return "Redis"
}

func (r *RedisService) Init(ctx *AppContext) error {
	if r.Conf.UseTLS {
		switch r.Conf.Mode {
		case "standalone":
			redCli := redis.NewClient(&redis.Options{
				Addr:     r.Conf.Host,
				Password: r.Conf.Password,
				DB:       r.Conf.DefaultDB,
				TLSConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			})
			r.client = redCli
		case "sentinel":
			redCli := redis.NewFailoverClient(&redis.FailoverOptions{
				MasterName:    r.Conf.MasterName,
				SentinelAddrs: strings.Split(r.Conf.SentinelHosts, ","),
				Password:      r.Conf.Password,
				DB:            r.Conf.DefaultDB,
				TLSConfig:     &tls.Config{InsecureSkipVerify: true},
			})
			r.client = redCli
		default:
			return fmt.Errorf("redis mode only support standalone or sentinel")
		}
	} else {
		switch r.Conf.Mode {
		case "standalone":
			redCli := redis.NewClient(&redis.Options{
				Addr:     r.Conf.Host,
				Password: r.Conf.Password,
				DB:       r.Conf.DefaultDB,
			})
			r.client = redCli
		case "sentinel":
			redCli := redis.NewFailoverClient(&redis.FailoverOptions{
				MasterName:    r.Conf.MasterName,
				SentinelAddrs: strings.Split(r.Conf.SentinelHosts, ","),
				Password:      r.Conf.Password,
				DB:            r.Conf.DefaultDB,
			})
			r.client = redCli
		default:
			return fmt.Errorf("redis mode only support standalone or sentinel")
		}
	}

	baseCtx, cancel := context.WithCancel(context.Background())
	r.ctx = baseCtx
	r.cancel = cancel

	// 验证连接
	timeoutCtx, timeoutCancel := context.WithTimeout(baseCtx, r.Conf.DialTimeout)
	defer timeoutCancel()

	if _, err := r.client.Ping(timeoutCtx).Result(); err != nil {
		return err
	}

	ctx.Redis = r
	return nil
}

func (r *RedisService) Close() error {
	if r.cancel != nil {
		r.cancel()
	}
	if r.client != nil {
		if err := r.client.Close(); err != nil {
			return err
		}
	}
	return nil
}

// ===============================Redis API=============================================
type Redis interface {
	SetNX(key string, value interface{}, expiration time.Duration) error
	Get(key string, out interface{}) error
	Delete(keys ...string) error

	ReleaseLock(key string, value interface{}) (bool, error)
	AcquireLockWithRetry(key string, value interface{}, expiration, retryInterval, timeout time.Duration) error

	Publish(channel string, message interface{}) error
	Subscribe(channel string, handler func(message *redis.Message)) error
	PSubscribe(channel string, handler func(message *redis.Message)) error
}

// =============================Redis manager============================================
func (r *RedisService) SetNX(key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.client.SetNX(r.ctx, key, data, expiration).Err()
}
func (r *RedisService) Get(key string, dest interface{}) error {
	data, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(data), dest)
}
func (r *RedisService) Delete(keys ...string) error {
	return r.client.Del(r.ctx, keys...).Err()
}

// AcquireLock 尝试获取分布式锁
func (r *RedisService) acquireLock(key string, value interface{}, expiration time.Duration) error {
	return r.SetNX(key, value, expiration)
}

// ReleaseLock 释放锁（仅释放自己加的锁）
func (r *RedisService) ReleaseLock(key string, value interface{}) (bool, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return false, err
	}

	script := redis.NewScript(`
		if redis.call("get", KEYS[1]) == ARGV[1] then
			return redis.call("del", KEYS[1])
		else
			return 0
		end
	`)
	res, err := script.Run(r.ctx, r.client, []string{key}, data).Int()
	return res == 1, err
}

// AcquireLockWithRetry 尝试获取分布式锁，失败自动重试直到超时
func (r *RedisService) AcquireLockWithRetry(key string, value interface{}, expiration, retryInterval, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for {
		err := r.acquireLock(key, value, expiration)
		if err != nil {
			return err
		}

		if time.Now().After(deadline) {
			return errors.New("AcquireLockWithRetry timeout: failed to acquire lock")
		}

		select {
		case <-r.ctx.Done():
			return r.ctx.Err()
		case <-time.After(retryInterval):
		}
	}
}
func (r *RedisService) Publish(channel string, message interface{}) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return r.client.Publish(r.ctx, channel, data).Err()
}

func (r *RedisService) Subscribe(channel string, handler func(message *redis.Message)) error {
	pubsub := r.client.Subscribe(r.ctx, channel)
	_, err := pubsub.Receive(r.ctx)
	if err != nil {
		return err
	}

	ch := pubsub.Channel()
	go func() {
		for {
			select {
			case msg, ok := <-ch:
				if !ok {
					return
				}
				handler(msg)
			case <-r.ctx.Done():
				_ = pubsub.Close()
				return
			}
		}
	}()
	return nil
}
func (r *RedisService) PSubscribe(channel string, handler func(message *redis.Message)) error {
	pubsub := r.client.PSubscribe(r.ctx, channel)
	_, err := pubsub.Receive(r.ctx)
	if err != nil {
		return err
	}

	ch := pubsub.Channel()
	go func() {
		for {
			select {
			case msg, ok := <-ch:
				if !ok {
					return
				}
				handler(msg)
			case <-r.ctx.Done():
				_ = pubsub.Close()
				return
			}
		}
	}()
	return nil
}
