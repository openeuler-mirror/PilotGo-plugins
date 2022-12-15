package main

import (
	"fmt"

	"openeuler.org/PilotGo/plugin-sdk/plugin"
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
		Url:         "http://10.1.167.93:9090/",
		ReverseDest: "http://10.1.167.93:9090/",
	})

	client.Serve(":99090")
}
