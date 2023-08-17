package collector

import (
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

type Data_collector interface {
	Collect_process_instant_data() error
	Collect_host_data() error
	Collect_netconnection_all_data() error
}

type Host struct {
	Hostname             string `json:"hostname"`
	Uptime               uint64 `json:"uptime"`
	BootTime             uint64 `json:"bootTime"`
	Procs                uint64 `json:"procs"`           // number of processes
	OS                   string `json:"os"`              // ex: freebsd, linux
	Platform             string `json:"platform"`        // ex: ubuntu, linuxmint
	PlatformFamily       string `json:"platformFamily"`  // ex: debian, rhel
	PlatformVersion      string `json:"platformVersion"` // version of the complete OS
	KernelVersion        string `json:"kernelVersion"`   // version of the OS kernel (if available)
	KernelArch           string `json:"kernelArch"`      // native cpu architecture queried at runtime, as returned by `uname -m` or empty string in case of error
	VirtualizationSystem string `json:"virtualizationSystem"`
	VirtualizationRole   string `json:"virtualizationRole"` // guest or host
	MachineUUID          string `json:"MachineUUID"`        // ex: pilotgo agent uuid
}

type Process struct {
	Pid  int32   `json:"pid"`
	Ppid int32   `json:"ppid"`
	Cpid []int32 `json:"cpid"`
	Tid  []int32 `json:"tid"`
	Uids []int32 `json:"uids"`
	Gids []int32 `json:"gids"`

	Username   string `json:"username"`
	Status     string `json:"status"`
	CreateTime int64  `json:"createtime"`
	ExePath    string `json:"exepath"`
	ExeName    string `json:"exename"`
	Cmdline    string `json:"cmdline"`
	Cwd        string `json:"cwd"`

	Nice   int32 `json:"nice"`
	IOnice int32 `json:"ionice"`

	Connections   []net.ConnectionStat `json:"connections"`
	NetIOCounters []net.IOCountersStat `json:"netiocounters"`

	IOCounters process.IOCountersStat `json:"iocounters"`

	OpenFiles []process.OpenFilesStat `json:"openfiles"`
	NumFDs    int32                   `json:"numfds"`

	NumCtxSwitches process.NumCtxSwitchesStat `json:"numctxswitches"`
	PageFaults     process.PageFaultsStat     `json:"pagefaults"`
	MemoryInfo     process.MemoryInfoStat     `json:"memoryinfo"`
	CPUPercent     float64                    `json:"cpupercent"`
	MemoryPercent  float64                    `json:"memorypercent"`
}

type Thread struct {
}

type Netconnection struct {
	Fd     uint32            `json:"fd"`
	Family uint32            `json:"family"`
	Type   uint32            `json:"type"`
	Laddr  map[string]string `json:"localaddr"`
	Raddr  map[string]string `json:"remoteaddr"`
	Status string            `json:"status"`
	Uids   []int32           `json:"uids"`
	Pid    int32             `json:"pid"`
}

type NetIOcounters struct {
}

type Resource struct {
}

type Container struct {
}
