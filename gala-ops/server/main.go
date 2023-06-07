package main

import (
	"fmt"
	"os"

	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"gitee.com/openeuler/PilotGo-plugins/sdk/plugin/client"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/gala-ops-plugin/config"
	"openeuler.org/PilotGo/gala-ops-plugin/database"
	"openeuler.org/PilotGo/gala-ops-plugin/httphandler"
)

const Version = "0.0.1"

var GlobalClient *client.Client

var PluginInfo = &client.PluginInfo{
	Name:        "gala-ops",
	Version:     Version,
	Description: "gala-ops智能运维工具",
	Author:      "guozhengxin",
	Email:       "guozhengxin@kylinos.cn",
	Url:         "http://192.168.75.100:9999/plugin/gala-ops",
	// ReverseDest: "http://192.168.48.163:3000/",
}

func main() {
	fmt.Println("hello gala-ops")

	if err := database.MysqlInit(config.Config().Mysql); err != nil {
		fmt.Println("failed to initialize database")
		os.Exit(-1)
	}

	InitLogger()

	server := gin.Default()

	GlobalClient := client.DefaultClient(PluginInfo)
	// 临时给server赋值
	GlobalClient.Server = "http://192.168.75.100:8888"
	GlobalClient.RegisterHandlers(server)
	InitRouter(server)

	if err := server.Run(config.Config().Http.Addr); err != nil {
		logger.Fatal("failed to run server")
	}
}

func InitLogger() {
	if err := logger.Init(config.Config().Logopts); err != nil {
		fmt.Printf("logger init failed, please check the config file: %s", err)
		os.Exit(-1)
	}
}

func InitRouter(router *gin.Engine) {
	api := router.Group("/plugin/gala-ops/api")
	{
		// 脚本执行结果接口
		// api.PUT("/run_script_result", httphandler.RunScriptResult)

		// 安装/升级/卸载gala-gopher监控终端
		api.PUT("/install_gopher", func(ctx *gin.Context) {
			httphandler.InstallGopher(ctx, GlobalClient)
		})
		api.PUT("/upgrade_gopher", func(ctx *gin.Context) {
			httphandler.UpgradeGopher(ctx, GlobalClient)
		})
		api.DELETE("/uninstall_gopher", func(ctx *gin.Context) {
			httphandler.UninstallGopher(ctx, GlobalClient)
		})
	}
}
