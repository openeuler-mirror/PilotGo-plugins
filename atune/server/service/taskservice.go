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

func QueryTaskLists(query *response.PaginationQ) ([]*model.Tasks, int, error) {
	if data, total, err := dao.QueryTaskLists(query); err != nil {
		return nil, 0, err
	} else {
		return data, int(total), nil
	}
}

func SaveTask(cmd string, task_name string, uuids []string, tuneId int) (int, error) {
	task := &model.Tasks{
		TaskName:   task_name,
		TuneID:     tuneId,
		Script:     cmd,
		TaskStatus: dao.Executing,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}

	taskid, err := dao.SaveTask(task)
	if err != nil {
		return 0, errors.New("保存执行任务失败：" + err.Error())
	}

	for _, uuid := range uuids {
		result := &model.RunResult{
			TaskID:      taskid,
			MachineUUID: uuid,
			Command:     cmd,
			IsUpdate:    false,
		}
		if err := dao.SaveRusult(result); err != nil {
			logger.Error("save result uuid failed: %v", err.Error())
		}
	}
	return taskid, nil
}
func DeleteTask(taskId []int) error {
	if len(taskId) == 0 {
		return errors.New("请输入调优模板ID")
	}

	for _, tid := range taskId {
		if err := dao.DeleteTask(tid); err != nil {
			logger.Error("%v", strconv.Itoa(tid)+"未删除成功")
		}
	}
	return nil
}

func SearchTask(search string, query *response.PaginationQ) ([]*model.Tasks, int, error) {
	if data, total, err := dao.SearchTask(search, query); err != nil {
		return nil, 0, err
	} else {
		return data, int(total), nil
	}
}
