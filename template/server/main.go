package main

import (
	"gitee.com/openeuler/PilotGo-plugin-template/conf"
	"gitee.com/openeuler/PilotGo-plugin-template/handler"
	"gitee.com/openeuler/PilotGo-plugin-template/logger"
	"gitee.com/openeuler/PilotGo-plugin-template/pluginclient"
	"gitee.com/openeuler/PilotGo-plugin-template/signal"
)

func main() {
	/*
		init config
	*/
	conf.InitConfig()

	/*
		init logger
	*/
	logger.InitLogger()

	/*
		init plugin client
	*/
	pluginclient.InitPluginClient()

	/*
		init web server
	*/
	handler.InitWebServer()

	/*
		业务模块
	*/

	/*
		终止进程信号监听
	*/
	signal.SignalMonitoring()
}
