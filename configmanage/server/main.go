package main

import (
	"flag"
	"fmt"
	"os"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
	"openeuler.org/PilotGo/configmanage-plugin/config"
	"openeuler.org/PilotGo/configmanage-plugin/db"
	"openeuler.org/PilotGo/configmanage-plugin/global"
	"openeuler.org/PilotGo/configmanage-plugin/router"
)

var config_file string

func main() {
	fmt.Println("hello plugin-config")

	flag.StringVar(&config_file, "conf", "./config.yaml", "plugin-config configuration file")
	flag.Parse()
	err := config.Init(config_file)
	if err != nil {
		fmt.Println("failed to load configure, exit..", err)
		os.Exit(-1)
	}

	if err := logger.Init(config.Config().Logopts); err != nil {
		fmt.Printf("logger init failed, please check the config file: %s", err)
		os.Exit(-1)
	}
	logger.Info("Thanks to choose PilotGo!")

	// mysql db初始化
	if err := db.MysqldbInit(config.Config().Mysql); err != nil {
		logger.Error("mysql db init failed, please check again: %s", err)
		os.Exit(-1)
	}

	server := router.InitRouter()
	global.GlobalClient = client.DefaultClient(global.Init(config.Config().ConfigPlugin))

	go router.RegisterAPIs(server)
	if err := server.Run(config.Config().HttpServer.Addr); err != nil {
		logger.Error("failed to run server: %s", err)
		os.Exit(-1)
	}
}
