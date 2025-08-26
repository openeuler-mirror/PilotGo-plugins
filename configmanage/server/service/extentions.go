/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin licensed under the Mulan Permissive Software License, Version 2.
 * See LICENSE file for more details.
 * Author: wubijie <wubijie@kylinos.cn>
 * Date: Thu Oct 31 14:15:31 2024 +0800
 */
package service

import (
	"gitee.com/openeuler/PilotGo/sdk/common"
)

// 添加扩展点信息
func GetExtentions() []common.Extention {
	var ex []common.Extention
	pe1 := &common.PageExtention{
		Type:       common.ExtentionPage,
		Name:       "添加配置",
		URL:        "/add",
		Permission: "plugin.configmanage.page/menu",
	}
	pe2 := &common.PageExtention{
		Type:       common.ExtentionPage,
		Name:       "查看配置",
		URL:        "/view",
		Permission: "plugin.configmanage.page/menu",
	}
	ex = append(ex, pe1, pe2)
	return ex
}

// 添加权限信息
func GetPermission() []common.Permission {
	var ps []common.Permission
	p1 := common.Permission{
		Resource: "plugin.configmanage",
		Operate:  "agent_install",
	}
	p2 := common.Permission{
		Resource: "plugin.configmanage",
		Operate:  "agent_uninstall",
	}
	ps = append(ps, p1, p2)
	return ps
}
