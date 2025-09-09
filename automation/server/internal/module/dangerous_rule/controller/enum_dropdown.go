package controller

import (
	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum/rule"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum/script"
)

func ScriptTypeListHandler(c *gin.Context) {
	scriptTypes := script.GetScriptType()
	response.Success(c, scriptTypes, "success")
}

func ActionListHandler(c *gin.Context) {
	actions := rule.GetActions()
	response.Success(c, actions, "success")
}
