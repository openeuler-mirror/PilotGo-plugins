package config

import (
	"flag"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/utils/config"
)

type PluginEvent struct {
	URL        string `yaml:"url"`
	PluginType string `yaml:"plugin_type"`
}

type HttpServer struct {
	Addr string `yaml:"addr"`
}
type ServerConfig struct {
	PluginEvent *PluginEvent    `yaml:"plugin_event"`
	HttpServer  *HttpServer     `yaml:"http_server"`
	Logopts     *logger.LogOpts `yaml:"log"`
	Influxd     *Influxd        `yaml:"influxd"`
}

type Influxd struct {
	URL          string `yaml:"url"`
	Token        string `yaml:"token"`
	Organization string `yaml:"organization"`
	Bucket       string `yaml:"bucket"`
}

var config_file string
var global_config ServerConfig

func Init() error {
	flag.StringVar(&config_file, "conf", "./config.yaml", "pilotgo-plugin-event configuration file")
	flag.Parse()
	return config.Load(config_file, &global_config)
}

func Config() *ServerConfig {
	return &global_config
}
