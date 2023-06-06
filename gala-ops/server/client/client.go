package client

import (
	"gitee.com/openeuler/PilotGo-plugins/sdk/plugin/client"
	"openeuler.org/PilotGo/gala-ops-plugin/config"
)

const Version = "0.0.1"

var globalClient *client.Client

func init() {
	globalClient = client.DefaultClient(&client.PluginInfo{
		Name:        "gala-ops",
		Version:     Version,
		Description: "gala-ops智能运维工具",
		Author:      "guozhengxin",
		Email:       "guozhengxin@kylinos.cn",
		Url:         "http://192.168.48.163:9999/plugin/gala-ops",
		// ReverseDest: "http://192.168.48.163:3000/",
	}, config.Config().Logopts)
}

func Client() *client.Client {
	return globalClient
}

func StartClient(conf *config.HttpConf) {
	globalClient.Serve(conf.Addr)
}
