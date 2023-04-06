package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/gala-ops-plugin/config"
	"openeuler.org/PilotGo/gala-ops-plugin/httphandler"
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

	registerHandlers(client.HttpEngine)

}

func registerHandlers(engine *gin.Engine) {
	api := engine.Group("/plugin/gala-ops/api")
	{
		// 安装/升级/卸载gala-gopher监控终端
		api.PUT("/install_gopher", httphandler.InstallGopher)
		api.PUT("/upgrade_gopher", httphandler.UpgradeGopher)
		api.DELETE("/uninstall_gopher", httphandler.UninstallGopher)
	}
}
