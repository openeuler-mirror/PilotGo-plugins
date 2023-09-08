package conf

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
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

func config_file() string {
	_, thisfilepath, _, _ := runtime.Caller(0)
	dirpath := filepath.Dir(thisfilepath)
	configfilepath := path.Join(dirpath, "..", "..", "conf", config_type)
	return configfilepath
}

var global_config ServerConfig

func init() {
	err := readConfig(config_file(), &global_config)
	if err != nil {
		err = errors.Wrap(err, "**2")
		// errors.EROE(err)
		fmt.Printf("%+v\n", err)
		os.Exit(-1)
	}
}

func Config() *ServerConfig {
	return &global_config
}

func readConfig(file string, config interface{}) error {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return errors.Errorf("open file failed: %s, %s**2", file, err.Error())
	}

	err = yaml.Unmarshal(bytes, config)
	if err != nil {
		return errors.Errorf("yaml unmarshal failed: %s**2", string(bytes))
	}
	return nil
}
