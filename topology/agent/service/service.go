package service

import (
	"fmt"

	"gitee.com/openeuler/PilotGo-plugin-topology-agent/collector"
	"gitee.com/openeuler/PilotGo-plugin-topology-agent/conf"
	"gitee.com/openeuler/PilotGo-plugin-topology-agent/utils"
	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
)

func DataCollectorService() (utils.Data_collector, error) {
	datasource := conf.Config().Topo.Datasource
	switch datasource {
	case "gopsutil":
		gops := &collector.PsutilCollector{}
		err := gops.Collect_host_data()
		if err != nil {
			filepath, line, funcname := utils.CallerInfo(err)
			logger.Error("file: %s, line: %d, func: %s, err: %s\n", filepath, line, funcname, err.Error())
			return nil, fmt.Errorf("file: %s, line: %d, func: %s, err -> %s", filepath, line, funcname, err.Error())
		}

		err = gops.Collect_netconnection_all_data()
		if err != nil {
			filepath, line, funcname := utils.CallerInfo(err)
			logger.Error("file: %s, line: %d, func: %s, err: %s\n", filepath, line, funcname, err.Error())
			return nil, fmt.Errorf("file: %s, line: %d, func: %s, err -> %s", filepath, line, funcname, err.Error())
		}

		err = gops.Collect_process_instant_data()
		if err != nil {
			filepath, line, funcname := utils.CallerInfo(err)
			logger.Error("file: %s, line: %d, func: %s, err: %s\n", filepath, line, funcname, err.Error())
			return nil, fmt.Errorf("file: %s, line: %d, func: %s, err -> %s", filepath, line, funcname, err.Error())
		}

		err = gops.Collect_addrInterfaceMap_data()
		if err != nil {
			filepath, line, funcname := utils.CallerInfo(err)
			logger.Error("file: %s, line: %d, func: %s, err: %s\n", filepath, line, funcname, err.Error())
			return nil, fmt.Errorf("file: %s, line: %d, func: %s, err -> %s", filepath, line, funcname, err.Error())
		}

		return gops, nil
	case "ebpf":

	}

	return nil, fmt.Errorf("wrong data source")
}
