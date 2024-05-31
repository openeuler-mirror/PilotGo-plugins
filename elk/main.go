package main

import (
	"gitee.com/openeuler/PilotGo-plugin-elk/conf"
	"gitee.com/openeuler/PilotGo-plugin-elk/db"
	"gitee.com/openeuler/PilotGo-plugin-elk/errormanager"
	"gitee.com/openeuler/PilotGo-plugin-elk/handler"
	"gitee.com/openeuler/PilotGo-plugin-elk/kibanaClient"
	"gitee.com/openeuler/PilotGo-plugin-elk/logger"
	"gitee.com/openeuler/PilotGo-plugin-elk/pluginclient"
	"gitee.com/openeuler/PilotGo-plugin-elk/signal"
)

func main() {
	/*
		init config
	*/
	conf.InitConfig()

	/*
		init plugin client
	*/
	pluginclient.InitPluginClient()

	/*
		init error control
	*/
	errormanager.InitErrorManager()

	/*
		init web server
	*/
	handler.InitWebServer()

	/*
		init logger
	*/
	logger.InitLogger()

	/*
		init database
		neo4j mysql redis prometheus
	*/
	db.InitDB()

	/*
		init kibana client
	*/
	kibanaClient.InitKibanaClient()

	/*
		终止进程信号监听
	*/
	signal.SignalMonitoring()
}
