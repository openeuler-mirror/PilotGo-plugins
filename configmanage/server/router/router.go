package router

import (
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/configmanage-plugin/controller"
	"openeuler.org/PilotGo/configmanage-plugin/global"
)

// gin.egnine充当server的角色
func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	return router
}

func RegisterAPIs(router *gin.Engine) {
	//输出插件初始化的信息
	global.GlobalClient.RegisterHandlers(router)
	pg := router.Group("/plugin/" + global.GlobalClient.PluginInfo.Name)
	{
		pg.POST("/repoadd", controller.AddRepoHandler)
		pg.POST("/repoapply", controller.RepoApply)

	}
}
