package router

import (
	execscript "ant-agent/exec-script"
	"ant-agent/pkg/global"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"github.com/gin-gonic/gin"
)

func HttpServerInit() *gin.Engine {
	server := initRouters()

	logger.Debug("http server successfully started, listening on %s", global.HttpAddr)
	return server
}

// 后端路由
func initRouters() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	Router := gin.New()
	Router.Use(gin.Recovery())
	Router.Use(logger.RequestLogger([]string{}))

	// 注册各自的路由模块
	api := Router.Group("/plugin/automation")
	execscript.ExecScriptHandler(api)
	return Router
}
