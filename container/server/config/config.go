package config

import (
	"fmt"
	"os"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gopkg.in/yaml.v3"
)

type HttpConf struct {
	Addr string `yaml:"addr"`
}

type MysqlConf struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type ServerConfig struct {
	Http    *HttpConf       `yaml:"http"`
	Mysql   *MysqlConf      `yaml:"mysql"`
	Logopts *logger.LogOpts `yaml:"log"`
}

const config_file = "./config.yml"

var global_config ServerConfig

func init() {
	err := readConfig(config_file, &global_config)
	if err != nil {
		fmt.Printf("")
		os.Exit(-1)
	}
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
