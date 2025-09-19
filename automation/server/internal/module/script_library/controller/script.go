package controller

import (
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/model"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/service"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/pkg/response"
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

func DeleteScriptHandler(c *gin.Context) {
	var id struct {
		ID string `json:"id"`
	}
	if err := c.ShouldBindJSON(&id); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	if err := service.DeleteScript(id.ID); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "success")
}
