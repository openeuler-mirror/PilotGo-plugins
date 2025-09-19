package dao

import (
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/global"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/job_workflow/model"
)

func GetTemplateByID(id int) (*model.TaskTemplate, error) {
	var template model.TaskTemplate
	if err := global.App.MySQL.First(&template, id).Error; err != nil {
		return nil, err
	}
	return &template, nil
}

func GetAllTemplates() ([]*model.TaskTemplate, error) {
	var templates []*model.TaskTemplate
	if err := global.App.MySQL.Find(&templates).Error; err != nil {
		return nil, err
	}
	return templates, nil
}

func CreateTemplate(t *model.TaskTemplate) error {
	return global.App.MySQL.Create(t).Error
}

func UpdateTemplate(id int, t *model.TaskTemplate) error {
	return global.App.MySQL.Model(&model.TaskTemplate{}).Where("id = ?", id).Updates(t).Error
}

func DeleteTemplate(id int) error {
	return global.App.MySQL.Delete(&model.TaskTemplate{}, id).Error
}
