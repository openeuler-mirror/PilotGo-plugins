package controller

import (
	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/model"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/service"
)

func GetScriptVersionsHandler(c *gin.Context) {
	script_id := c.Param("script_id")

	data, err := service.GetScriptVersions(script_id)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, data, "success")
}

func AddScriptVersionHandler(c *gin.Context) {
	script_id := c.Param("script_id")

	var scriptVersion model.ScriptVersion
	if err := c.ShouldBindJSON(&scriptVersion); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	if err := service.AddScriptVersion(script_id, &scriptVersion); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "success")
}

func UpdateScriptVersionHandler(c *gin.Context) {
	script_id := c.Param("script_id")

	var scriptVersion model.ScriptVersion
	if err := c.ShouldBindJSON(&scriptVersion); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	if err := service.UpdateScriptVersion(script_id, &scriptVersion); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "success")
}

func DeleteScriptVersionHandler(c *gin.Context) {
	script_id := c.Param("script_id")

	var id struct {
		ID int `json:"id"`
	}
	if err := c.ShouldBindJSON(&id); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	if err := service.DeleteScriptVersion(id.ID, script_id); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "success")
}
