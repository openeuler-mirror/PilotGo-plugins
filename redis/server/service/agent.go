package service

import (
	"gitee.com/openeuler/PilotGo-plugins/sdk/common"
	"gitee.com/openeuler/PilotGo-plugins/sdk/plugin/client"

	"openeuler.org/PilotGo/redis-plugin/global"
	"openeuler.org/PilotGo/redis-plugin/plugin"
)

func FormatData(cmdResults []*client.CmdResult) ([]string, []interface{}, error) {
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
			// TODO: add or delete redis exporter to prometheus monitor target here
			// default exporter port :9121
			monitorTargets = append(monitorTargets, result.MachineIP+":9121")
		}

		ret = append(ret, d)
	}
	return monitorTargets, ret, nil
}

func Install(param *common.Batch) ([]interface{}, error) {
	cmd := "yum install -y redis_exporter && systemctl start redis_exporter"

	cmdResults, err := global.GlobalClient.RunScript(param, cmd)
	if err != nil {
		return nil, err
	}

	monitorTargets, ret, err := FormatData(cmdResults)
	if err != nil {
		return nil, err
	}
	err = plugin.MonitorTargets(monitorTargets)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func UnInstall(param *common.Batch) ([]interface{}, error) {
	cmd := "systemctl stop redis_exporter && yum autoremove -y redis_exporter"
	cmdResults, err := global.GlobalClient.RunScript(param, cmd)
	if err != nil {
		return nil, err
	}
	monitorTargets, ret, err := FormatData(cmdResults)
	if err != nil {
		return nil, err
	}
	err = plugin.MonitorTargets(monitorTargets)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func Restart(param *common.Batch) error {
	cmd := "systemctl restart redis_exporter && systemctl status redis_exporter"
	cmdResults, err := global.GlobalClient.RunScript(param, cmd)
	if err != nil {
		return err
	}
	_, _, err = FormatData(cmdResults)
	return nil
}

func Stop(param *common.Batch) error {
	cmd := "systemctl stop redis_exporter && systemctl status redis_exporter"
	cmdResults, err := global.GlobalClient.RunScript(param, cmd)
	if err != nil {
		return err
	}
	_, _, err = FormatData(cmdResults)
	return nil
}
