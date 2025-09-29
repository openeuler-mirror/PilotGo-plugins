package controller

import (
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/job_workflow/model"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/job_workflow/service"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/pkg/response"
)

func CreateTemplate(c *gin.Context) {
	var template model.TaskTemplateDTO
	if err := c.ShouldBindJSON(&template); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	if err := service.CreateTemplate(&template); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "success")
}

func QueryTemplate(c *gin.Context) {
	query := &response.PagedQuery{}
	if err := c.ShouldBindQuery(query); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	templates, total, err := service.QueryTemplate(query)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.DataPaged(c, templates, total, query)
}
