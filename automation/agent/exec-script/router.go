package execscript

import (
	"ant-agent/exec-script/controller"

	"github.com/gin-gonic/gin"
)

func ExecScriptHandler(router *gin.RouterGroup) {
	api := router.Group("/script")
	{
		api.POST("/exec", controller.ExecScript)
	}
}
