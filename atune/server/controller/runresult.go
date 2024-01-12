package controller

import (
	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/atune-plugin/service"
)

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
