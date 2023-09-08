package utils

import (
	"reflect"
	"strconv"
	"strings"

	"gitee.com/openeuler/PilotGo-plugin-topology-server/meta"
)

func StructToMap(obj interface{}) map[string]string {
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}

	if objValue.Kind() != reflect.Struct {
		return nil
	}

	objType := objValue.Type()
	fieldCount := objType.NumField()

	m := make(map[string]string)
	for i := 0; i < fieldCount; i++ {
		field := objType.Field(i)
		fieldValue := objValue.Field(i)

		m[field.Name] = fieldValue.Interface().(string)
	}

	return m
}

func HostToMap(host *meta.Host, a_i_map *map[string][]string) *map[string]string {
	host_metrics := StructToMap(host)
	interfaces_string := []string{}

	for key, value := range *a_i_map {
		interfaces_string = append(interfaces_string, key+":"+strings.Join(value, "**"))
	}

	host_metrics["interfaces"] = strings.Join(interfaces_string, "***")

	return &host_metrics
}

func ProcessToMap(process *meta.Process) *map[string]string {
	uids_string := []string{}
	for _, u := range process.Uids {
		uids_string = append(uids_string, strconv.Itoa(int(u)))
	}

	gids_string := []string{}
	for _, g := range process.Gids {
		gids_string = append(gids_string, strconv.Itoa(int(g)))
	}

	openfiles_string := []string{}
	for _, of := range process.OpenFiles {
		openfiles_string = append(openfiles_string, strconv.Itoa(int(of.Fd))+":"+of.Path)
	}

	cpid_string := []string{}
	for _, cid := range process.Cpid {
		cpid_string = append(cpid_string, strconv.Itoa(int(cid)))
	}

	return &map[string]string{
		"Pid":                         strconv.Itoa(int(process.Pid)),
		"Ppid":                        strconv.Itoa(int(process.Ppid)),
		"Cpid":                        strings.Join(cpid_string, "-"),
		"Uids":                        strings.Join(uids_string, "-"),
		"Gids":                        strings.Join(gids_string, "-"),
		"Status":                      process.Status,
		"CreateTime":                  strconv.Itoa(int(process.CreateTime)),
		"Cwd":                         process.Cwd,
		"ExePath":                     process.ExePath,
		"Cmdline":                     process.Cmdline,
		"Nice":                        strconv.Itoa(int(process.Nice)),
		"IOnice":                      strconv.Itoa(int(process.IOnice)),
		"DISK-rc":                     strconv.Itoa(int(process.IOCounters.ReadCount)),
		"DISK-rb":                     strconv.Itoa(int(process.IOCounters.ReadBytes)),
		"DISK-wc":                     strconv.Itoa(int(process.IOCounters.WriteCount)),
		"DISK-wb":                     strconv.Itoa(int(process.IOCounters.WriteBytes)),
		"fd":                          strings.Join(openfiles_string, "-"),
		"NumCtxSwitches-v":            strconv.Itoa(int(process.NumCtxSwitches.Voluntary)),
		"NumCtxSwitches-inv":          strconv.Itoa(int(process.NumCtxSwitches.Involuntary)),
		"PageFaults-MinorFaults":      strconv.Itoa(int(process.PageFaults.MinorFaults)),
		"PageFaults-MajorFaults":      strconv.Itoa(int(process.PageFaults.MajorFaults)),
		"PageFaults-ChildMinorFaults": strconv.Itoa(int(process.PageFaults.ChildMinorFaults)),
		"PageFaults-ChildMajorFaults": strconv.Itoa(int(process.PageFaults.ChildMajorFaults)),
		"CPUPercent":                  strconv.FormatFloat(process.CPUPercent, 'f', -1, 64),
		"MemoryPercent":               strconv.FormatFloat(process.MemoryPercent, 'f', -1, 64),
		"MemoryInfo":                  process.MemoryInfo.String(),
	}
}

func ThreadToMap(thread *meta.Thread) *map[string]string {
	return &map[string]string{
		"Tid":       strconv.Itoa(int(thread.Tid)),
		"Tgid":      strconv.Itoa(int(thread.Tgid)),
		"CPU":       thread.CPU,
		"User":      strconv.FormatFloat(thread.User, 'f', -1, 64),
		"System":    strconv.FormatFloat(thread.System, 'f', -1, 64),
		"Idle":      strconv.FormatFloat(thread.Idle, 'f', -1, 64),
		"Nice":      strconv.FormatFloat(thread.Nice, 'f', -1, 64),
		"Iowait":    strconv.FormatFloat(thread.Iowait, 'f', -1, 64),
		"Irq":       strconv.FormatFloat(thread.Irq, 'f', -1, 64),
		"Softirq":   strconv.FormatFloat(thread.Softirq, 'f', -1, 64),
		"Steal":     strconv.FormatFloat(thread.Steal, 'f', -1, 64),
		"Guest":     strconv.FormatFloat(thread.Guest, 'f', -1, 64),
		"GuestNice": strconv.FormatFloat(thread.GuestNice, 'f', -1, 64),
	}
}

// net节点的metrics字段 临时定义
func NetToMap(net *meta.Netconnection) *map[string]string {
	return &map[string]string{
		"Fd": strconv.Itoa(int(net.Fd)),
		"Family": strconv.Itoa(int(net.Family)),
		"Type": strconv.Itoa(int(net.Type)),
		"Laddr": net.Laddr,
		"Raddr": net.Raddr,  
		"Status": net.Status, 
		"Uids": strconv.Itoa(int(net.Uids[])),
		"Pid": strconv.Itoa(int(net.Pid)),  
	}
}
// func NetToMap(net *net.IOCountersStat, a_i_map *map[string][]string) *map[string]string {
// 	addrs := []string{}
// 	for key, value := range *a_i_map {
// 		if net.Name == key {
// 			addrs = value
// 		}
// 	}

// 	return &map[string]string{
// 		"Name":        net.Name,
// 		"addrs":       addrs[0],
// 		"BytesSent":   strconv.Itoa(int(net.BytesSent)),
// 		"BytesRecv":   strconv.Itoa(int(net.BytesRecv)),
// 		"PacketsSent": strconv.Itoa(int(net.PacketsSent)),
// 		"PacketsRecv": strconv.Itoa(int(net.PacketsRecv)),
// 		"Errin":       strconv.Itoa(int(net.Errin)),
// 		"Errout":      strconv.Itoa(int(net.Errout)),
// 		"Dropin":      strconv.Itoa(int(net.Dropin)),
// 		"Dropout":     strconv.Itoa(int(net.Dropout)),
// 		"Fifoin":      strconv.Itoa(int(net.Fifoin)),
// 		"Fifoout":     strconv.Itoa(int(net.Fifoout)),
// 	}
// }
