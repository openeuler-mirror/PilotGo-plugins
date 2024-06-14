package conf

type TemplateConf struct {
	Https_enabled bool   `yaml:"https_enabled"`
	CertFile      string `yaml:"cert_file"`
	KeyFile       string `yaml:"key_file"`
	Addr          string `yaml:"addr"`
}

type PilotGoConf struct {
	Addr string `yaml:"addr"`
}
