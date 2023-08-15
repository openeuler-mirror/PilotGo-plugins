package handler

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	api := router.Group("/plugin/topo/api")
	{
		api.PUT("/", Temp)
	}
}
