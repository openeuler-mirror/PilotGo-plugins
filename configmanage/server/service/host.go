package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"openeuler.org/PilotGo/configmanage-plugin/internal"
)

/*
host: 配置文件

一般方法：	1、在/etc/hosts中修改内容

	2、执行命令systemctl restart NetworkManager重启网络服务

考虑的问题：
*/
type HostFile = internal.HostFile

type HostConfig struct {
	UUID           string          `json:"uuid"`
	ConfigInfoUUID string          `json:"configinfouuid"`
	Content        json.RawMessage `json:"content"`
	Version        string          `json:"version"`
	Path           string          `json:"path"`
	Name           string          `json:"name"`
	//下发改变标志位
	IsActive bool `json:"isactive"`
}

func (hc *HostConfig) toHostFile() HostFile {
	return HostFile{
		UUID:           hc.UUID,
		ConfigInfoUUID: hc.ConfigInfoUUID,
		Path:           hc.Path,
		Name:           hc.Name,
		Content:        hc.Content,
		Version:        fmt.Sprintf("v%s", time.Now().Format("2006-01-02-15-04-05")),
		IsActive:       hc.IsActive,
	}
}

func (hc *HostConfig) Record() error {
	//检查info的uuid是否存在
	ci, err := GetInfoByUUID(hc.ConfigInfoUUID)
	if err != nil || ci.UUID == "" {
		return errors.New("configinfo uuid not exist")
	}

	hf := hc.toHostFile()
	return hf.Add()
}
