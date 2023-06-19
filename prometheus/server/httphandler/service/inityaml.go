package service

import (
	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"gitee.com/openeuler/PilotGo-plugins/sdk/utils/command"
	"openeuler.org/PilotGo/prometheus-plugin/global"
)

func InitPrometheusYML(httpaddr string) error {
	if err := backup(); err != nil {
		return err
	}

	if err := initYML(httpaddr); err != nil {
		return err
	}

	logger.Debug("prometheus yml init success")
	return nil
}

func backup() error {
	cmd := "cp " + global.GlobalPrometheusYml + " " + global.GlobalPrometheusYml + ".bak"
	exitcode, _, stderr, err := command.RunCommand(cmd)
	if exitcode == 0 && stderr == "" && err == nil {
		return nil
	}
	return err
}

func initYML(httaddr string) error {
	cmd := "sh " + global.GlobalPrometheusYmlInit + " " + httaddr + " " + global.GlobalPrometheusYml
	exitcode, _, stderr, err := command.RunCommand(cmd)
	if exitcode == 0 && stderr == "" && err == nil {
		return nil
	}
	return err
}
