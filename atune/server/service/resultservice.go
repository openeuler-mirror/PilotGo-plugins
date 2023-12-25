package service

import (
	"errors"
	"strconv"

	"gitee.com/openeuler/PilotGo/sdk/common"
	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/response"
	"openeuler.org/PilotGo/atune-plugin/dao"
	"openeuler.org/PilotGo/atune-plugin/model"
)

func DeleteResult(resultId []int) error {
	if len(resultId) == 0 {
		return errors.New("请输入删除的ID")
	}

	for _, rid := range resultId {
		if err := dao.DeleteResult(rid); err != nil {
			logger.Error("%v", strconv.Itoa(rid)+"未删除成功")
		}
	}
	return nil
}
func ProcessResult(res *common.CmdResult, command string, dbtaskid int) error {
	result := &model.RunResult{
		MachineIP: res.MachineIP,
		RetCode:   res.RetCode,
		Stdout:    res.Stdout,
		Stderr:    res.Stderr,
		IsUpdate:  true,
	}

	if err := dao.UpdateResult(dbtaskid, res.MachineUUID, result); err != nil {
		return errors.New("更新命令执行结果失败：" + err.Error())
	}

	if ok, err := dao.IsUpdateAll(); ok && err == nil {
		if err := dao.UpdateTask(dbtaskid); err != nil {
			return errors.New("更新任务列表失败：" + err.Error())
		}
	}

	return nil
}

func SearchResult(searchKey string, query *response.PaginationQ) ([]*model.RunResult, int, error) {
	if data, total, err := dao.SearchResult(searchKey, query); err != nil {
		return nil, 0, err
	} else {
		return data, int(total), nil
	}
}
