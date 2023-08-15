package main

import (
	"fmt"
	"os"

	"gitee.com/openeuler/PilotGo-plugin-topology-server/clientmanager"
	"gitee.com/openeuler/PilotGo-plugin-topology-server/conf"
	"gitee.com/openeuler/PilotGo-plugin-topology-server/handler"
	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"gitee.com/openeuler/PilotGo-plugins/sdk/plugin/client"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello topology")

	InitLogger()

	// TODO: init arangodb

	PluginClient := client.DefaultClient(clientmanager.PluginInfo)
	// 临时给server赋值
	PluginClient.Server = "http://192.168.75.100:8887"
	clientmanager.Galaops = &clientmanager.Opsclient{
		Sdkmethod: PluginClient,
	}

	// 设置router
	engine := gin.Default()
	clientmanager.Galaops.Sdkmethod.RegisterHandlers(engine)
	handler.InitRouter(engine)
	if err := engine.Run(conf.Config().Http.Server_addr); err != nil {
		logger.Fatal("failed to run server")
	}
}

func InitLogger() {
	if err := logger.Init(conf.Config().Logopts); err != nil {
		fmt.Printf("logger init failed, please check the config file: %s", err)
		os.Exit(1)
	}
}
