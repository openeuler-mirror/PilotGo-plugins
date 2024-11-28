/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: wubijie <wubijie@kylinos.cn>
 * Date: Fri Dec 1 16:30:26 2023 +0800
 */
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
	PATH   = "path"
)
