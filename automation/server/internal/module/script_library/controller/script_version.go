package controller

import (
	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/service"
)

func GetScriptVersions(c *gin.Context) {
	script_id := c.Param("script_id")

	data, err := service.GetScriptVersions(script_id)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, data, "success")
}
