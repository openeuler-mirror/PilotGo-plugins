package config

import (
	"ant-agent/cmd/config/options"

	"gitee.com/openeuler/PilotGo/sdk/logger"
)

func NewHttpServerOptions() *options.HttpServer {
	return &options.HttpServer{
		Addr: "localhost:8277",
	}
}
func NewLogOptions() *logger.LogOpts {
	return &logger.LogOpts{
		Level:   "debug",
		Driver:  "file",
		Path:    "./log/ant-agent.log",
		MaxFile: 1,
		MaxSize: 10485760,
	}
}

var DefaultConfigTemplate = options.ServerConfig{
	HttpServer: NewHttpServerOptions(),
	Logopts:    NewLogOptions(),
}
