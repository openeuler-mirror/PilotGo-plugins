package pluginclient

import (
	"context"

	"gitee.com/openeuler/PilotGo-plugin-template/conf"
	"gitee.com/openeuler/PilotGo/sdk/common"
	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
)

var Global_Client *client.Client

var Global_Context context.Context

func InitPluginClient() {
	if conf.Global_Config.Template.Https_enabled {
		PluginInfo.Url = "https://" + conf.Global_Config.Template.Addr
	} else {
		PluginInfo.Url = "http://" + conf.Global_Config.Template.Addr
	}

	Global_Client = client.DefaultClient(PluginInfo)

	// 注册插件扩展点
	var ex []common.Extention
	pe1 := &common.PageExtention{
		Type:       common.ExtentionPage,
		Name:       "主页面扩展",
		URL:        "/page",
		Permission: "plugin.template.page/menu",
	}
	me2 := &common.MachineExtention{
		Type:       common.ExtentionMachine,
		Name:       "机器扩展",
		URL:        "/host",
		Permission: "plugin.template/function",
	}
	be3 := &common.BatchExtention{
		Type:       common.ExtentionBatch,
		Name:       "批次扩展",
		URL:        "/batch",
		Permission: "plugin.template/function",
	}
	ex = append(ex, pe1, me2, be3)
	Global_Client.RegisterExtention(ex)

	Global_Context = context.Background()
}
