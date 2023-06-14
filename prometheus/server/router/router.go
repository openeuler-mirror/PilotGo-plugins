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

	// prometheus配置文件http方式获取监控target
	DBTarget := router.Group("/plugin/" + global.GlobalClient.PluginInfo.Name)
	{
		DBTarget.GET("target", httphandler.DBTargets)
	}

	// prometheus api代理
	prometheus := router.Group("/plugin/" + global.GlobalClient.PluginInfo.Name + "/api/v1")
	{
		prometheus.GET("/query", func(c *gin.Context) {
			c.Set("query", global.GlobalClient.PluginInfo.ReverseDest)
			httphandler.Query(c)
		})
		prometheus.GET("/query_range", func(c *gin.Context) {
			c.Set("query_range", global.GlobalClient.PluginInfo.ReverseDest)
			httphandler.QueryRange(c)
		})
		prometheus.GET("/targets", func(c *gin.Context) {
			c.Set("targets", global.GlobalClient.PluginInfo.ReverseDest)
			httphandler.PrometheusAPITargets(c)
		})

	}

	//prometheus target crud
	targetManager := router.Group("/plugin/" + global.GlobalClient.PluginInfo.Name)
	{
		targetManager.POST("addTarget", httphandler.AddPrometheusTarget)
		targetManager.DELETE("delTarget", httphandler.DeletePrometheusTarget)
	}
}
