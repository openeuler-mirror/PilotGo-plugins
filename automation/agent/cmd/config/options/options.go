package options

import (
	"fmt"
	"os"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gopkg.in/yaml.v2"
)

type HttpServer struct {
	Addr string `yaml:"addr"`
}

type ServerConfig struct {
	HttpServer *HttpServer     `yaml:"http_server"`
	Logopts    *logger.LogOpts `yaml:"log"`
}

const config_file = "./ant-agent.yaml"

type Options struct {
	ConfigFile string
	Config     *ServerConfig
}

func NewOptions() *Options {
	return &Options{
		ConfigFile: config_file,
		Config:     &ServerConfig{},
	}
}

func (O *Options) TryLoadFromDisk() (*Options, error) {
	o, err := loadConfigFromFile(O.ConfigFile)
	if err != nil {
		return nil, err
	}
	O.Config = o
	return O, nil
}

func loadConfigFromFile(path string) (*ServerConfig, error) {
	cfg := &ServerConfig{}
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("open %s failed! err = %s\n", path, err.Error())
		return nil, err
	}

	err = yaml.Unmarshal(bytes, cfg)
	if err != nil {
		fmt.Printf("yaml Unmarshal %s failed!\n", string(bytes))
		return nil, err
	}
	return cfg, nil
}
