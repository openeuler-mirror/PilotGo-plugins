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
	"openeuler.org/PilotGo/gala-ops-plugin/router"
)

const Version = "0.0.1"

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
		os.Exit(1)
	}

	InitLogger()

	engine := gin.Default()

	PluginClient := client.DefaultClient(PluginInfo)
	// 临时给server赋值
	PluginClient.Server = "http://192.168.75.100:8888"
	httphandler.Galaops = &httphandler.Opsclient{
		Sdkmethod:   PluginClient,
		PromePlugin: nil,
	}

	// 临时自定义获取prometheus地址方式
	promeplugin, err := httphandler.Galaops.Getplugininfo(PluginClient.Server, "Prometheus")
	if err != nil {
		logger.Error(err.Error())
	}
	httphandler.Galaops.PromePlugin = promeplugin

	// 检查prometheus插件是否在运行
	promepluginstatus, _ := httphandler.Galaops.CheckPrometheusPlugin()
	if !promepluginstatus {
		logger.Error("prometheus plugin is not running")
	}

	// 向prometheus插件发送可视化插件json模板    TODO: prometheus plugin 注册接收jsonmode的路由
	respbody, retcode, err := httphandler.Galaops.SendJsonMode("/abc")
	if err != nil || retcode != 201 {
		logger.Error("failed to send jsonmode to prometheus plugin: ", respbody, retcode, err)
	}

	// 设置router
	httphandler.Galaops.Sdkmethod.RegisterHandlers(engine)
	router.InitRouter(engine)
	if err := engine.Run(config.Config().Http.Addr); err != nil {
		logger.Fatal("failed to run server")
	}
}

func InitLogger() {
	if err := logger.Init(config.Config().Logopts); err != nil {
		fmt.Printf("logger init failed, please check the config file: %s", err)
		os.Exit(1)
	}
}
