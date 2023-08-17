package collector

import (
	"encoding/json"
	"fmt"

	"gitee.com/openeuler/PilotGo-plugin-topology-agent/utils"
	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"github.com/shirou/gopsutil/process"
	"github.com/shirou/gopsutil/v3/host"
)

type PsutilCollector struct {
	Host_0           *Host
	Processes_0      []*Process
	Netconnections_0 []*Netconnection
}

func (pc *PsutilCollector) Collect_host_data() error {
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

func (pc *PsutilCollector) Collect_process_instant_data() error {
	Echo_process_err := func(method string, err error, processid int32) {
		if err != nil {
			logger.Debug("failed to run process.%s: %d, %s", method, processid, err)
		}
	}

	processes_0, err := process.Processes()
	if err != nil {
		return fmt.Errorf("failed to run gopsutil/process.processes: %s", err)
	}

	p1 := &Process{}
	for _, p0 := range processes_0 {
		p1.Pid = p0.Pid

		p1.Ppid, err = p0.Ppid()
		Echo_process_err("ppid", err, p0.Pid)

		children, err := p0.Children()
		Echo_process_err("children", err, p0.Pid)
		for _, c := range children {
			p1.Cpid = append(p1.Cpid, c.Pid)
		}

		thread, err := p0.Threads()
		Echo_process_err("thread", err, p0.Pid)
		for k := range thread {
			p1.Tid = append(p1.Tid, k)
		}

		p1.Uids, err = p0.Uids()
		Echo_process_err("uids", err, p0.Pid)

		p1.Gids, err = p0.Gids()
		Echo_process_err("gids", err, p0.Pid)

		p1.Username, err = p0.Username()
		Echo_process_err("username", err, p0.Pid)

		p1.Status, err = p0.Status()
		Echo_process_err("status", err, p0.Pid)

		p1.CreateTime, err = p0.CreateTime()
		Echo_process_err("createtime", err, p0.Pid)

		p1.ExePath, err = p0.Exe()
		Echo_process_err("exe", err, p0.Pid)

		p1.ExeName, err = p0.Name()
		Echo_process_err("name", err, p0.Pid)

		p1.Cmdline, err = p0.Cmdline()
		Echo_process_err("cmdline", err, p0.Pid)

		p1.Cwd, err = p0.Cwd()
		Echo_process_err("cwd", err, p0.Pid)

		p1.Nice, err = p0.Nice()
		Echo_process_err("nice", err, p0.Pid)

		p1.IOnice, err = p0.IOnice()
		Echo_process_err("ionice", err, p0.Pid)

		p1.Connections, err = p0.Connections()
		Echo_process_err("connections", err, p0.Pid)

		p1.NetIOCounters, err = p0.NetIOCounters(true)
		Echo_process_err("netiocounters", err, p0.Pid)

		iocounters, err := p0.IOCounters()
		Echo_process_err("iocounters", err, p0.Pid)
		p1.IOCounters = *iocounters

		p1.OpenFiles, err = p0.OpenFiles()
		Echo_process_err("openfiles", err, p0.Pid)

		p1.NumFDs, err = p0.NumFDs()
		Echo_process_err("numfds", err, p0.Pid)

		numctxswitches, err := p0.NumCtxSwitches()
		Echo_process_err("numctxswitches", err, p0.Pid)
		p1.NumCtxSwitches = *numctxswitches

		pagefaults, err := p0.PageFaults()
		Echo_process_err("pagefaults", err, p0.Pid)
		p1.PageFaults = *pagefaults

		memoryinfo, err := p0.MemoryInfo()
		Echo_process_err("memoryinfo", err, p0.Pid)
		p1.MemoryInfo = *memoryinfo

		p1.CPUPercent, err = p0.CPUPercent()
		Echo_process_err("cpupercent", err, p0.Pid)

		memorypercent, err := p0.MemoryPercent()
		Echo_process_err("memorypercent", err, p0.Pid)
		p1.MemoryPercent = float64(memorypercent)

		pc.Processes_0 = append(pc.Processes_0, p1)
	}

	return nil
}
