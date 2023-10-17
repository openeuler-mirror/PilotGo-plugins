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
	d := &common.CmdStruct{}
	if err := c.ShouldBind(d); err != nil {
		logger.Debug("bind batch param error:%s", err)
		response.Fail(c, nil, "parameter error")
		return
	}

	run_result := func(result []*common.RunResult) {
		for _, res := range result {
			if err := service.ProcessResult(res, d.Command); err != nil {
				logger.Error("%v", err.Error())
			}
		}
	}

	err := plugin.GlobalClient.RunCommandAsync(d.Batch, d.Command, run_result)
	if err != nil {
		logger.Error("%v", err.Error())
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "指令下发完成")
}

func QueryResults(c *gin.Context) {
	query := &response.PaginationQ{}
	err := c.ShouldBindQuery(query)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	data, total, err := service.QueryResults(query)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	response.DataPagination(c, data, total, query)
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
