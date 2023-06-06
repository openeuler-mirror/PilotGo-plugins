package main

import (
	"fmt"

	"gitee.com/openeuler/PilotGo-plugins/sdk/plugin"
	Router "openeuler.org/PilotGo/prometheus-plugin/router"
)

func main() {
	fmt.Println("hello prometheus")

	client := Router.DefaultClient(&plugin.PluginInfo{
		Name:        "Prometheus",
		Version:     "Version",
		Description: "Prometheus开源系统监视和警报工具包",
		Author:      "zhanghan",
		Email:       "zhanghan@kylinos.cn",
		Url:         "http://localhost:8090",
		ReverseDest: "http://localhost:9090",
	})

	client.Serve(":8090")
}
