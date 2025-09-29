package service

import (
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/job_workflow/dao"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/job_workflow/model"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/pkg/response"
)

func CreateTemplate(data *model.TaskTemplateDTO) error {
	if err := dao.CreateTemplate(data); err != nil {
		return err
	}
	return nil
}

func QueryTemplate(query *response.PagedQuery) ([]model.TaskTemplate, int, error) {
	templates, total, err := dao.QueryTemplates(query)
	if err != nil {
		return nil, 0, err
	}
	return templates, total, nil
}

func GetTemplateById(id string) (interface{}, error) {
	info, err := dao.GetTemplateById(id)
	if err != nil {
		return nil, err
	}
	return info, nil
}
