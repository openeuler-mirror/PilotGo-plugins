package dao

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/global"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/model"
)

func GetScriptVersions(scriptId string) (*model.ScriptVersionResponse, error) {
	sql := `
SELECT
    s.id AS script_id,
    s.name AS name,
    s.script_type AS script_type,
    s.description AS description,
    s.is_public AS is_public,
    CAST((
        SELECT COALESCE(JSON_ARRAYAGG(
            JSON_OBJECT(
                'id', sv.id,
                'script_id', sv.script_id,
                'content', sv.content,
                'version', sv.version,
                'version_desc', sv.version_desc,
                'status', sv.status = 1,
                'creator', sv.creator,
                'created_at', sv.created_at,
                'last_modify_user', sv.last_modify_user,
                'last_modify_updated_at', sv.last_modify_updated_at
            )
        ), JSON_ARRAY())
        FROM script_version sv
        WHERE sv.script_id = s.id
    ) AS CHAR) AS versions,
    CAST((
        SELECT COALESCE(JSON_ARRAYAGG(
            JSON_OBJECT(
                'id', t.id,
                'name', t.name,
                'description', t.description,
                'creator', t.creator,
                'created_at', t.created_at,
                'last_modify_user', t.last_modify_user,
                'last_modify_updated_at', t.last_modify_updated_at
            )
        ), JSON_ARRAY())
        FROM tag t
        WHERE FIND_IN_SET(t.name, s.tags)
    ) AS CHAR) AS tags
FROM script s
WHERE s.id = ?
`
	var row model.RawScriptVersion
	if err := global.App.MySQL.Raw(sql, scriptId).Scan(&row).Error; err != nil {
		return &model.ScriptVersionResponse{}, fmt.Errorf("查询脚本版本失败: %w", err)
	}

	var tags []model.Tag
	if row.Tags == "" {
		tags = []model.Tag{}
	} else if err := json.Unmarshal([]byte(row.Tags), &tags); err != nil {
		return &model.ScriptVersionResponse{}, fmt.Errorf("解析标签失败: %w", err)
	}

	var scriptVersions []model.ScriptVersion
	if row.Versions == "" {
		scriptVersions = []model.ScriptVersion{}
	} else if err := json.Unmarshal([]byte(row.Versions), &scriptVersions); err != nil {
		return &model.ScriptVersionResponse{}, fmt.Errorf("解析版本失败: %w", err)
	}

	resp := &model.ScriptVersionResponse{
		ID:             row.ScriptID,
		Name:           row.Name,
		ScriptType:     row.ScriptType,
		Description:    row.Description,
		IsPublic:       row.IsPublic,
		Tags:           tags,
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
	return global.App.MySQL.Model(&model.ScriptVersion{}).Where("id = ? AND script_id = ?", id, scriptId).Updates(sv).Error
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
