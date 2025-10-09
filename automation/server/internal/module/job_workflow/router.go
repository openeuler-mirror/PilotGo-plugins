package jobworkflow

import (
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/job_workflow/controller"
)

func WorkflowHandler(router *gin.RouterGroup) {
	api := router.Group("/workflows")
	{
		api.POST("/create", controller.CreateTemplate)
		api.POST("/update", controller.UpdateTemplate)
		api.GET("/query", controller.QueryTemplate)
		api.GET("/get", controller.GetTemplateById)
	}
}
