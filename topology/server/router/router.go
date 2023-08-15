package router

import (
	"gitee.com/openeuler/PilotGo-plugin-topology-server/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	api := router.Group("/plugin/topo/api")
	{
		api.PUT("/", handler.Temp)
	}
}
