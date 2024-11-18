package service

import (
	"encoding/json"
	"fmt"
	"time"

	"openeuler.org/PilotGo/configmanage-plugin/internal"
)

/*
DNS客户端配置: 配置文件，resolv.conf文件只有一个

一般方法：	1、在/etc/resolv.conf中修改内容

	2、执行命令systemctl restart networking重启网络服务

考虑的问题：
*/
type DNSFile = internal.DNSFile

type DNSConfig struct {
	UUID           string          `json:"uuid"`
	ConfigInfoUUID string          `json:"configinfouuid"`
	Content        json.RawMessage `json:"content"`
	Version        string          `json:"version"`
	Path           string          `json:"path"`
	Name           string          `json:"name"`
	//下发改变标志位
	IsActive bool `json:"isactive"`
}

func (dc *DNSConfig) toDNSFile() DNSFile {
	return DNSFile{
		UUID:           dc.UUID,
		ConfigInfoUUID: dc.ConfigInfoUUID,
		Path:           dc.Path,
		Name:           dc.Name,
		Content:        dc.Content,
		Version:        fmt.Sprintf("v%s", time.Now().Format("2006-01-02-15-04-05")),
		IsActive:       dc.IsActive,
		IsFromHost:     false,
	}
}

func toDNSConfig(df *DNSFile) DNSConfig {
	return DNSConfig{
		UUID:           df.UUID,
		ConfigInfoUUID: df.ConfigInfoUUID,
		Path:           df.Path,
		Name:           df.Name,
		Content:        df.Content,
		Version:        df.Version,
		IsActive:       df.IsActive,
	}
}

func (dc *DNSConfig) Record() error {
	df := dc.toDNSFile()
	return df.Add()
}

func (dc *DNSConfig) Load() error {
	// 加载正在使用的某配置文件
	df, err := internal.GetDNSFileByInfoUUID(dc.ConfigInfoUUID, true)
	if err != nil {
		return err
	}
	dc.UUID = df.UUID
	dc.Path = df.Path
	dc.Name = df.Name
	dc.Content = df.Content
	dc.Version = df.Version
	dc.IsActive = df.IsActive
	return nil
}

func GetDNSFileByInfoUUID(uuid string, isindex interface{}) (DNSFile, error) {
	return internal.GetDNSFileByInfoUUID(uuid, isindex)
}

// 根据配置uuid获取所有配置文件
func GetDNSFilesByConfigUUID(uuid string) ([]DNSFile, error) {
	return internal.GetDNSFilesByConfigUUID(uuid)
}

// 查看某台机器某种类型的的历史配置信息
func GetDNSFilesByNode(nodeid string) ([]DNSConfig, error) {
	// 查找本台机器所属的配置uuid
	config_nodes, err := internal.GetConfigNodesByNode(nodeid)
	if err != nil {
		return nil, err
	}
	var dcs []DNSConfig
	for _, v := range config_nodes {
		df, err := internal.GetDNSFileByInfoUUID(v.ConfigInfoUUID, nil)
		if err != nil {
			return nil, err
		}
		dc := toDNSConfig(&df)
		dcs = append(dcs, dc)
	}
	return dcs, nil
}
