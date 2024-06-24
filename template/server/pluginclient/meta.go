package pluginclient

import "gitee.com/openeuler/PilotGo/sdk/plugin/client"

const Version = "1.0.1"

var PluginInfo = &client.PluginInfo{
	Name:        "template",
	Version:     Version,
	Description: "template plugin for PilotGo",
	Author:      "",
	Email:       "",
	Url:         "",   // 插件服务端地址，非插件配置文件中web服务器的监听地址
	PluginType:  "micro-app",
}
