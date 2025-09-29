package dao

import (
	"sort"
	"time"

	"gorm.io/gorm"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/global"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/job_workflow/model"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/pkg/response"
)

func CreateTemplate(dto *model.TaskTemplateDTO) error {
	return global.App.MySQL.Transaction(func(tx *gorm.DB) error {
		// 1. 插入模板
		template := &model.TaskTemplate{
			Name:        dto.Template.Name,
			Description: dto.Template.Description,
			Tags:        dto.Template.Tags,
			ModifyUser:  dto.Template.ModifyUser,
			ModifyTime:  time.Now().Format("2006-01-02 15:04:05"),
		}
		if err := tx.Create(template).Error; err != nil {
			return err
		}
		templateId := template.ID

		// 2. 插入变量
		if len(dto.Variables) > 0 {
			for i := range dto.Variables {
				dto.Variables[i].TemplateId = templateId
			}
			if err := tx.Create(&dto.Variables).Error; err != nil {
				return err
			}
		}

		// 3. 插入步骤 & 脚本
		if len(dto.Steps) > 0 {
			// 3.1 按 stepNum 排序，补全链路
			sort.Slice(dto.Steps, func(i, j int) bool {
				return dto.Steps[i].StepNum < dto.Steps[j].StepNum
			})

			for i := range dto.Steps {
				dto.Steps[i].TemplateId = templateId
				if i > 0 {
					dto.Steps[i].PreviousStepNum = dto.Steps[i-1].StepNum
				}
				if i < len(dto.Steps)-1 {
					dto.Steps[i].NextStepNum = dto.Steps[i+1].StepNum
				}
			}

			if err := tx.Create(&dto.Steps).Error; err != nil {
				return err
			}

			// 3.2 设置模板的首尾步骤
			template.FirstStepNum = dto.Steps[0].StepNum
			template.LastStepNum = dto.Steps[len(dto.Steps)-1].StepNum
			// 3.4 回写模板首尾步骤
			if err := tx.Model(&model.TaskTemplate{}).
				Where("id = ?", templateId).
				Updates(map[string]interface{}{
					"first_step_num": template.FirstStepNum,
					"last_step_num":  template.LastStepNum,
				}).Error; err != nil {
				return err
			}
		}

		if len(dto.Scripts) > 0 {
			for i := range dto.Scripts {
				dto.Scripts[i].TemplateId = templateId
			}
			if err := tx.Create(&dto.Scripts).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func QueryTemplates(query *response.PagedQuery) ([]model.TaskTemplate, int, error) {
	var templates []model.TaskTemplate
	q := global.App.MySQL.Model(&model.TaskTemplate{}).Limit(query.PageSize).Offset((query.CurrentPage - 1) * query.PageSize)
	if err := q.Find(&templates).Error; err != nil {
		return nil, 0, err
	}

	var total int64
	if err := global.App.MySQL.Model(&model.TaskTemplate{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return templates, int(total), nil
}

func GetTemplateById(id string) (interface{}, error) {
	var template model.TaskTemplate
	// 1. 查询模板基本信息
	if err := global.App.MySQL.Model(&model.TaskTemplate{}).Where("id = ?", id).First(&template).Error; err != nil {
		return nil, err
	}
	// 2. 查询变量
	var variables []model.TaskTemplateVariable
	if err := global.App.MySQL.Model(&model.TaskTemplateVariable{}).Where("template_id = ?", id).Find(&variables).Error; err != nil {
		return nil, err
	}
	// 3. 查询步骤
	var steps []model.TaskTemplateStep
	if err := global.App.MySQL.Model(&model.TaskTemplateStep{}).Where("template_id = ?", id).Find(&steps).Error; err != nil {
		return nil, err
	}
	var data = map[string]interface{}{}
	data["id"] = template.ID
	data["name"] = template.Name
	data["description"] = template.Description
	data["tags"] = template.Tags
	data["variableList"] = variables
	data["stepList"] = steps
	return data, nil
}
