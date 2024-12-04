/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugins licensed under the Mulan Permissive Software License, Version 2.
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Wed Dec 4 14:45:21 2024 +0800
 */
package controller

import (
	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-event/db"
)

func EventsQueryHandler(c *gin.Context) {
	_start := c.Query("start")
	_stop := c.Query("stop")
	searchKey := c.Query("search")
	result, _ := db.Query(_start, _stop, searchKey)
	response.Success(c, result, "获取到数据")
}
