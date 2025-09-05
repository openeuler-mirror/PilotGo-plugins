package customscripts

import "github.com/gin-gonic/gin"

func CustomScriptsHandler(router *gin.RouterGroup) {
	api := router.Group("/custom_scripts")
	{
		api.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "test ok"})
		})
	}
}
