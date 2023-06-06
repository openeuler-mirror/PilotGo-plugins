package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"openeuler.org/PilotGo/gala-ops-plugin/client"
	"openeuler.org/PilotGo/gala-ops-plugin/config"
	"openeuler.org/PilotGo/gala-ops-plugin/database"
	"openeuler.org/PilotGo/gala-ops-plugin/httphandler"
)

func main() {
	fmt.Println("hello gala-ops")

	// config.Init()

	if err := database.MysqlInit(config.Config().Mysql); err != nil {
		fmt.Println("failed to initialize database")
		os.Exit(-1)
	}

	engine := client.Client().HttpEngine
	registerHandlers(engine)
	client.StartClient(config.Config().Http)
}

func registerHandlers(engine *gin.Engine) {
	manager := engine.Group("/plugin_manage")
	{
		manager.GET("/info", httphandler.PluginInfo)
	}

	api := engine.Group("/plugin/gala-ops/api")
	{
		// 脚本执行结果接口
		// api.PUT("/run_script_result", httphandler.RunScriptResult)

		// 安装/升级/卸载gala-gopher监控终端
		api.PUT("/install_gopher", httphandler.InstallGopher)
		api.PUT("/upgrade_gopher", httphandler.UpgradeGopher)
		api.DELETE("/uninstall_gopher", httphandler.UninstallGopher)
	}
}
