/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugins licensed under the Mulan Permissive Software License, Version 2.
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Wed Dec 4 14:45:21 2024 +0800
 */
package controller

import (
	"net/http"

	"gitee.com/openeuler/PilotGo-plugins/event/sdk"
	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-event/db"
)

func EventsQueryHandler(c *gin.Context) {
	_start := c.Query("start")
	_stop := c.Query("stop")
	searchKey := c.Query("search")

	query := &response.PaginationQ{}
	err := c.ShouldBindQuery(query)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	result, err := db.Query(_start, _stop, searchKey)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	data, err := response.DataPaging(query, result, len(result))
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"ok":      true,
		"data":    data,
		"msgType": sdk.MessageTypes,
		"total":   len(result),
		"page":    query.Page,
		"size":    query.PageSize})
}
