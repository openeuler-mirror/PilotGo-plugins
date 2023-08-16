package conf

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"gopkg.in/yaml.v3"
)

type TopoConf struct {
	Agent_addr string `yaml:"agent_addr"`
	Datasource string `yaml:"datasource"`
}

type PilotGoConf struct {
	Addr string `yaml:"addr"`
}

type ServerConfig struct {
	Http    *TopoConf       `yaml:"topohttp"`
	PilotGo *PilotGoConf    `yaml:"PilotGo"`
	Logopts *logger.LogOpts `yaml:"log"`
}

const config_type = "config_agent.yaml"

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
