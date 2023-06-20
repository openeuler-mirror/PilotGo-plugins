package utils

import (
	"strings"

	"gitee.com/openeuler/PilotGo-plugins/sdk/common"
	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"openeuler.org/PilotGo/gala-ops-plugin/agentmanager"
	"openeuler.org/PilotGo/gala-ops-plugin/database"
)

// 获取业务机集群gala-gopher版本信息
func GetPackageVersion(machines []*database.AopsDepolyStatus, batch *common.Batch, cmd string) ([]*database.AopsDepolyStatus, error) {
	cmdresults, err := agentmanager.Galaops.Sdkmethod.RunCommand(batch, cmd)
	if err == nil {
		for _, result := range cmdresults {
			if result.RetCode == 1 && strings.Contains(result.Stdout, "is not installed") && result.Stderr == "" {
				logger.Error("%s not installed happened when getpackageversion: %s, %s, %s; ", "gala-gopher", result.MachineUUID, result.MachineIP, result.Stderr)
				continue
			} else if result.RetCode == 127 && result.Stdout == "" && strings.Contains(result.Stderr, "command not found") {
				logger.Error("rpm not installed when getpackageversion: %s, %s, %s", result.MachineUUID, result.MachineIP, result.Stderr)
				continue
			} else if result.RetCode == 0 && len(result.Stdout) > 0 && result.Stderr == "" {
				reader := strings.NewReader(result.Stdout)
				v, err := ReadInfo(reader, `^Version.*`)
				if err != nil && len(v) != 0 {
					logger.Error("failed to read RPM package version when getpackageversion: %s, %s, %s", result.MachineUUID, result.MachineIP, result.Stderr)
					continue
				}
				for _, m := range machines {
					if m.UUID == result.MachineUUID {
						m.Gopher_version = v
					}
				}
			} else {
				logger.Error("failed to run command: %s in %s, %s, %s when getpackageversion", "rpm -qi gala-gopher", result.MachineUUID, result.MachineIP, result.Stderr)
				continue
			}
		}
		return machines, nil
	}
	return nil, err
}
