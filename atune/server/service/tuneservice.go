package service

import (
	"errors"
	"strconv"
	"time"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/response"
	"openeuler.org/PilotGo/atune-plugin/dao"
	"openeuler.org/PilotGo/atune-plugin/model"
)

func QueryTunes(query *response.PaginationQ) ([]*model.Tunes, int, error) {
	if data, total, err := dao.QueryTunes(query); err != nil {
		return nil, 0, err
	} else {
		return data, int(total), nil
	}
}

func SaveTune(tune model.Tunes) error {
	if tune.TuneName == "" {
		return errors.New("内容为空，请检查输入内容")
	}
	t := model.Tunes{
		TuneName:      tune.TuneName,
		Description:   tune.Description,
		CreateTime:    time.Now().Format("2006-01-02 15:04:05"),
		WorkDirectory: tune.WorkDirectory,
		Prepare:       tune.Prepare,
		Tune:          tune.Tune,
		Restore:       tune.Restore,
		Notes:         tune.Notes,
	}
	if err := dao.SaveTune(&t); err != nil {
		return err
	}

	return nil
}

func DeleteTune(tuneId []int) error {
	if len(tuneId) == 0 {
		return errors.New("请输入调优模板ID")
	}

	for _, tid := range tuneId {
		if err := dao.DeleteTune(tid); err != nil {
			logger.Error("%v", strconv.Itoa(tid)+"未删除成功")
		}
	}
	return nil
}

func UpdateTune(t model.Tunes) error {
	updatetune := &model.Tunes{
		Description:   t.Description,
		UpdateTime:    time.Now().Format("2006-01-02 15:04:05"),
		WorkDirectory: t.WorkDirectory,
		Prepare:       t.Prepare,
		Tune:          t.Tune,
		Restore:       t.Restore,
		Notes:         t.Notes,
	}

	if err := dao.UpdateTune(t.TuneName, updatetune); err != nil {
		return err
	}
	return nil
}

func SearchTune(tune_name string, query *response.PaginationQ) ([]*model.Tunes, int, error) {
	if data, total, err := dao.SearchTune(tune_name, query); err != nil {
		return nil, 0, err
	} else {
		return data, int(total), nil
	}
}
