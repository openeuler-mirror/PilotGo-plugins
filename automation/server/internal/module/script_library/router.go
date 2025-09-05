package scriptlibrary

import "github.com/gin-gonic/gin"

func ScriptLibraryHandler(router *gin.RouterGroup) {
	api := router.Group("/scriptlibrary")
	{
		api.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "test ok"})
		})
	}
}
