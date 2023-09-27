package service

import (
	"errors"
	"strconv"

	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"gitee.com/openeuler/PilotGo-plugins/sdk/response"
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

	if ok, err := dao.IsExistTuneName(tune.TuneName); err == nil && ok {
		return errors.New("已存在该模板，勿重复添加")
	}

	if err := dao.SaveTune(&tune); err != nil {
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
