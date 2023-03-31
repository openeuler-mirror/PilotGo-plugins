package main

import (
	"fmt"

	"openeuler.org/PilotGo/gala-ops-plugin/config"
	"openeuler.org/PilotGo/plugin-sdk/plugin"
)

const Version = "0.0.1"

func main() {
	fmt.Println("hello gala-ops")

	config.Init()

	client := plugin.DefaultClient(&plugin.PluginInfo{
		Name:        "gala-ops",
		Version:     Version,
		Description: "gala-ops智能运维工具",
		Author:      "guozhengxin",
		Email:       "guozhengxin@kylinos.cn",
		Url:         "http://192.168.48.163:9999/plugin/grafana",
		// ReverseDest: "http://192.168.48.163:3000/",
	})

	client.Serve(":8888")
}
