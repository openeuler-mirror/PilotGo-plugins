package handler

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	api := router.Group("/plugin/topo_a/api")
	{
		api.PUT("/metric", Raw_metric_data)
	}
}
