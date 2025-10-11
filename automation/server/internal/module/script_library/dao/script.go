package dao

import (
	"gorm.io/gorm"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/global"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/model"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/pkg/response"
)

func AddScript(script *model.Script, scriptVersion *model.ScriptVersion) error {
	return global.App.MySQL.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(script).Error; err != nil {
			return err
		}

		if err := tx.Create(scriptVersion).Error; err != nil {
			return err
		}

		return nil
	})
}

func GetScripts(query *response.PagedQuery) ([]*model.ScriptResponse, int, error) {
	// 查询数据
	var scripts []*model.Script
	q := global.App.MySQL.Model(&model.Script{}).Limit(query.PageSize).Offset((query.CurrentPage - 1) * query.PageSize)
	if err := q.Order("modify_time desc").Find(&scripts).Error; err != nil {
		return nil, 0, err
	}
	// 组装 ScriptResponse
	var scriptResponses []*model.ScriptResponse
	for _, s := range scripts {
		sr := &model.ScriptResponse{
			ID:          s.ID,
			Name:        s.Name,
			ScriptType:  s.ScriptType.String(),
			UsageType:   s.UsageType,
			Description: s.Description,
			ModifyUser:  s.ModifyUser,
			ModifyTime:  s.ModifyTime,
		}

		var tag model.Tag
		if len(s.Tag) > 0 {
			if err := global.App.MySQL.Model(&model.Tag{}).Where("name = ?", s.Tag).Find(&tag).Error; err != nil {
				return nil, 0, err
			}
		}
		sr.Tag = tag

		scriptResponses = append(scriptResponses, sr)
	}

	// 统计数目
	var total int64
	qc := global.App.MySQL.Model(&model.Script{})
	if err := qc.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return scriptResponses, int(total), nil
}

func UpdateScript(id string, s *model.Script) error {
	return global.App.MySQL.Model(&model.Script{}).Where("id = ?", id).Updates(s).Error
}

func DeleteScript(id string) error {
	return global.App.MySQL.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("script_id = ?", id).Delete(&model.ScriptVersion{}).Error; err != nil {
			return err
		}

		if err := tx.Where("id = ?", id).Delete(&model.Script{}).Error; err != nil {
			return err
		}
		return nil
	})
}
