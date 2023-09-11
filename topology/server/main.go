package main

import (
	"fmt"

	"gitee.com/openeuler/PilotGo-plugin-topology-server/agentmanager"
	"gitee.com/openeuler/PilotGo-plugin-topology-server/handler"
)

func main() {
	fmt.Println("hello topology")

	/*
		init plugin client
	*/
	agentmanager.Topo.InitPluginClient()

	/*
		init logger
	*/
	agentmanager.Topo.InitLogger()

	/*
		init arangodb
		TODO:
	*/
	agentmanager.Topo.InitArangodb()

	/*
		init machine agent list
		TODO: 实时更新machine agent、topo agent的状态
	*/
	agentmanager.Topo.InitMachineList()

	/*
		init topo agent status
		TODO:
	*/

	/*
		init web server
	*/
	handler.InitWebServer()
}
