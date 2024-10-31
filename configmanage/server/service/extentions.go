package service

import (
	"gitee.com/openeuler/PilotGo/sdk/common"
	"openeuler.org/PilotGo/configmanage-plugin/global"
)

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
