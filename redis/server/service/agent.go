package service

import (
	"gitee.com/openeuler/PilotGo-plugins/sdk/common"

	"openeuler.org/PilotGo/redis-plugin/global"
	"openeuler.org/PilotGo/redis-plugin/plugin"
)

func Install(param *common.Batch) ([]interface{}, error) {
	cmd := "yum install -y redis_exporter && systemctl start redis_exporter"

	cmdResults, err := global.GlobalClient.RunScript(param, cmd)
	if err != nil {
		return nil, err
	}
	ret := []interface{}{}
	monitorTargets := []string{}
	for _, result := range cmdResults {
		d := struct {
			MachineUUID   string
			MachineIP     string
			InstallStatus string
			Error         string
		}{
			MachineUUID:   result.MachineUUID,
			MachineIP:     result.MachineIP,
			InstallStatus: "ok",
			Error:         "",
		}

		if result.RetCode != 0 {
			d.InstallStatus = "error"
			d.Error = result.Stderr
		} else {
			// TODO: add redis exporter to prometheus monitor target here
			// default exporter port :9121
			monitorTargets = append(monitorTargets, result.MachineIP+":9121")
		}

		ret = append(ret, d)
	}

	err = plugin.MonitorTargets(monitorTargets)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
