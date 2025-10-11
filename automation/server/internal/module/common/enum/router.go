package enum

import (
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum/rule"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum/script"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum/workflow"
)

func EnumHandler(router *gin.RouterGroup) {
	api := router.Group("/enum")
	{
		api.GET("/ruleAction", rule.RuleActionListHandler)
		api.GET("/scriptType", script.ScriptTypeListHandler)
		api.GET("/scriptStatus", script.ScriptPublishStatusListHandler)
		api.GET("/workflowStatus", workflow.WorkflowPublishStatusListHandler)
		api.GET("/workflowStepType", workflow.StepTypeListHandler)
	}
}
