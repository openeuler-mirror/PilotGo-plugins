package router

import (
	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/atune-plugin/config"
	"openeuler.org/PilotGo/atune-plugin/httphandler"
	"openeuler.org/PilotGo/atune-plugin/plugin"
)

func HttpServerInit(conf *config.HttpServer) error {

	go func() {
		r := setupRouter()

		logger.Info("start http service on: http://%s", conf.Addr)
		if err := r.Run(conf.Addr); err != nil {
			logger.Error("start http server failed:%v", err)
		}

	}()

	return nil
}
func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(logger.LoggerDebug())
	router.Use(gin.Recovery())

	registerAPIs(router)

	return router
}
func registerAPIs(router *gin.Engine) {
	logger.Debug("router register")
	plugin.GlobalClient.RegisterHandlers(router)

	atune := router.Group("/plugin/" + plugin.GlobalClient.PluginInfo.Name)
	{
		atune.GET("all", httphandler.GetAtuneAll)
		atune.GET("info", httphandler.GetAtuneInfo)
	}
}
