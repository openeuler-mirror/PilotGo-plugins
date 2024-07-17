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
