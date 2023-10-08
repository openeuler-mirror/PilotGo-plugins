package dao

import (
	"gitee.com/openeuler/PilotGo/sdk/response"
	"openeuler.org/PilotGo/atune-plugin/db"
	"openeuler.org/PilotGo/atune-plugin/model"
)

func QueryTunes(query *response.PaginationQ) ([]*model.Tunes, int64, error) {
	var tunes []*model.Tunes
	if err := db.MySQL().Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&tunes).Error; err != nil {
		return nil, 0, err
	}

	var total int64
	if err := db.MySQL().Model(&tunes).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return tunes, total, nil
}

func SaveTune(tune *model.Tunes) error {
	if err := db.MySQL().Create(&tune).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTune(tuneName string, tune *model.Tunes) error {
	var t model.Tunes
	if err := db.MySQL().Model(&t).Where("tune_name = ?", tuneName).Updates(&tune).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTune(tuneId int) error {
	var t model.Tunes
	if err := db.MySQL().Where("id = ?", tuneId).Unscoped().Delete(&t).Error; err != nil {
		return err
	}
	return nil
}

func IsExistTuneName(tuneName string) (bool, error) {
	var t model.Tunes
	err := db.MySQL().Where("tune_name = ?", tuneName).Find(&t).Error
	return t.ID != 0, err
}
