package conf

import (
	"flag"
	"fmt"
	"os"
	"path"

	"gitee.com/openeuler/PilotGo-plugin-template/global"
	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gopkg.in/yaml.v2"
)

var Global_Config *ServerConfig

const config_type = "template.yaml"

var config_dir string

type ServerConfig struct {
	Template *TemplateConf
	PilotGo  *PilotGoConf
	Logopts  *logger.LogOpts `yaml:"log"`
}

func ConfigFile() string {
	configfilepath := path.Join(config_dir, config_type)

	return configfilepath
}

func InitConfig() {
	flag.StringVar(&config_dir, "conf", "/opt/PilotGo/plugin/template", "template configuration directory")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s -conf /path/to/template.yaml(default:/opt/PilotGo/plugin/template) \n", os.Args[0])
	}
	flag.Parse()

	bytes, err := global.FileReadBytes(ConfigFile())
	if err != nil {
		flag.Usage()
		fmt.Printf("open file failed: %s, %s", ConfigFile(), err.Error())
		os.Exit(1)
	}

	Global_Config = &ServerConfig{}

	err = yaml.Unmarshal(bytes, Global_Config)
	if err != nil {
		fmt.Printf("yaml unmarshal failed: %s", err.Error())
		os.Exit(1)
	}
}
