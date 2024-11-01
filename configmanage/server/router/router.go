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
			// 提供配置列表
			list.GET("/config_infos", controller.ConfigInfosHandler)
			// 根据列表中的configinfo_uuid获取某一个config信息
			list.GET("/config_info", controller.ConfigInfoHandler)
		}
		// 添加配置管理
		api.POST("/add", controller.AddConfigHandler)
		// 根据列表中的configinfo_uuid获取某一个具体的正在使用的config信息
		api.GET("/load", controller.LoadConfigHandler)
		// 下发配置管理
		api.POST("/apply", controller.ApplyConfigHandler)
	}
}
