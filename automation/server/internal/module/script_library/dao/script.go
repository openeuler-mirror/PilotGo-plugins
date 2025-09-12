package dao

import (
	"strings"

	"gitee.com/openeuler/PilotGo/sdk/response"
	"gorm.io/gorm"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/global"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/model"
)

func AddScript(script *model.Script, scriptVersion *model.ScriptVersion) error {
	return global.App.MySQL.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(script).Error; err != nil {
			return err
		}

		if err := tx.Save(scriptVersion).Error; err != nil {
			return err
		}

		return nil
	})
}

func GetScripts(query *response.PaginationQ) ([]*model.ScriptResponse, int, error) {
	// 查询数据
	var scripts []*model.Script
	q := global.App.MySQL.Model(&model.Script{}).Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize)
	if err := q.Order("created_at desc").Find(&scripts).Error; err != nil {
		return nil, 0, err
	}

	// 组装 ScriptResponse
	var scriptResponses []*model.ScriptResponse
	for _, s := range scripts {
		sr := &model.ScriptResponse{
			ID:                  s.ID,
			Name:                s.Name,
			ScriptType:          s.ScriptType,
			Description:         s.Description,
			IsPublic:            s.IsPublic,
			Creator:             s.Creator,
			CreatedAt:           s.CreatedAt,
			LastModifyUser:      s.LastModifyUser,
			LastModifyUpdatedAt: s.LastModifyUpdatedAt,
		}

		tagNames := strings.Split(s.Tags, ",")
		var tags []model.Tag
		if len(tagNames) > 0 {
			if err := global.App.MySQL.Model(&model.Tag{}).Where("name IN ?", tagNames).Find(&tags).Error; err != nil {
				return nil, 0, err
			}
		}
		sr.Tags = tags

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
