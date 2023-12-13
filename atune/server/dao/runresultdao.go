package dao

import (
	"gitee.com/openeuler/PilotGo/sdk/response"
	"openeuler.org/PilotGo/atune-plugin/db"
	"openeuler.org/PilotGo/atune-plugin/model"
)

func QueryResults(query *response.PaginationQ) ([]*model.RunResult, int64, error) {
	var results []*model.RunResult
	if err := db.MySQL().Order("id desc").Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&results).Error; err != nil {
		return nil, 0, err
	}

	var total int64
	if err := db.MySQL().Model(&results).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return results, total, nil
}

func SaveRusult(result *model.RunResult) error {
	if err := db.MySQL().Create(&result).Error; err != nil {
		return err
	}
	return nil
}

func UpdateResult(machine_uuid string, result *model.RunResult) error {
	var r model.RunResult
	if err := db.MySQL().Model(&r).Where("machine_uuid = ?", machine_uuid).Updates(&result).Error; err != nil {
		return err
	}
	return nil
}

func DeleteResult(resultId int) error {
	var r model.RunResult
	if err := db.MySQL().Where("id = ?", resultId).Unscoped().Delete(&r).Error; err != nil {
		return err
	}
	return nil
}

func SearchResult(searchKey string, query *response.PaginationQ) ([]*model.RunResult, int64, error) {
	var result []*model.RunResult
	if err := db.MySQL().Order("id desc").Limit(query.PageSize).Offset((query.Page-1)*query.PageSize).Where("machine_ip LIKE ?", "%"+searchKey+"%").Find(&result).Error; err != nil {
		return nil, 0, nil
	}

	var total int64
	if err := db.MySQL().Where("machine_ip LIKE ?", "%"+searchKey+"%").Model(&result).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return result, total, nil
}

func IsExistCommandResult(machine_uuid string, command string) (bool, error) {
	var r model.RunResult
	err := db.MySQL().Where("machine_uuid = ? AND command = ?", machine_uuid, command).Find(&r).Error
	if err != nil {
		return false, err
	}
	if r.ID != 0 && command == r.Command {
		return true, nil
	}
	return false, nil
}
