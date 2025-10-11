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

func DeleteTemplate(c *gin.Context) {
	var ids struct {
		ID []int `json:"id"`
	}
	if err := c.ShouldBindJSON(&ids); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	if err := service.DeleteTemplate(ids.ID); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "success")
}

func UpdateTemplate(c *gin.Context) {
	var template model.TaskTemplateDTO
	if err := c.ShouldBindJSON(&template); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	if err := service.UpdateTemplate(&template); err != nil {
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

func GetTemplateById(c *gin.Context) {
	templateId := c.Query("id")
	info, err := service.GetTemplateById(templateId)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, info, "success")
}

func PublishTemplate(c *gin.Context) {
	var id struct {
		ID        int    `json:"id"`
		NewStatus string `json:"new_status"`
	}
	if err := c.ShouldBindJSON(&id); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	if err := service.PublishTemplate(id.ID, id.NewStatus); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "success")
}
