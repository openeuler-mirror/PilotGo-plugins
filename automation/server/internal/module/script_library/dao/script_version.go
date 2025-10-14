package dao

import (
	"encoding/json"
	"fmt"
	"strconv"

	"gorm.io/gorm"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/global"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum/script"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/model"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/pkg/utils"
)

func GetScriptVersions(scriptId string) (*model.ScriptVersionResponse, error) {
	sql := `
SELECT
    s.id AS script_id,
    s.name AS name,
    s.usage_type AS usage_type,
    s.script_type AS script_type,
    s.description AS description,
    CAST((
        SELECT COALESCE(JSON_ARRAYAGG(
            JSON_OBJECT(
                'id', sv.id,
                'script_id', sv.script_id,
                'content', sv.content,
                'params', sv.params,
                'version', sv.version,
                'version_desc', sv.version_desc,
                'status', sv.status,
                'modify_user', sv.modify_user,
                'modify_time', sv.modify_time
            )
        ), JSON_ARRAY())
        FROM script_version sv
        WHERE sv.script_id = s.id
    ) AS CHAR) AS versions,
    CAST((
        SELECT COALESCE(JSON_OBJECT(
           'id', t.id,
           'name', t.name,
           'description', t.description,
           'modify_user', t.modify_user,
           'modify_time', t.modify_time
        ), JSON_OBJECT())
        FROM tag t
        WHERE FIND_IN_SET(t.name, s.tag)
    ) AS CHAR) AS tag
FROM script s
WHERE s.id = ?
`
	var row model.RawScriptVersion
	if err := global.App.MySQL.Raw(sql, scriptId).Scan(&row).Error; err != nil {
		return &model.ScriptVersionResponse{}, fmt.Errorf("查询脚本版本失败: %w", err)
	}

	var tag model.Tag
	if row.Tag == "" {
		tag = model.Tag{}
	} else if err := json.Unmarshal([]byte(row.Tag), &tag); err != nil {
		return &model.ScriptVersionResponse{}, fmt.Errorf("解析标签失败: %w", err)
	}

	var scriptVersions []model.ScriptVersion
	if row.Versions == "" {
		scriptVersions = []model.ScriptVersion{}
	} else if err := json.Unmarshal([]byte(row.Versions), &scriptVersions); err != nil {
		return &model.ScriptVersionResponse{}, fmt.Errorf("解析版本失败: %w", err)
	}

	script_type, _ := strconv.Atoi(row.ScriptType)
	resp := &model.ScriptVersionResponse{
		ID:             row.ScriptID,
		Name:           row.Name,
		ScriptType:     script.ScriptTypeMap[script_type],
		Description:    row.Description,
		Tag:            tag,
		ScriptVersions: scriptVersions,
	}

	return resp, nil
}

func AddScriptVersion(sv *model.ScriptVersion) error {
	return global.App.MySQL.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(sv).Error; err != nil {
			return err
		}
		return nil
	})
}

func UpdateScriptVersion(id int, scriptId string, sv *model.ScriptVersion) error {
	return global.App.MySQL.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.ScriptVersion{}).Where("id = ? AND script_id = ?", id, scriptId).Updates(sv).Error; err != nil {
			return err
		}
		return nil
	})
}

func DeleteScriptVersion(id int, scriptId string) error {
	return global.App.MySQL.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ? AND script_id = ?", id, scriptId).Delete(&model.ScriptVersion{}).Error; err != nil {
			return err
		}

		var count int64
		if err := tx.Model(&model.ScriptVersion{}).Where("script_id = ?", scriptId).Count(&count).Error; err != nil {
			return err
		}

		if count == 0 {
			if err := tx.Where("id = ?", scriptId).Delete(&model.Script{}).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func PublishScriptVersion(id int, scriptId string, newStatus string) error {
	return global.App.MySQL.Transaction(func(tx *gorm.DB) error {
		if newStatus == script.Published.String() {
			if err := tx.Model(&model.ScriptVersion{}).Where("script_id = ?", scriptId).Update("status", script.Develop).Error; err != nil {
				return err
			}
		}

		if err := tx.Model(&model.ScriptVersion{}).Where("id = ? AND script_id = ?", id, scriptId).Update("status", script.ParseScriptPublishStatus(newStatus)).Error; err != nil {
			return err
		}

		return nil
	})
}

func GetLatestScriptVersion(scriptId string) (string, error) {
	var sv model.ScriptVersion
	if err := global.App.MySQL.Model(&model.ScriptVersion{}).Where("script_id = ?", scriptId).Order("id DESC").First(&sv).Error; err != nil {
		return "", err
	}
	return sv.Version, nil
}

func GetPublishedScriptByScriptId(scriptId string) (string, string, error) {
	var s model.Script
	if err := global.App.MySQL.Model(&model.Script{}).Where("script_id = ?", scriptId).First(&s).Error; err != nil {
		return "", "", err
	}
	var sv model.ScriptVersion
	if err := global.App.MySQL.Model(&model.ScriptVersion{}).Where("script_id = ? AND status = ?", scriptId, script.Published).First(&sv).Error; err != nil {
		return s.ScriptType.String(), "", err
	}
	return s.ScriptType.String(), utils.EncodeScriptContent(sv.Content), nil
}
