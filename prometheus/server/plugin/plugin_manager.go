package plugin

import (
	"gitee.com/openeuler/PilotGo-plugins/sdk/plugin/client"
	"openeuler.org/PilotGo/prometheus-plugin/config"
)

func Init(conf *config.Prometheus) *client.PluginInfo {
	PluginInfo := client.PluginInfo{
		Name:        "Prometheus",
		Version:     "0.0.1",
		Description: "Prometheus开源系统监视和警报工具包",
		Author:      "zhanghan",
		Email:       "zhanghan@kylinos.cn",
		Url:         conf.URL,
		PluginType:  conf.PluginType,
		ReverseDest: conf.ReverseDest,
	}

	return &PluginInfo
}
