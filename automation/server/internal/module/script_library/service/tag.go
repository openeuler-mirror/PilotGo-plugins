package service

import (
	"time"

	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/dao"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/model"
)

func GetTags() ([]model.Tag, error) {
	return dao.GetTags()
}

func CreateTag(tag *model.Tag) error {
	return dao.CreateTag(&model.Tag{
		Name:        tag.Name,
		Description: tag.Description,
		ModifyUser:  tag.ModifyUser,
		ModifyTime:  time.Now().Format("2006-01-02 15:04:05"),
	})
}

func UpdateTag(tag *model.Tag) error {
	return dao.UpdateTag(&model.Tag{
		ID:          tag.ID,
		Name:        tag.Name,
		Description: tag.Description,
		ModifyUser:  tag.ModifyUser,
		ModifyTime:  time.Now().Format("2006-01-02 15:04:05"),
	})
}

func DeleteTag(id int) error {
	return dao.DeleteTag(id)
}
