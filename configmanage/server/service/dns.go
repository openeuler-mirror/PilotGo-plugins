package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"gitee.com/openeuler/PilotGo/sdk/common"
	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
	"gitee.com/openeuler/PilotGo/sdk/utils/httputils"
	"openeuler.org/PilotGo/configmanage-plugin/global"
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

func (dc *DNSConfig) Apply() ([]NodeResult, error) {
	// 从数据库获取下发的信息
	df, err := internal.GetDNSFileByUUID(dc.UUID)
	if err != nil {
		return nil, err
	}
	if df.ConfigInfoUUID != dc.ConfigInfoUUID || df.UUID != dc.UUID {
		return nil, errors.New("数据库不存在此配置")
	}

	batchids, err := internal.GetConfigBatchByUUID(dc.ConfigInfoUUID)
	if err != nil {
		return nil, err
	}
	departids, err := internal.GetConfigDepartByUUID(dc.ConfigInfoUUID)
	if err != nil {
		return nil, err
	}
	nodes, err := internal.GetConfigNodesByUUID(dc.ConfigInfoUUID)
	if err != nil {
		return nil, err
	}

	// 从hc中解析下发的文件内容，逐一进行下发
	dnsfile := common.File{}
	err = json.Unmarshal([]byte(df.Content), &dnsfile)
	if err != nil {
		return nil, err
	}
	results := []NodeResult{}
	de := Deploy{
		DeployBatch: common.Batch{
			BatchIds:      batchids,
			DepartmentIDs: departids,
			MachineUUIDs:  nodes,
		},
		DeployPath:     dnsfile.Path,
		DeployFileName: dnsfile.Name,
		DeployText:     dnsfile.Content,
	}
	url := "http://" + client.GetClient().Server() + "/api/v1/pluginapi/file_deploy"
	r, err := httputils.Post(url, &httputils.Params{
		Body: de,
	})
	if err != nil {
		return nil, err
	}
	if r.StatusCode != http.StatusOK {
		return nil, errors.New("server process error:" + strconv.Itoa(r.StatusCode))
	}

	resp := &common.CommonResult{}
	if err := json.Unmarshal(r.Body, resp); err != nil {
		return nil, err
	}
	if resp.Code != http.StatusOK {
		return nil, errors.New(resp.Message)
	}

	data := []common.NodeResult{}
	if err := resp.ParseData(&data); err != nil {
		return nil, err
	}
	// 将执行失败的文件、机器信息和原因添加到结果字符串中
	for _, d := range data {
		// 存储每一台机器的执行结果
		dfNode := DNSFile{
			UUID:           df.UUID,
			ConfigInfoUUID: df.ConfigInfoUUID,
			Path:           df.Path,
			Name:           df.Name,
			Content:        df.Content,
			Version:        df.Version,
			IsActive:       true,
			IsFromHost:     false,
			Hostuuid:       d.UUID,
			CreatedAt:      time.Now(),
		}

		// 返回执行失败的机器详情
		if d.Error != "" {
			dfNode.IsActive = false
			results = append(results, NodeResult{
				Type:     global.DNS,
				NodeUUID: d.UUID,
				Detail:   dnsfile.Content,
				Result:   false,
				Err:      d.Error,
			})
		}
		err = dfNode.Add()
		if err != nil {
			results = append(results, NodeResult{
				Type:     global.DNS,
				NodeUUID: d.UUID,
				Detail:   "failed to collect dns config to db",
				Result:   false,
				Err:      err.Error(),
			})
		}
	}

	// 全部下发成功直接修改数据库是否激活字段
	if results == nil {
		//下发成功修改数据库应用版本
		err = df.UpdateByuuid()
		return nil, err
	}
	return results, errors.New("failed to apply dns config")
}

// TODO:
func (hc *DNSConfig) Collect() ([]NodeResult, error) {
	results := []NodeResult{}
	return results, errors.New("failed to apply dns config")
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
