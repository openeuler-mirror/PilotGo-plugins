package main

import (
	"fmt"

	splugin "gitee.com/openeuler/PilotGo-plugins/sdk/plugin"
)

const Version = "0.0.1"

func main() {
	fmt.Println("hello grafana")
	client := splugin.DefaultClient(&splugin.PluginInfo{
		Name:        "grafana",
		Version:     "Version",
		Description: "grafana可视化工具支持",
		Author:      "guozhengxin",
		Email:       "guozhengxin@kylinos.cn",
		Url:         "http://localhost:9999/plugin/grafana",
		ReverseDest: "http://192.168.28.232:3000/",
	})

	client.Serve(":9999")
}
