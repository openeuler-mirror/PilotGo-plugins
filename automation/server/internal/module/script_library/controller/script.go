package controller

import (
	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/model"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/service"
)

func AddScriptHandler(c *gin.Context) {
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

func ScriptListHandler(c *gin.Context) {
	query := &response.PaginationQ{}
	if err := c.ShouldBindQuery(query); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	scripts, total, err := service.GetScripts(query)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.DataPagination(c, scripts, total, query)
}

func UpdateScriptHandler(c *gin.Context) {
	var script model.Script
	if err := c.ShouldBindJSON(&script); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	if err := service.UpdateScript(&script); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "success")
}
