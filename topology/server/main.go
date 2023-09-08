package main

import (
	"fmt"

	"gitee.com/openeuler/PilotGo-plugin-topology-server/agentmanager"
	"gitee.com/openeuler/PilotGo-plugin-topology-server/collector"
	"github.com/pkg/errors"
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

	/*
		init machine agent list
		TODO: 实时更新machine agent、topo agent的状态
	*/
	agentmanager.Topo.InitMachineList()

	/*
		init topo agent status
		TODO:
	*/

	// ttcode
	datacollector := collector.CreateDataCollector()
	errorlist := datacollector.Collect_instant_data()
	if len(errorlist) != 0 {
		for _, err := range errorlist {
			err = errors.Wrap(err, "**3")
			// errors.EORE(err)
			fmt.Printf("%+v\n", err)
		}
	}
	agentmanager.Topo.AgentMap.Range(
		func(key, value any) bool {
			agent := value.(*agentmanager.Agent_m)
			fmt.Printf("\033[32m%s\033[0m: \n", agent.UUID)
			for _, net := range agent.Netconnections_2 {
				fmt.Printf("\t%+v\n", net)
			}
			return true
		},
	)

	/*
		init web server
	*/
	agentmanager.Topo.InitWebServer()
}
