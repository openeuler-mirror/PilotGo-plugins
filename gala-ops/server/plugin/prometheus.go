package plugin

import (
	"openeuler.org/PilotGo/gala-ops-plugin/client"
)

// 请求prometheus插件接口，将gala-ops targets添加到监控清单当中
func addTargets(targets []string, url string) error {
	// TODO:
	// jobName := "gala-gopher"
	// url := url+"/api/add_targets"

	return nil
}

func MonitorTargets(targets []string) error {
	c := client.Client()
	plugin, err := c.GetPluginInfo("prometheus")
	if err != nil {
		return err
	}

	if err := addTargets(targets, plugin.Url); err != nil {
		return err
	}

	return nil
}

func deleteTargets(targets []string, url string) error {
	// TODO:

	return nil
}

func DeleteTargets(targets []string) error {
	c := client.Client()
	plugin, err := c.GetPluginInfo("prometheus")
	if err != nil {
		return err
	}

	if err := deleteTargets(targets, plugin.Url); err != nil {
		return err
	}

	return nil
}
