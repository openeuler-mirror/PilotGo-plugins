package client

import (
	splugin "gitee.com/openeuler/PilotGo-plugins/sdk/plugin"
	"openeuler.org/PilotGo/gala-ops-plugin/config"
)

const Version = "0.0.1"

var globalClient *splugin.Client

func init() {
	globalClient = splugin.DefaultClient(&splugin.PluginInfo{
		Name:        "gala-ops",
		Version:     Version,
		Description: "gala-ops智能运维工具",
		Author:      "guozhengxin",
		Email:       "guozhengxin@kylinos.cn",
		Url:         "http://192.168.48.163:9999/plugin/grafana",
		// ReverseDest: "http://192.168.48.163:3000/",
	})
}

func Client() *splugin.Client {
	return globalClient
}

func StartClient(conf *config.HttpConf) {
	globalClient.Serve(conf.Addr)
}
