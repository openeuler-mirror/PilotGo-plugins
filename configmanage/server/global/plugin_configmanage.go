package global

import (
	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
	"openeuler.org/PilotGo/configmanage-plugin/config"
)

const Version = "1.0.1"

func Init(plugin *config.ConfigPlugin) *client.PluginInfo {
	PluginInfo := client.PluginInfo{
		Name:        "configmanage",
		Version:     Version,
		Description: "configmanage-plugin",
		Author:      "wubijie",
		Email:       "wubijie@kylinos.cn",
		Url:         plugin.URL,
		PluginType:  "micro-app",
	}
	return &PluginInfo
}

const (
	Repo   = "repo"
	Host   = "host"
	SSH    = "ssh"
	SSHD   = "sshd"
	Sysctl = "sysctl"
	DNS    = "dns"
)
