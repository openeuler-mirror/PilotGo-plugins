package handler

import (
	"fmt"
	"os"

	"gitee.com/openeuler/PilotGo-plugin-topology-server/agentmanager"
	"gitee.com/openeuler/PilotGo-plugin-topology-server/conf"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func InitWebServer() {
	engine := gin.Default()
	agentmanager.Topo.Sdkmethod.RegisterHandlers(engine)
	InitRouter(engine)
	err := engine.Run(conf.Config().Topo.Server_addr)
	if err != nil {
		fmt.Printf("%+v\n", errors.Errorf("%s**2", err.Error())) // err top
		os.Exit(1)
	}
}

func InitRouter(router *gin.Engine) {
	api := router.Group("/plugin/api")
	{
		api.GET("/single_host/:uuid", SingleHostHandle)
		api.GET("/multi_host", MultiHostHandle)
	}
}
