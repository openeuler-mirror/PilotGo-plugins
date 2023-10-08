package plugin

import (
	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
	"openeuler.org/PilotGo/atune-plugin/config"
)

var (
	GlobalClient *client.Client
)

func Init(plugin *config.PluginAtune) *client.PluginInfo {
	PluginInfo := client.PluginInfo{
		Name:        "atune",
		Version:     "1.0.1",
		Description: "A-Tune智能运维调优工具",
		Author:      "zhanghan",
		Email:       "zhanghan@kylinos.cn",
		Url:         plugin.URL,
		PluginType:  plugin.PluginType,
	}

	return &PluginInfo
}
