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

func UpdateTag(tag *model.Tag) error {
	return global.App.MySQL.Model(&model.Tag{}).Where("id = ?", tag.ID).Updates(tag).Error
}

func DeleteTag(id int) error {
	return global.App.MySQL.Delete(&model.Tag{}, id).Error
}
