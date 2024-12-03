/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugins licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Wed Jul 17 11:37:57 2024 +0800
 */
package service

import (
	"gitee.com/openeuler/PilotGo/sdk/common"
	plugin_manage "openeuler.org/PilotGo/PilotGo-plugin-event/client"
)

func AddExtentions() {
	var ex []common.Extention
	pe1 := &common.PageExtention{
		Type:       common.ExtentionPage,
		Name:       "事件日志",
		URL:        "/event",
		Permission: "plugin.event.page/menu",
	}
	ex = append(ex, pe1)
	plugin_manage.EventClient.RegisterExtention(ex)
}
