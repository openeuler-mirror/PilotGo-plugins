package main

import (
	"github.com/gomodule/redigo/redis"
	"strconv"
	"strings"
)

type RedisService struct {
	conn redis.Conn
}

func connectToRedis() (*RedisService, error) {
	c, err := redis.DialURL("redis://localhost:6379")
	if err != nil {
		return nil, err
	}
	return &RedisService{conn: c}, nil
}

func (r *RedisService) doRedisCmd(cmd string, args ...interface{}) (interface{}, error) {
	res, err := r.conn.Do(cmd, args...)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (r *RedisService) get() (map[string]string, []CommandStats, map[string]string, map[string]string, map[string]string, map[string]string, error) {
	servers, comm, cpu, memory, stats, cluster, err := r.getInfo()
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	return servers, r.getCommandstats(comm), cpu, memory, stats, cluster, nil
}

type CommandStats struct {
	Comm        string  `json:"comm"`
	Calls       float64 `json:"calls"`
	Usec        float64 `json:"usec"`
	UsecPerCall float64 `json:"usec_per_call"`
}

func (r *RedisService) getCommandstats(res map[string]string) []CommandStats {
	var p []CommandStats
	for k, v := range res {
		q := CommandStats{}
		q.Comm = strings.Split(k, "_")[1]
		for _, m := range strings.Split(v, ",") {
			d := strings.Split(m, "=")
			if d[0] == "calls" {
				q.Calls, _ = strconv.ParseFloat(d[1], 64)
			} else if d[0] == "usec" {
				q.Usec, _ = strconv.ParseFloat(d[1], 64)
			} else if d[0] == "usec_per_call" {
				q.UsecPerCall, _ = strconv.ParseFloat(d[1], 64)
			}
		}
		p = append(p, q)
	}
	return p
}

func (r *RedisService) getInfo() (map[string]string, map[string]string, map[string]string, map[string]string, map[string]string, map[string]string, error) {
	infoAll, err := redis.String(r.doRedisCmd("INFO", "ALL"))
	if err != nil || infoAll == "" {

		infoAll, err = redis.String(r.doRedisCmd("INFO"))
		if err != nil {
			return nil, nil, nil, nil, nil, nil, err
		}
	}
	servers := map[string]string{}
	comm := map[string]string{}
	cpu := map[string]string{}
	memory := map[string]string{}
	stats := map[string]string{}
	cluster := map[string]string{}
	lines := strings.Split(infoAll, "\n")
	fieldClass := ""
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) > 0 && strings.HasPrefix(line, "# ") {
			fieldClass = line[2:]
			continue
		}
		if len(line) < 2 || (!strings.Contains(line, ":")) {
			continue
		}
		split := strings.SplitN(line, ":", 2)
		fieldKey := split[0]
		fieldValue := split[1]
		switch fieldClass {
		case "Stats":
			stats[fieldKey] = fieldValue
		case "Server":
			servers[fieldKey] = fieldValue
		case "Cluster":
			cluster[fieldKey] = fieldValue
		case "Commandstats":
			comm[fieldKey] = fieldValue
		case "Memory":
			memory[fieldKey] = fieldValue
		case "CPU":
			cpu[fieldKey] = fieldValue
		}
	}
	return servers, comm, cpu, memory, stats, cluster, nil
}
