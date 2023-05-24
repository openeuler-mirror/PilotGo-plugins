package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"openeuler.org/PilotGo/container-plugin/client"
	"openeuler.org/PilotGo/container-plugin/config"
	"openeuler.org/PilotGo/container-plugin/database"
	"openeuler.org/PilotGo/container-plugin/httphandler"
)

func main() {
	fmt.Println("hello container")

	config.Init()

	if err := database.MysqlInit(config.Config().Mysql); err != nil {
		fmt.Println("failed to initialize database")
		os.Exit(-1)
	}

	engine := client.Client().HttpEngine
	registerHandlers(engine)
	client.StartClient(config.Config().Http)
}

func registerHandlers(engine *gin.Engine) {
	api := engine.Group("/plugin/container/api")
	{
		api.PUT("/deploy_docker", httphandler.DeployDocker)
	}
}
