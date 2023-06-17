package service

import (
	"errors"
	"fmt"

	"gitee.com/openeuler/PilotGo-plugins/sdk/utils/command"
	"openeuler.org/PilotGo/prometheus-plugin/global"
)

// Check if prometheus is installed
func CheckPrometheus() error {
	exec := "ls /etc/prometheus/prometheus.yml /etc/prometheus/prometheus.yaml"
	_, stdout, stderr, err := command.RunCommand(exec)
	if len(stdout) > 0 {
		fmt.Println("prometheus already installed")
		global.GlobalPrometheusYml = stdout
		return nil
	}
	return errors.New(stderr + err.Error())
}
