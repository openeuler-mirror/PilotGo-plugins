/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: wubijie <wubijie@kylinos.cn>
 * Date: Wed Nov 15 14:46:48 2023 +0800
 */
package config

import (
	"fmt"
	"os"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gopkg.in/yaml.v2"
)

type ConfigPlugin struct {
	URL string `yaml:"url"`
}

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

type ServerConfig struct {
	ConfigPlugin *ConfigPlugin   `yaml:"config_plugin"`
	HttpServer   *HttpServer     `yaml:"http_server"`
	Logopts      *logger.LogOpts `yaml:"log"`
	Mysql        *MysqlDBInfo    `yaml:"mysql"`
}

var global_config ServerConfig

func Init(config_file string) error {
	err := readConfig(config_file, &global_config)
	return err
}

func Config() *ServerConfig {
	return &global_config
}

func readConfig(file string, config interface{}) error {
	bytes, err := os.ReadFile(file)
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
