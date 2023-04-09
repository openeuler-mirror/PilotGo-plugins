package client

import (
	"openeuler.org/PilotGo/gala-ops-plugin/config"
	"openeuler.org/PilotGo/plugin-sdk/plugin"
)

const Version = "0.0.1"

var globalClient *plugin.Client

func init() {
	globalClient = plugin.DefaultClient(&plugin.PluginInfo{
		Name:        "gala-ops",
		Version:     Version,
		Description: "gala-ops智能运维工具",
		Author:      "guozhengxin",
		Email:       "guozhengxin@kylinos.cn",
		Url:         "http://192.168.48.163:9999/plugin/grafana",
		// ReverseDest: "http://192.168.48.163:3000/",
	})
}

func Client() *plugin.Client {
	return globalClient
}

func StartClient(conf *config.HttpConf) {
	globalClient.Serve(conf.Addr)
}
