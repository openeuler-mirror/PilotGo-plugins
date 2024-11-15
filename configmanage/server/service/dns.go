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
func (dc *DNSConfig) Record() error {
	df := dc.toDNSFile()
	return df.Add()
}
