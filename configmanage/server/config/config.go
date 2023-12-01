package config

import (
	"fmt"
	"os"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gopkg.in/yaml.v2"
)

type ConfigPlugin struct {
	URL        string `yaml:"url"`
	PluginType string `yaml:"plugin_type"`
}

type HttpServer struct {
	Addr string `yaml:"addr"`
}

type MysqlDBInfo struct {
	HostName string `yaml:"host"`
	Port     int    `yaml:"port"`
	UserName string `yaml:"user"`
	Password string `yaml:"password"`
	DataBase string `yaml:"database"`
}

type ServerConfig struct {
	ConfigPlugin *ConfigPlugin   `yaml:"config_plugin"`
	HttpServer   *HttpServer     `yaml:"http_server"`
	Logopts      *logger.LogOpts `yaml:"log"`
	Mysql        *MysqlDBInfo    `yaml:"mysql"`
}

var global_config ServerConfig

func Init(config_file string) error {
	err := readConfig(config_file, &global_config)
	return err
}

func Config() *ServerConfig {
	return &global_config
}

func readConfig(file string, config interface{}) error {
	bytes, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("open %s failed! err = %s\n", file, err.Error())
		return err
	}

	err = yaml.Unmarshal(bytes, config)
	if err != nil {
		fmt.Printf("yaml Unmarshal %s failed!\n", string(bytes))
		return err
	}
	return nil
}
