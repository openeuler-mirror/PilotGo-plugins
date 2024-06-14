package pluginclient

import "gitee.com/openeuler/PilotGo/sdk/plugin/client"

const Version = "1.0.1"

var PluginInfo = &client.PluginInfo{
	Name:        "template",
	Version:     Version,
	Description: "template plugin for PilotGo",
	Author:      "",
	Email:       "",
	Url:         "",
	PluginType:  "micro-app",
}
