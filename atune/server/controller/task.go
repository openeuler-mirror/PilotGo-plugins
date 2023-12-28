package controller

import (
	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/atune-plugin/service"
)

func TaskLists(c *gin.Context) {
	query := &response.PaginationQ{}
	err := c.ShouldBindQuery(query)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	data, total, err := service.QueryTaskLists(query)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	response.DataPagination(c, data, total, query)
}

func DeleteTask(c *gin.Context) {
	taskdel := struct {
		TaskID []int `json:"ids"`
	}{}
	if err := c.Bind(&taskdel); err != nil {
		response.Fail(c, nil, "parameter error")
		return
	}

	if err := service.DeleteTask(taskdel.TaskID); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "已删除任务")
}

func SearchTask(c *gin.Context) {
	search := c.Query("search")

	query := &response.PaginationQ{}
	if err := c.ShouldBindQuery(query); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	data, total, err := service.SearchTask(search, query)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.DataPagination(c, data, total, query)
}
