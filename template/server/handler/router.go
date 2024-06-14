package handler

import (
	"os"

	"gitee.com/openeuler/PilotGo-plugin-template/conf"
	"gitee.com/openeuler/PilotGo-plugin-template/pluginclient"
	"gitee.com/openeuler/PilotGo/sdk/logger"
	"github.com/gin-gonic/gin"
)

func InitWebServer() {
	if pluginclient.Global_Client == nil {
		logger.Error("Global_Client is nil")
		os.Exit(1)
	}

	go func() {
		engine := gin.Default()
		gin.SetMode(gin.ReleaseMode)
		pluginclient.Global_Client.RegisterHandlers(engine)
		InitRouter(engine)
		StaticRouter(engine)

		if conf.Global_Config.Template.Https_enabled {
			err := engine.RunTLS(conf.Global_Config.Template.Addr, conf.Global_Config.Template.CertFile, conf.Global_Config.Template.KeyFile)
			if err != nil {
				logger.Error(err.Error())
				os.Exit(1)
			}
		} else {
			err := engine.Run(conf.Global_Config.Template.Addr)
			if err != nil {
				logger.Error(err.Error())
				os.Exit(1)
			}
		}
	}()
}

func InitRouter(router *gin.Engine) {
	api := router.Group("/plugin/template/api")
	{
		api.GET("/do_something", DoSomethingHandle)
	}
}
