package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"openeuler.org/PilotGo/configmanage-plugin/internal"
)

/*
修改内核参数的方法：1、sysctl 命令；2、挂载于 /proc/sys/ 目录的虚拟文件系统；3、/etc/sysctl.d/ 目录中的配置文件
采用1、2方法可以修改临时配置，重启后恢复原来配置；采用1、3方法可以永久修改配置

sysctl: 配置文件，sysctl.conf文件只有一个

一般方法：	1、在/etc/ysctl.conf中修改内容，/etc/sysctl.conf是系统控制文件，主要用于运行时配置内核参数等系统信息，它的内容全部对应于/proc/sys/目录及其子目录。

	2、有以下两种方法使其生效：

（1）重启机器reboot。但是一般不推荐，因为可能机器上正在运行的程序

（2）使用如下命令刷新配置，使其立即生效。

/sbin/sysctl -p

/sbin/sysctl -w net.ipv4.route.flush=1

考虑的问题：1、目前只考虑创建、下发/etc/sysctl.conf，不考虑其他子文件的，后续可继续完善。
*/
type SysctlFile = internal.SysctlFile
type SysctlConfig struct {
	UUID           string          `json:"uuid"`
	ConfigInfoUUID string          `json:"configinfouuid"`
	Content        json.RawMessage `json:"content"`
	Version        string          `json:"version"`
	Path           string          `json:"path"`
	Name           string          `json:"name"`
	//下发改变标志位
	IsActive bool `json:"isactive"`
}

func (sysc *SysctlConfig) toSysctlFile() SysctlFile {
	return SysctlFile{
		UUID:           sysc.UUID,
		ConfigInfoUUID: sysc.ConfigInfoUUID,
		Path:           sysc.Path,
		Name:           sysc.Name,
		Content:        sysc.Content,
		Version:        fmt.Sprintf("v%s", time.Now().Format("2006-01-02-15-04-05")),
		IsActive:       sysc.IsActive,
	}
}

func (sysc *SysctlConfig) Record() error {
	//检查info的uuid是否存在
	ci, err := GetInfoByUUID(sysc.ConfigInfoUUID)
	if err != nil || ci.UUID == "" {
		return errors.New("configinfo uuid not exist")
	}

	sysf := sysc.toSysctlFile()
	return sysf.Add()
}

func (sysc *SysctlConfig) Load() error {
	// 加载正在使用的某配置文件
	sysf, err := internal.GetSysctlFileByInfoUUID(sysc.ConfigInfoUUID, true)
	if err != nil {
		return err
	}
	sysc.UUID = sysf.UUID
	sysc.Path = sysf.Path
	sysc.Name = sysf.Name
	sysc.Content = sysf.Content
	sysc.Version = sysf.Version
	sysc.IsActive = sysf.IsActive
	return nil
}
