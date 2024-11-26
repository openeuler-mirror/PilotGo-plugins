package service

import (
	"encoding/json"
	"fmt"
	"time"

	"openeuler.org/PilotGo/configmanage-plugin/internal"
)

/*
path: 环境变量文件只有一个

一般方法：	1、在/etc/profile中修改内容

考虑的问题：
*/

type PathFile = internal.PathFile

type PathConfig struct {
	UUID           string          `json:"uuid"`
	ConfigInfoUUID string          `json:"configinfouuid"`
	Content        json.RawMessage `json:"content"`
	Version        string          `json:"version"`
	Path           string          `json:"path"`
	Name           string          `json:"name"`
	//下发改变标志位
	IsActive bool `json:"isactive"`
}

func (pc *PathConfig) toPathFile() PathFile {
	return PathFile{
		UUID:           pc.UUID,
		ConfigInfoUUID: pc.ConfigInfoUUID,
		Path:           pc.Path,
		Name:           pc.Name,
		Content:        pc.Content,
		Version:        fmt.Sprintf("v%s", time.Now().Format("2006-01-02-15-04-05")),
		IsActive:       pc.IsActive,
		IsFromHost:     false,
	}
}

func (pc *PathConfig) Record() error {
	pf := pc.toPathFile()
	return pf.Add()
}
