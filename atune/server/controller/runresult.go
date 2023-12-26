package controller

import (
	"gitee.com/openeuler/PilotGo/sdk/common"
	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/atune-plugin/plugin"
	"openeuler.org/PilotGo/atune-plugin/service"
)

func RunCommand(c *gin.Context) {
	d := &struct {
		MachineUUIDs []string `json:"machine_uuids"`
		Command      string   `json:"command"`
		TuneID       int      `json:"tune_id"`
		TaskName     string   `json:"task_name"`
	}{}
	if err := c.ShouldBind(d); err != nil {
		logger.Debug("绑定批次参数失败：%s", err)
		response.Fail(c, nil, "parameter error")
		return
	}
	dbtaskid, err := service.SaveTask(d.Command, d.TaskName, d.MachineUUIDs, d.TuneID)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	run_result := func(result []*common.CmdResult) {
		for _, res := range result {
			logger.Info("结果：%v", *res)
			if err := service.ProcessResult(res, d.Command, dbtaskid); err != nil {
				logger.Error("处理结果失败：%v", err.Error())
			}
		}
	}

	dd := &common.Batch{
		MachineUUIDs: d.MachineUUIDs,
	}
	err = plugin.GlobalClient.RunCommandAsync(dd, d.Command, run_result)
	if err != nil {
		logger.Error("远程调用失败：%v", err.Error())
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "指令下发完成")
}

func DeleteResult(c *gin.Context) {
	resultdel := struct {
		ResultID []int `json:"ids"`
	}{}
	if err := c.Bind(&resultdel); err != nil {
		response.Fail(c, nil, "parameter error")
		return
	}

	if err := service.DeleteResult(resultdel.ResultID); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "已删除")
}

func SearchResult(c *gin.Context) {
	searchKey := c.Query("searchKey")

	query := &response.PaginationQ{}
	if err := c.ShouldBindQuery(query); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	data, total, err := service.SearchResult(searchKey, query)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.DataPagination(c, data, total, query)
}
