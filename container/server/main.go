package main

import (
	"fmt"
	"os"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/container-plugin/config"
	"openeuler.org/PilotGo/container-plugin/database"
	"openeuler.org/PilotGo/container-plugin/httphandler"
)

const Version = "0.0.1"

var PluginInfo = &client.PluginInfo{
	Name:        "container",
	Version:     Version,
	Description: "Container management plugin",
	Author:      "wangjunqi",
	Email:       "wangjunqi@kylinos.cn",
	Url:         "http://192.168.75.100:9998/plugin/container",
	// ReverseDest: "",
}

func main() {
	fmt.Println("hello gala-ops")

	if err := database.MysqlInit(config.Config().Mysql); err != nil {
		fmt.Println("failed to initialize database")
		os.Exit(-1)
	}

	InitLogger()

	server := gin.Default()

	GlobalClient := client.DefaultClient(PluginInfo)
	// 临时给server赋值
	GlobalClient.Server = "http://192.168.75.100:8888"
	GlobalClient.RegisterHandlers(server)
	InitRouter(server)

	if err := server.Run(config.Config().Http.Addr); err != nil {
		logger.Fatal("failed to run server")
	}
}

func InitLogger() {
	if err := logger.Init(config.Config().Logopts); err != nil {
		fmt.Printf("logger init failed, please check the config file: %s", err)
		os.Exit(-1)
	}
}

func InitRouter(router *gin.Engine) {
	api := router.Group("/plugin/container/api")
	{
		api.POST("/deploy_docker", httphandler.DeployDocker)
	}
}
