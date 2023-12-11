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
		logger.Debug("绑定批次参数失败：%s", err)
		response.Fail(c, nil, "parameter error")
		return
	}

	run_result := func(result []*common.RunResult) {
		for _, res := range result {
			logger.Info("结果：%v", *res)
			if err := service.ProcessResult(res, d.Command); err != nil {
				logger.Error("处理结果失败：%v", err.Error())
			}
		}
	}

	err := plugin.GlobalClient.RunCommandAsync(d.Batch, d.Command, run_result)
	if err != nil {
		logger.Error("远程调用失败：%v", err.Error())
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
