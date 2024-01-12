package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
	"openeuler.org/PilotGo/atune-plugin/config"
	"openeuler.org/PilotGo/atune-plugin/db"
	"openeuler.org/PilotGo/atune-plugin/plugin"
	"openeuler.org/PilotGo/atune-plugin/router"
	"openeuler.org/PilotGo/atune-plugin/service"
)

func main() {
	fmt.Println("hello atune")

	config.Init()

	if err := logger.Init(config.Config().Logopts); err != nil {
		fmt.Printf("logger init failed, please check the config file: %s", err)
		os.Exit(-1)
	}

	if err := db.MysqldbInit(config.Config().Mysql); err != nil {
		logger.Error("mysql db init failed, please check again: %s", err)
		os.Exit(-1)
	}

	plugin.GlobalClient = client.DefaultClient(plugin.Init(config.Config().PluginAtune))
	service.GetTags()
	service.AddExtentions()

	err := router.HttpServerInit(config.Config().HttpServer)
	if err != nil {
		logger.Error("http server init failed, error:%v", err)
		os.Exit(-1)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for {
		s := <-c
		switch s {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			logger.Info("signal interrupted: %s", s.String())
			goto EXIT
		default:
			logger.Info("unknown signal: %s", s.String())
		}
	}

EXIT:
	logger.Info("exit system, bye~")
}
