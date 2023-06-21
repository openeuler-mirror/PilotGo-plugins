package plugin

import (
	"gitee.com/openeuler/PilotGo-plugins/sdk/plugin/client"
	"openeuler.org/PilotGo/redis-plugin/config"
)

const Version = "0.0.1"

func Init(conf *config.Redis) *client.PluginInfo {
	PluginInfo := client.PluginInfo{
		Name:        "redis",
		Version:     Version,
		Description: "redis",
		Author:      "wubijie",
		Email:       "wubijie@kylinos.cn",
		Url:         conf.URL,
		ReverseDest: conf.ReverseDest,
	}
	return &PluginInfo
}

// 请求prometheus插件接口，将gala-ops targets添加到监控清单当中
func addTargets(targets []string, url string) error {
	// TODO:
	// jobName := "redis"
	// url := url+"/api/add_targets"
	/*
	   - job_name: 'redis'
	     static_configs:
	       - targets: ['172.20.32.218:9121']
	*/
	return nil
}

func deleteTargets(targets []string, url string) error {
	// TODO:
	// jobName := "redis"
	// 删除yml文件中有关配置
	return nil
}

func MonitorTargets(targets []string) error {
	plugin, err := client.GetClient().GetPluginInfo("redis")
	if err != nil {
		return err
	}

	if err := addTargets(targets, plugin.Url); err != nil {
		return err
	}

	return nil
}
