package router

import (
	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/redis-plugin/global"
	"openeuler.org/PilotGo/redis-plugin/httphandler"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(logger.LoggerDebug())
	router.Use(gin.Recovery())

	return router
}

func RegisterAPIs(router *gin.Engine) {
	global.GlobalClient.RegisterHandlers(router)

	pg := router.Group("/plugin/" + global.GlobalClient.PluginInfo.Name)
	{
		pg.GET("/query", func(c *gin.Context) {
			c.Set("query", global.GlobalClient.PluginInfo.ReverseDest)
			httphandler.InstallRedisExporter(c)
		})

	}

}
