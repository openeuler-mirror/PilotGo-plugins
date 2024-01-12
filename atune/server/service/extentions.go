package service

import (
	"gitee.com/openeuler/PilotGo/sdk/common"
	"openeuler.org/PilotGo/atune-plugin/plugin"
)

func AddExtentions() {
	var ex []common.Extention
	me1 := &common.MachineExtention{
		Type:       common.ExtentionMachine,
		Name:       "安装a-tune",
		URL:        "/plugin/atune/atune_install",
		Permission: "plugin.atune.agent/install",
	}
	me2 := &common.MachineExtention{
		Type:       common.ExtentionMachine,
		Name:       "卸载a-tune",
		URL:        "/plugin/atune/atune_uninstall",
		Permission: "plugin.atune.agent/uninstall",
	}
	pe1 := &common.PageExtention{
		Type:       common.ExtentionPage,
		Name:       "plugin-atune",
		URL:        "/plugin/atune/task",
		Permission: "plugin.prometheus.page/menu",
	}
	pe2 := &common.PageExtention{
		Type:       common.ExtentionPage,
		Name:       "plugin-atune",
		URL:        "/plugin/atune/template",
		Permission: "plugin.prometheus.page/menu",
	}
	ex = append(ex, me1, me2, pe1, pe2)
	plugin.GlobalClient.RegisterExtention(ex)
}
