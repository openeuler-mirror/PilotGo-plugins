package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"openeuler.org/PilotGo/configmanage-plugin/internal"
)

/*
sshd: 配置文件，sshd_config文件只有一个

一般方法：	1、在/etc/ssh/sshd_config中修改内容

	2、执行命令systemctl restart sshd重启sshd服务

考虑的问题：
*/
type SSHDFile = internal.SSHDFile
type SSHDConfig struct {
	UUID           string          `json:"uuid"`
	ConfigInfoUUID string          `json:"configinfouuid"`
	Content        json.RawMessage `json:"content"`
	Version        string          `json:"version"`
	Path           string          `json:"path"`
	Name           string          `json:"name"`
	//下发改变标志位
	IsActive bool `json:"isactive"`
}

func (sdc *SSHDConfig) toSSHDFile() SSHDFile {
	return SSHDFile{
		UUID:           sdc.UUID,
		ConfigInfoUUID: sdc.ConfigInfoUUID,
		Path:           sdc.Path,
		Name:           sdc.Name,
		Content:        sdc.Content,
		Version:        fmt.Sprintf("v%s", time.Now().Format("2006-01-02-15-04-05")),
		IsActive:       sdc.IsActive,
	}
}

func (sdc *SSHDConfig) Record() error {
	//检查info的uuid是否存在
	ci, err := GetInfoByUUID(sdc.ConfigInfoUUID)
	if err != nil || ci.UUID == "" {
		return errors.New("configinfo uuid not exist")
	}

	sdf := sdc.toSSHDFile()
	return sdf.Add()
}

func (sdc *SSHDConfig) Load() error {
	// 加载正在使用的某配置文件
	sdf, err := internal.GetSSHFileByInfoUUID(sdc.ConfigInfoUUID, true)
	if err != nil {
		return err
	}
	sdc.UUID = sdf.UUID
	sdc.Path = sdf.Path
	sdc.Name = sdf.Name
	sdc.Content = sdf.Content
	sdc.Version = sdf.Version
	sdc.IsActive = sdf.IsActive
	return nil
}
