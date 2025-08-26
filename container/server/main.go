package main

import (
	"fmt"
	"os"

	"gitee.com/openeuler/PilotGo/sdk/go-micro/registry"
	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-event/service"
	cli "openeuler.org/PilotGo/container-plugin/client"
	"openeuler.org/PilotGo/container-plugin/config"
	"openeuler.org/PilotGo/container-plugin/database"
	"openeuler.org/PilotGo/container-plugin/httphandler"
)

func main() {
	fmt.Println("hello gala-ops")

	if err := database.MysqlInit(config.Config().Mysql); err != nil {
		fmt.Println("failed to initialize database")
		os.Exit(-1)
	}

	InitLogger()

	server := gin.Default()

	sr, err := registry.NewServiceRegistrar(&registry.Options{
		Endpoints:   config.Config().Etcd.Endpoints,
		ServiceAddr: config.Config().Http.Addr,
		ServiceName: config.Config().Etcd.ServiveName,
		Version:     config.Config().Etcd.Version,
		MenuName:    config.Config().Etcd.MenuName,
		Icon:        config.Config().Etcd.Icon,
		DialTimeout: config.Config().Etcd.DialTimeout,
		Extentions:  service.GetExtentions(),
		Permissions: service.GetPermissions(),
	})
	if err != nil {
		logger.Error("failed to initialize registry: %s", err)
		os.Exit(-1)
	}

	client, err := client.NewClient(config.Config().Etcd.ServiveName, sr.Registry)
	if err != nil {
		logger.Error("failed to create plugin client: %s", err)
		os.Exit(-1)
	}

	cli.ContainerClient = client
	cli.ContainerClient.RegisterHandlers(server)
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
