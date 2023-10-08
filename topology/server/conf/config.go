package conf

import (
	"path"

	"gitee.com/openeuler/PilotGo/sdk/logger"
)

type TopoConf struct {
	Server_addr string `yaml:"server_addr"`
	Agent_port  string `yaml:"agent_port"`
	Period      int    `yaml:"period"`
}

type PilotGoConf struct {
	Addr string `yaml:"http_addr"`
}

type ArangodbConf struct {
	Addr     string `yaml:"addr"`
	Database string `yaml:"database"`
}

type ServerConfig struct {
	Topo     *TopoConf       `yaml:"topo"`
	PilotGo  *PilotGoConf    `yaml:"PilotGo"`
	Logopts  *logger.LogOpts `yaml:"log"`
	Arangodb *ArangodbConf   `yaml:"arangodb"`
}

const config_type = "config_server.yaml"

var Config_dir string

func Config_file() string {
	// _, thisfilepath, _, _ := runtime.Caller(0)
	// dirpath := filepath.Dir(thisfilepath)
	// configfilepath := path.Join(dirpath, "..", "..", "conf", config_type)

	// ttcode:
	configfilepath := path.Join(Config_dir, config_type)

	return configfilepath
}

var Global_config ServerConfig

func Config() *ServerConfig {
	return &Global_config
}
