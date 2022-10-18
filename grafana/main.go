package main

import (
	"fmt"

	"openeuler.org/PilotGo/plugin-sdk/plugin"
)

const Version = "0.0.1"

func main() {
	fmt.Println("hello grafana")
	client := plugin.DefaultClient(&plugin.PluginInfo{
		Name:        "grafana",
		Version:     "Version",
		Description: "grafana可视化工具支持",
		Author:      "guozhengxin",
		Email:       "guozhengxin@kylinos.cn",
	})

	client.Serve()
}
