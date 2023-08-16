package collector

import (
	"encoding/json"
	"fmt"

	"gitee.com/openeuler/PilotGo-plugin-topology-agent/utils"
	"github.com/shirou/gopsutil/v3/host"
)

type PsutilCollector struct {
	Host_0           *Host
	Processes_0      []*Process
	Netconnections_0 []*Netconnection
}

func (pc *PsutilCollector) Get_host_info() error {
	hostinit, err := host.Info()
	if err != nil {
		return fmt.Errorf("failed to get hostinfo by using gopsutil/v3/host.Info: %s", err)
	}

	m_u_bytes, err := utils.FileReadBytes(utils.Agentuuid_filepath)
	if err != nil {
		return fmt.Errorf("failed to get agent machineuuid: %s", err)
	}
	type machineuuid struct {
		Agentuuid string `json:"agent_uuid"`
	}
	m_u_struct := &machineuuid{}
	json.Unmarshal(m_u_bytes, m_u_struct)

	pc.Host_0 = &Host{
		Hostname:             hostinit.Hostname,
		Uptime:               hostinit.Uptime,
		BootTime:             hostinit.BootTime,
		Procs:                hostinit.Procs,
		OS:                   hostinit.OS,
		Platform:             hostinit.Platform,
		PlatformFamily:       hostinit.PlatformFamily,
		PlatformVersion:      hostinit.PlatformVersion,
		KernelVersion:        hostinit.KernelVersion,
		KernelArch:           hostinit.KernelArch,
		VirtualizationSystem: hostinit.VirtualizationSystem,
		VirtualizationRole:   hostinit.VirtualizationRole,
		MachineUUID:          m_u_struct.Agentuuid,
	}

	return nil
}
