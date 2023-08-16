package clientmanager

import (
	"gitee.com/openeuler/PilotGo-plugins/sdk/plugin/client"
)

const Version = "0.0.1"

var PluginInfo = &client.PluginInfo{
	Name:        "topo",
	Version:     Version,
	Description: "system application architecture perception",
	Author:      "wangjunqi",
	Email:       "wangjunqi@kylinos.cn",
	Url:         "http://192.168.75.100:9995/plugin/topo",
}
