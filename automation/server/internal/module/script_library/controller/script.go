package controller

import (
	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/model"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/service"
)

func AddScript(c *gin.Context) {
	var script model.ScriptWithVersion
	if err := c.ShouldBindJSON(&script); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	if err := service.AddScript(&script); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "success")
}
