package config

import (
	"time"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/cmd/config/options"
)

func NewHttpServerOptions() *options.HttpServer {
	return &options.HttpServer{
		Addr: "localhost:8288",
	}
}
func NewLogOptions() *logger.LogOpts {
	return &logger.LogOpts{
		Level:   "debug",
		Driver:  "file",
		Path:    "./log/plugin_automation.log",
		MaxFile: 1,
		MaxSize: 10485760,
	}
}

func NewMysqlDBInfoOptions() *options.MysqlDBInfo {
	return &options.MysqlDBInfo{
		HostName: "localhost",
		Port:     3306,
		UserName: "root",
		Password: "Qwer!234578",
		DataBase: "PluginAutomation",
	}
}

func NewRedisDBInfoOptions() *options.RedisDBInfo {
	return &options.RedisDBInfo{
		Mode:          "standalone",
		Host:          "127.0.0.1:6379",
		SentinelHosts: "1.1.1.1:26379,1.1.1.2:26379,1.1.1.3:26379",
		MasterName:    "mymaster",
		Password:      "Qwer!234578",
		DefaultDB:     0,
		DialTimeout:   5 * time.Second,
		EnableRedis:   true,
		UseTLS:        false,
	}
}

func NewEtcdOptions() *options.Etcd {
	return &options.Etcd{
		Endpoints:   []string{"localhost:2379"},
		ServiceName: "automation-service",
		Version:     "3.0",
		DialTimeout: 5 * time.Second,
		MenuName:    "智能运维调度中心",
		Icon:        "el-icon-setting",
	}
}

var DefaultConfigTemplate = options.ServerConfig{
	HttpServer: NewHttpServerOptions(),
	Logopts:    NewLogOptions(),
	Mysql:      NewMysqlDBInfoOptions(),
	Redis:      NewRedisDBInfoOptions(),
	Etcd:       NewEtcdOptions(),
}
