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

func toPathConfig(pf *PathFile) PathConfig {
	return PathConfig{
		UUID:           pf.UUID,
		ConfigInfoUUID: pf.ConfigInfoUUID,
		Path:           pf.Path,
		Name:           pf.Name,
		Content:        pf.Content,
		Version:        pf.Version,
		IsActive:       pf.IsActive,
	}
}

func (pc *PathConfig) Record() error {
	pf := pc.toPathFile()
	return pf.Add()
}

func (pc *PathConfig) Load() error {
	// 加载正在使用的某配置文件
	pf, err := internal.GetPathFileByInfoUUID(pc.ConfigInfoUUID, true)
	if err != nil {
		return err
	}
	pc.UUID = pf.UUID
	pc.Path = pf.Path
	pc.Name = pf.Name
	pc.Content = pf.Content
	pc.Version = pf.Version
	pc.IsActive = pf.IsActive
	return nil
}

// 根据配置uuid获取所有配置文件
func GetPathFilesByConfigUUID(uuid string) ([]PathFile, error) {
	return internal.GetPathFilesByConfigUUID(uuid)
}

// 查看某台机器某种类型的的历史配置信息
func GetPathFilesByNode(nodeid string) ([]PathConfig, error) {
	// 查找本台机器所属的配置uuid
	config_nodes, err := internal.GetConfigNodesByNode(nodeid)
	if err != nil {
		return nil, err
	}
	var pcs []PathConfig
	for _, v := range config_nodes {
		pf, err := internal.GetPathFileByInfoUUID(v.ConfigInfoUUID, nil)
		if err != nil {
			return nil, err
		}
		pc := toPathConfig(&pf)
		pcs = append(pcs, pc)
	}
	return pcs, nil
}
