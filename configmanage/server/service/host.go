/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: wubijie <wubijie@kylinos.cn>
 * Date: Mon Nov 4 11:31:01 2024 +0800
 */
package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"gitee.com/openeuler/PilotGo/sdk/common"
	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
	"gitee.com/openeuler/PilotGo/sdk/utils/httputils"
	"github.com/google/uuid"
	"openeuler.org/PilotGo/configmanage-plugin/global"
	"openeuler.org/PilotGo/configmanage-plugin/internal"
)

/*
host: 配置文件，hosts文件只有一个

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
		IsFromHost:     false,
	}
}

func toHostConfig(hf *HostFile) HostConfig {
	return HostConfig{
		UUID:           hf.UUID,
		ConfigInfoUUID: hf.ConfigInfoUUID,
		Path:           hf.Path,
		Name:           hf.Name,
		Content:        hf.Content,
		Version:        hf.Version,
		IsActive:       hf.IsActive,
	}
}

func (hc *HostConfig) Record() error {
	hf := hc.toHostFile()
	return hf.Add()
}

func (hc *HostConfig) Load() error {
	// 加载正在使用的某配置文件
	hf, err := internal.GetHostFileByInfoUUID(hc.ConfigInfoUUID, true)
	if err != nil {
		return err
	}
	hc.UUID = hf.UUID
	hc.Path = hf.Path
	hc.Name = hf.Name
	hc.Content = hf.Content
	hc.Version = hf.Version
	hc.IsActive = hf.IsActive
	return nil
}

func (hc *HostConfig) Apply() ([]NodeResult, error) {
	//从数据库获取下发的信息
	hf, err := internal.GetHostFileByUUID(hc.UUID)
	if err != nil {
		return nil, err
	}
	if hf.ConfigInfoUUID != hc.ConfigInfoUUID || hf.UUID != hc.UUID {
		return nil, errors.New("数据库不存在此配置")
	}

	batchids, err := internal.GetConfigBatchByUUID(hc.ConfigInfoUUID)
	if err != nil {
		return nil, err
	}
	departids, err := internal.GetConfigDepartByUUID(hc.ConfigInfoUUID)
	if err != nil {
		return nil, err
	}
	nodes, err := internal.GetConfigNodesByUUID(hc.ConfigInfoUUID)
	if err != nil {
		return nil, err
	}

	//从hc中解析下发的文件内容，逐一进行下发
	hostfile := common.File{}
	err = json.Unmarshal([]byte(hf.Content), &hostfile)
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
		DeployPath:     hostfile.Path,
		DeployFileName: hostfile.Name,
		DeployText:     hostfile.Content,
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
		hfNode := HostFile{
			UUID:           hf.UUID,
			ConfigInfoUUID: hf.ConfigInfoUUID,
			Path:           hf.Path,
			Name:           hf.Name,
			Content:        hf.Content,
			Version:        hf.Version,
			IsActive:       true,
			IsFromHost:     false,
			Hostuuid:       d.UUID,
			CreatedAt:      time.Now(),
		}

		// 返回执行失败的机器详情
		if d.Error != "" {
			hfNode.IsActive = false
			results = append(results, NodeResult{
				Type:     global.Host,
				NodeUUID: d.UUID,
				Detail:   hostfile.Content,
				Result:   false,
				Err:      d.Error,
			})
		}
		err = hfNode.Add()
		if err != nil {
			results = append(results, NodeResult{
				Type:     global.Host,
				NodeUUID: d.UUID,
				Detail:   "failed to collect host config to db",
				Result:   false,
				Err:      err.Error(),
			})
		}
	}

	// 全部下发成功直接修改数据库是否激活字段
	if results == nil {
		//下发成功修改数据库应用版本
		err = hf.UpdateByuuid()
		return nil, err
	}
	return results, errors.New("failed to apply host config")
}

func (hc *HostConfig) Collect() ([]NodeResult, error) {
	ci, err := GetConfigByUUID(hc.ConfigInfoUUID)
	if err != nil {
		return nil, err
	}

	//发请求获取配置详情
	url := "http://" + client.GetClient().Server() + "/api/v1/pluginapi/getnodefiles"
	p := struct {
		DeployBatch common.Batch `json:"deploybatch"`
		Path        string       `json:"path"`
		FileName    string       `json:"filename"`
	}{
		DeployBatch: common.Batch{
			BatchIds:      ci.BatchIds,
			DepartmentIDs: ci.DepartIds,
			MachineUUIDs:  ci.Nodes,
		},
		Path:     hc.Path,
		FileName: hc.Name,
	}
	r, err := httputils.Post(url, &httputils.Params{
		Body: p,
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
	results := []NodeResult{}
	for _, v := range data {
		if v.Error == "" {
			file, _ := json.Marshal(v.Data)
			hf := HostFile{
				UUID:           uuid.New().String(),
				ConfigInfoUUID: hc.ConfigInfoUUID,
				Path:           hc.Path,
				Name:           hc.Name,
				Content:        file,
				Version:        fmt.Sprintf("v%s", time.Now().Format("2006-01-02-15-04-05")),
				IsActive:       true,
				IsFromHost:     true,
				Hostuuid:       v.UUID,
				CreatedAt:      time.Now(),
			}
			err = hf.Add()
			if err != nil {
				logger.Error("failed to add hostconfig: %s", err.Error())
				results = append(results, NodeResult{
					Type:     global.Host,
					NodeUUID: v.UUID,
					Detail:   "failed to collect host config to db",
					Result:   false,
					Err:      err.Error()})
			}
		} else {
			results = append(results, NodeResult{
				Type:     global.Host,
				NodeUUID: v.UUID,
				Detail:   "failed to collect host config:" + v.Data.(string),
				Result:   false,
				Err:      v.Error})
		}
	}
	if results != nil {
		return results, errors.New("failed to collect host config")
	}
	return nil, nil
}

func GetHostFileByInfoUUID(uuid string, isindex interface{}) (HostFile, error) {
	return internal.GetHostFileByInfoUUID(uuid, isindex)
}

// 根据配置uuid获取所有配置文件
func GetHostFilesByConfigUUID(uuid string) ([]HostFile, error) {
	return internal.GetHostFilesByConfigUUID(uuid)
}

// 查看某台机器某种类型的的历史配置信息
func GetHostFilesByNode(nodeid string) ([]HostConfig, error) {
	// 查找本台机器所属的配置uuid
	config_nodes, err := internal.GetConfigNodesByNode(nodeid)
	if err != nil {
		return nil, err
	}
	var hcs []HostConfig
	for _, v := range config_nodes {
		hf, err := internal.GetHostFileByInfoUUID(v.ConfigInfoUUID, nil)
		if err != nil {
			return nil, err
		}
		hc := toHostConfig(&hf)
		hcs = append(hcs, hc)
	}
	return hcs, nil
}
