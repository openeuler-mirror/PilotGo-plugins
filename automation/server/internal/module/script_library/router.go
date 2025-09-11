package scriptlibrary

import (
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/controller"
)

func ScriptLibraryHandler(router *gin.RouterGroup) {
	api := router.Group("/scriptlibrary")
	{
		api.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "test ok"})
		})
	}

	tag := router.Group("/tag")
	{
		tag.GET("/list", controller.GetTagsHandler)
		tag.POST("/create", controller.CreateTagHandler)
	}
}
