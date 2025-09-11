package dao

import (
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/global"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/model"
)

func GetTags() ([]model.Tag, error) {
	var tags []model.Tag
	err := global.App.MySQL.Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func CreateTag(tag *model.Tag) error {
	return global.App.MySQL.Create(tag).Error
}
