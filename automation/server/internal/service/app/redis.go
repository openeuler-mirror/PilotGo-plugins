package app

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v8"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/cmd/config/options"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/global"
)

type RedisService struct {
	Conf   *options.RedisDBInfo
	client *redis.Client
	cancel context.CancelFunc
}

func (r *RedisService) Name() string {
	return "Redis"
}

func (r *RedisService) Init(ctx *global.AppContext) error {
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

	timeoutCtx, cancel := context.WithTimeout(context.Background(), r.Conf.DialTimeout)
	r.cancel = cancel

	// 验证连接
	if _, err := r.client.Ping(timeoutCtx).Result(); err != nil {
		return err
	}

	ctx.Redis = r.client
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
