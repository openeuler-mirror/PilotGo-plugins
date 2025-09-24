package dao

import (
	"sort"
	"time"

	"gorm.io/gorm"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/global"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/job_workflow/model"
)

func CreateTemplate(dto *model.TaskTemplateDTO) error {
	return global.App.MySQL.Transaction(func(tx *gorm.DB) error {
		// 1. 插入模板
		template := &model.TaskTemplate{
			Name:                dto.Template.Name,
			Description:         dto.Template.Description,
			Tags:                dto.Template.Tags,
			Creator:             dto.Template.Creator,
			CreatedAt:           time.Now().Format("2006-01-02 15:04:05"),
			LastModifyUser:      dto.Template.Creator,
			LastModifyUpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
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
			// 3.1 按 stepId 排序，补全链路
			sort.Slice(dto.Steps, func(i, j int) bool {
				return dto.Steps[i].StepId < dto.Steps[j].StepId
			})

			for i := range dto.Steps {
				dto.Steps[i].TemplateId = templateId
				if i > 0 {
					dto.Steps[i].PreviousStepId = dto.Steps[i-1].StepId
				}
				if i < len(dto.Steps)-1 {
					dto.Steps[i].NextStepId = dto.Steps[i+1].StepId
				}
			}

			if err := tx.Create(&dto.Steps).Error; err != nil {
				return err
			}

			// 3.2 设置模板的首尾步骤
			template.FirstStepId = dto.Steps[0].StepId
			template.LastStepId = dto.Steps[len(dto.Steps)-1].StepId
			// 3.4 回写模板首尾步骤
			if err := tx.Model(&model.TaskTemplate{}).
				Where("id = ?", templateId).
				Updates(map[string]interface{}{
					"first_step_id": template.FirstStepId,
					"last_step_id":  template.LastStepId,
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
