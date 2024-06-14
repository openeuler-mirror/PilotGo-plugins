package handler

import (
	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
)

func DoSomethingHandle(ctx *gin.Context) {
	query := &response.PaginationQ{}
	err := ctx.ShouldBindQuery(query)
	if err != nil || (query.PageSize == 0 && query.Page == 0) {
		logger.Warn(err.Error())
		response.Fail(ctx, nil, err.Error())
		return
	}

	// condition: 1 成功响应 2 失败响应 3 翻页查询响应
	condition := 1
	switch condition {
	case 1:
		response.Success(ctx, nil, "")
	case 2:
		response.Fail(ctx, nil, "")
		return
	default:
		results := []string{}
		total := 0
		response.DataPagination(ctx, results, total, query)
	}
}