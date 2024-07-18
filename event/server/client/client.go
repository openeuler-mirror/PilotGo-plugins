package client

import (
	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
	"openeuler.org/PilotGo/PilotGo-plugin-event/config"
)

var (
	EventClient *client.Client
)

func Init(plugin *config.PluginEvent) *client.PluginInfo {
	PluginInfo := client.PluginInfo{
		Name:        "event",
		Version:     "1.0.0",
		Description: "时间轴事件插件",
		Author:      "zhanghan",
		Email:       "zhanghan@kylinos.cn",
		Url:         plugin.URL,
		PluginType:  plugin.PluginType,
	}

	return &PluginInfo
}
