package service

import (
	"errors"
	"time"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/response"
	"openeuler.org/PilotGo/atune-plugin/dao"
	"openeuler.org/PilotGo/atune-plugin/global"
	"openeuler.org/PilotGo/atune-plugin/model"
)

func QueryTaskLists(query *response.PaginationQ) ([]*model.Tasks, int, error) {
	if data, total, err := dao.QueryTaskLists(query); err != nil {
		return nil, 0, err
	} else {
		return data, int(total), nil
	}
}

func SaveTask(cmd string, task_name string, uuids []string) (int, error) {
	creat_time := time.Now().Format("2006-01-02 15:04:05")
	task := &model.Tasks{
		TaskName:   task_name,
		Script:     cmd,
		TaskStatus: global.Executing,
		CreateTime: creat_time,
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
