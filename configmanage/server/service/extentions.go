package service

import (
	"gitee.com/openeuler/PilotGo/sdk/common"
	"openeuler.org/PilotGo/configmanage-plugin/global"
)

// 添加扩展点信息
func AddExtentions() {
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
	global.GlobalClient.RegisterExtention(ex)
}

// 添加权限信息
func AddPermission() {
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
	global.GlobalClient.RegisterPermission(ps)
}
