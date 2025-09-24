package service

import (
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/job_workflow/dao"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/job_workflow/model"
)

func CreateTemplate(data *model.TaskTemplateDTO) error {
	if err := dao.CreateTemplate(data); err != nil {
		return err
	}
	return nil
}
