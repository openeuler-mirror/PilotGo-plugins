package conf

type ElkConf struct {
	Addr string `yaml:"http_addr"`
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
