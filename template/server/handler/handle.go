package handler

import (
	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
)

func HelloWorldHandle(ctx *gin.Context) {
	// 数据库分页查询
	query := &response.PaginationQ{}
	err := ctx.ShouldBindQuery(query)
	if err != nil {
		logger.Warn(err.Error())
		response.Fail(ctx, nil, err.Error())
		return
	}

	if query.PageSize == 0 && query.Page == 0 {
		// 成功响应
		response_data := "hello world"
		response.Success(ctx, response_data, "")
		return

		// 失败响应
		// response.Fail(ctx, nil, "")
		// return
	}

	results := []string{
		"hello world1",
		"hello world2",
		"hello world3",
	}
	total := 0
	response.DataPagination(ctx, results, total, query)
}
