package controller

import (
	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/model"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/service"
)

func GetTagsHandler(c *gin.Context) {
	tags, err := service.GetTags()
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, tags, "success")
}

func CreateTagHandler(c *gin.Context) {
	var tag model.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	if err := service.CreateTag(&tag); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "success")
}
