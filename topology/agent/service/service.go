package service

import (
	"fmt"

	"gitee.com/openeuler/PilotGo-plugin-topology-agent/collector"
	"gitee.com/openeuler/PilotGo-plugin-topology-agent/conf"
)

func DataCollectorService() (collector.Data_collector, error) {
	datasource := conf.Config().Topo.Datasource
	switch datasource {
	case "gopsutil":
		gops := &collector.PsutilCollector{}
		err := gops.Collect_host_data()
		if err != nil {
			return nil, fmt.Errorf("(datacollectorservice->Collect_host_data: %s)", err)
		}

		err = gops.Collect_netconnection_all_data()
		if err != nil {
			return nil, fmt.Errorf("(datacollectorservice->Collect_netconnection_all_data: %s)", err)
		}

		err = gops.Collect_process_instant_data()
		if err != nil {
			return nil, fmt.Errorf("(datacollectorservice->Collect_process_instant_data: %s)", err)
		}

		return gops, nil
	case "ebpf":

	}

	return nil, fmt.Errorf("wrong data source")
}
