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
	api := router.Group("/plugin/" + global.GlobalClient.PluginInfo.Name)
	{
		list := api.Group("/list")
		{
			// 提供配置文件类型的列表
			list.GET("/config_type", controller.ConfigTypeListHandler)
		}
		api.POST("/add", controller.AddConfigHandler)
		api.GET("/load", controller.LoadConfigHandler)
		api.POST("/apply", controller.ApplyConfigHandler)
	}
}
