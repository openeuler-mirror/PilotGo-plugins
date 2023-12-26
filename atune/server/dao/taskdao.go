package dao

import (
	"time"

	"gitee.com/openeuler/PilotGo/sdk/response"
	"openeuler.org/PilotGo/atune-plugin/db"
	"openeuler.org/PilotGo/atune-plugin/global"
	"openeuler.org/PilotGo/atune-plugin/model"
)

func QueryTaskLists(query *response.PaginationQ) ([]*model.Tasks, int64, error) {
	var tasks []*model.Tasks
	if err := db.MySQL().Preload("Tune").Preload("RunResults").Order("id desc").Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&tasks).Error; err != nil {
		return nil, 0, err
	}

	var total int64
	if err := db.MySQL().Model(&tasks).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return tasks, total, nil
}

func SaveTask(task *model.Tasks) (int, error) {
	if err := db.MySQL().Create(&task).Error; err != nil {
		return 0, err
	}
	return task.ID, nil
}
func UpdateTask(dbtaskid int) error {
	var t model.Tasks
	update_time := time.Now().Format("2006-01-02 15:04:05")
	task := model.Tasks{
		TaskStatus: global.Completed,
		UpdateTime: update_time,
	}
	if err := db.MySQL().Model(&t).Where("id = ?", dbtaskid).Updates(&task).Error; err != nil {
		return err
	}
	return nil
}
