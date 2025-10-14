package jobaction

import (
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/job_action/controller"
)

func JobActionHandler(router *gin.RouterGroup) {
	api := router.Group("/action")
	{
		api.POST("/exec", controller.ExecScriptHandler)
	}
}
