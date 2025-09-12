package scriptlibrary

import (
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/controller"
)

func ScriptLibraryHandler(router *gin.RouterGroup) {
	api := router.Group("/scriptlibrary")
	{
		api.POST("/add", controller.AddScriptHandler)
		api.GET("/list", controller.ScriptListHandler)
		api.PUT("/update", controller.UpdateScriptHandler)
		api.DELETE("delete", controller.DeleteScriptHandler)

		versionGroup := api.Group("/:script_id/scriptVersion")
		{
			versionGroup.GET("/list", controller.GetScriptVersionsHandler)
			versionGroup.POST("/add", controller.AddScriptVersionHandler)
			versionGroup.PUT("/update", controller.UpdateScriptVersionHandler)
		}
	}

	tag := router.Group("/tag")
	{
		tag.GET("/list", controller.GetTagsHandler)
		tag.POST("/create", controller.CreateTagHandler)
		tag.PUT("/update", controller.UpdateTagHandler)
		tag.DELETE("/delete", controller.DeleteTagHandler)
	}
}
