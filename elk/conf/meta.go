package conf

type ElkConf struct {
	Https_enabled      bool   `yaml:"https_enabled"`
	Public_certificate string `yaml:"public_certificate"`
	Private_key        string `yaml:"private_key"`
	Addr               string `yaml:"addr"`
}

type PilotGoConf struct {
	Addr string `yaml:"http_addr"`
}

type ElasticConf struct {
	Addr     string `yaml:"http_addr"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type LogstashConf struct {
	Addr string `yaml:"http_addr"`
}

type KibanaConf struct {
	Addr     string `yaml:"http_addr"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
