package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"openeuler.org/PilotGo/configmanage-plugin/internal"
)

/*
ssh: 配置文件，ssh_config文件只有一个

一般方法：	1、在/etc/ssh/ssh_config中修改内容

	2、执行命令systemctl restart ssh重启ssh服务

考虑的问题：
*/
type SSHFile = internal.SSHFile
type SSHConfig struct {
	UUID           string          `json:"uuid"`
	ConfigInfoUUID string          `json:"configinfouuid"`
	Content        json.RawMessage `json:"content"`
	Version        string          `json:"version"`
	Path           string          `json:"path"`
	Name           string          `json:"name"`
	//下发改变标志位
	IsActive bool `json:"isactive"`
}

func (sc *SSHConfig) toSSHFile() SSHFile {
	return SSHFile{
		UUID:           sc.UUID,
		ConfigInfoUUID: sc.ConfigInfoUUID,
		Path:           sc.Path,
		Name:           sc.Name,
		Content:        sc.Content,
		Version:        fmt.Sprintf("v%s", time.Now().Format("2006-01-02-15-04-05")),
		IsActive:       sc.IsActive,
	}
}

func (sc *SSHConfig) Record() error {
	//检查info的uuid是否存在
	ci, err := GetInfoByUUID(sc.ConfigInfoUUID)
	if err != nil || ci.UUID == "" {
		return errors.New("configinfo uuid not exist")
	}

	sf := sc.toSSHFile()
	return sf.Add()
}
func (sc *SSHConfig) Load() error {
	// 加载正在使用的某配置文件
	sf, err := internal.GetSSHFileByInfoUUID(sc.ConfigInfoUUID, true)
	if err != nil {
		return err
	}
	sc.UUID = sf.UUID
	sc.Path = sf.Path
	sc.Name = sf.Name
	sc.Content = sf.Content
	sc.Version = sf.Version
	sc.IsActive = sf.IsActive
	return nil
}

// TODO:
func (rc *SSHConfig) Apply() (json.RawMessage, error) {
	return nil, errors.New("failed to apply SSHConfig")
}

// TODO:
func (rc *SSHConfig) Collect() error {
	return nil
}