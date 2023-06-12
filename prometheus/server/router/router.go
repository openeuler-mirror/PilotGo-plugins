package router

import (
	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/prometheus-plugin/global"
	"openeuler.org/PilotGo/prometheus-plugin/httphandler"
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

	pg := router.Group("/plugin/" + global.GlobalClient.PluginInfo.Name + "/api/v1")
	{
		pg.GET("/query", func(c *gin.Context) {
			c.Set("query", global.GlobalClient.PluginInfo.ReverseDest)
			httphandler.Query(c)
		})
		pg.GET("/query_range", func(c *gin.Context) {
			c.Set("query_range", global.GlobalClient.PluginInfo.ReverseDest)
			httphandler.QueryRange(c)
		})
		pg.GET("/targets", func(c *gin.Context) {
			c.Set("targets", global.GlobalClient.PluginInfo.ReverseDest)
			httphandler.PrometheusAPITargets(c)
		})

	}

	target := router.Group("/plugin/" + global.GlobalClient.PluginInfo.Name)
	{
		target.GET("target", httphandler.DBTargets)
	}
}
