package options

import (
	"fmt"
	"os"
	"time"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gopkg.in/yaml.v2"
)

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

type Etcd struct {
	Endpoints   []string      `yaml:"endpoints"`
	ServiceName string        `yaml:"service_name"`
	Version     string        `yaml:"version"`
	DialTimeout time.Duration `yaml:"dialTimeout"`
	MenuName    string        `yaml:"menu_name"`
	Icon        string        `yaml:"icon"`
}

type ServerConfig struct {
	HttpServer *HttpServer     `yaml:"http_server"`
	Logopts    *logger.LogOpts `yaml:"log"`
	Mysql      *MysqlDBInfo    `yaml:"mysql"`
	Etcd       *Etcd           `yaml:"etcd" mapstructure:"etcd"`
}

const config_file = "./automation.yaml"

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
