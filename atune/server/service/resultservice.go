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

func QueryResults(query *response.PaginationQ) ([]*model.RunResult, int, error) {
	if data, total, err := dao.QueryResults(query); err != nil {
		return nil, 0, err
	} else {
		return data, int(total), nil
	}
}

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
func ProcessResult(res *common.RunResult, command string) error {
	result := &model.RunResult{
		MachineUUID:   res.CmdResult.MachineUUID,
		MachineIP:     res.CmdResult.MachineIP,
		Command:       command,
		RetCode:       res.CmdResult.RetCode,
		Stdout:        res.CmdResult.Stdout,
		Stderr:        res.CmdResult.Stderr,
		ResponseError: res.Error.(string),
	}

	ok, err := dao.IsExistCommandResult(res.CmdResult.MachineUUID, command)
	if ok && err == nil {
		if Err := dao.UpdateResult(res.CmdResult.MachineUUID, result); Err != nil {
			return errors.New("更新命令执行结果失败：" + Err.Error())
		}
	}
	if !ok && err == nil {
		if Err := dao.SaveRusult(result); Err != nil {
			return errors.New("保存命令执行结果失败：" + Err.Error())
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
