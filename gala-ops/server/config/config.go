package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type GrafanaConf struct {
	Addr   string `yaml:"addr"`
	User   string `yaml:"user"`
	Passwd string `yaml:"passwd"`
}

type HttpConf struct {
	Addr string `yaml:"addr"`
}

type PilotGoConf struct {
	Addr string `yaml:"addr"`
}

type ServerConfig struct {
	Grafana *GrafanaConf `yaml:"grafana"`
	Http    *HttpConf    `yaml:"http"`
	PilotGo *PilotGoConf `yaml:"PilotGo"`
}

const config_file = "./config.yaml"

var global_config ServerConfig

func Init() {
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
	bytes, err := ioutil.ReadFile(file)
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
