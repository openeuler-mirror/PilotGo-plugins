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
		Name:                tag.Name,
		Description:         tag.Description,
		Creator:             tag.Creator,
		CreatedAt:           time.Now().Format("2006-01-02 15:04:05"),
		LastModifyUser:      tag.Creator,
		LastModifyUpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	})
}
