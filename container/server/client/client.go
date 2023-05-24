package client

import (
	splugin "gitee.com/openeuler/PilotGo-plugins/sdk/plugin"
	"openeuler.org/PilotGo/container-plugin/config"
)

const Version = "0.0.1"

var globalClient *splugin.Client

func init() {
	globalClient = splugin.DefaultClient(&splugin.PluginInfo{
		Name:        "container",
		Version:     Version,
		Description: "Container management plugin",
		Author:      "wangjunqi",
		Email:       "wangjunqi@kylinos.cn",
		Url:         "http://192.168.75.100:9999/plugin/container",
		ReverseDest: "",
	})
}

func Client() *splugin.Client {
	return globalClient
}

func StartClient(conf *config.HttpConf) {
	globalClient.Serve(conf.Addr)
}
