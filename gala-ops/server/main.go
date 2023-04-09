package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"openeuler.org/PilotGo/gala-ops-plugin/httphandler"

	"openeuler.org/PilotGo/gala-ops-plugin/client"
	"openeuler.org/PilotGo/gala-ops-plugin/config"
)

func main() {
	fmt.Println("hello gala-ops")

	config.Init()

	engine := client.Client().HttpEngine
	registerHandlers(engine)
	client.StartClient(config.Config().Http)
}

func registerHandlers(engine *gin.Engine) {
	api := engine.Group("/plugin/gala-ops/api")
	{
		// 脚本执行结果接口
		api.PUT("/run_script_result", httphandler.RunScriptResult)

		// 安装/升级/卸载gala-gopher监控终端
		api.PUT("/install_gopher", httphandler.InstallGopher)
		api.PUT("/upgrade_gopher", httphandler.UpgradeGopher)
		api.DELETE("/uninstall_gopher", httphandler.UninstallGopher)
	}
}
