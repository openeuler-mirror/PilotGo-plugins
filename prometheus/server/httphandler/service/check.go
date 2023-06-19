package service

import (
	"errors"
	"strings"

	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"gitee.com/openeuler/PilotGo-plugins/sdk/utils/command"
	"openeuler.org/PilotGo/prometheus-plugin/global"
)

// Check if prometheus is installed
func CheckPrometheus() error {
	exec := "ls /etc/prometheus/prometheus.yml /etc/prometheus/prometheus.yaml"
	_, stdout, stderr, _ := command.RunCommand(exec)
	if len(stdout) > 0 {
		logger.Debug("prometheus already installed")
		global.GlobalPrometheusYml = strings.Trim(stdout, "\n")
		return nil
	}
	return errors.New(stderr)
}
