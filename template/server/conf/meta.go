package conf

import "time"

type TemplateConf struct {
	Https_enabled bool   `yaml:"https_enabled"`
	CertFile      string `yaml:"cert_file"`
	KeyFile       string `yaml:"key_file"`
	Addr          string `yaml:"addr"`
}

type Etcd struct {
	Endpoints   []string      `yaml:"endpoints"`
	ServiveName string        `yaml:"service_name"`
	Version     string        `yaml:"version"`
	DialTimeout time.Duration `yaml:"dialTimeout"`
	MenuName    string        `yaml:"menu_name"`
	Icon        string        `yaml:"icon"`
}
