package scriptlibrary

import (
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/job_history/controller"
)

func JobHistoryHandler(router *gin.RouterGroup) {
	api := router.Group("/jobhistory")
	{
		api.POST("/add", controller.JobActionHistory)
	}
}
