/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin licensed under the Mulan Permissive Software License, Version 2.
 * See LICENSE file for more details.
 * Author: wubijie <wubijie@kylinos.cn>
 * Date: Tue Nov 5 11:20:32 2024 +0800
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
	"gitee.com/openeuler/PilotGo/sdk/utils/httputils"
	"github.com/google/uuid"
	"openeuler.org/PilotGo/configmanage-plugin/global"
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
		IsFromHost:     false,
		CreatedAt:      time.Now(),
	}
}

func toSSHConfig(sf *SSHFile) SSHConfig {
	return SSHConfig{
		UUID:           sf.UUID,
		ConfigInfoUUID: sf.ConfigInfoUUID,
		Path:           sf.Path,
		Name:           sf.Name,
		Content:        sf.Content,
		Version:        sf.Version,
		IsActive:       sf.IsActive,
	}
}

func (sc *SSHConfig) Record() error {
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

func (sc *SSHConfig) Apply() ([]NodeResult, error) {
	// 从数据库获取下发的信息
	sf, err := internal.GetSSHFileByUUID(sc.UUID)
	if err != nil {
		return nil, err
	}
	if sf.ConfigInfoUUID != sc.ConfigInfoUUID || sf.UUID != sc.UUID {
		return nil, errors.New("数据库不存在此配置")
	}

	batchids, err := internal.GetConfigBatchByUUID(sc.ConfigInfoUUID)
	if err != nil {
		return nil, err
	}
	departids, err := internal.GetConfigDepartByUUID(sc.ConfigInfoUUID)
	if err != nil {
		return nil, err
	}
	nodes, err := internal.GetConfigNodesByUUID(sc.ConfigInfoUUID)
	if err != nil {
		return nil, err
	}

	// 从hc中解析下发的文件内容，逐一进行下发
	sshfile := common.File{}
	err = json.Unmarshal([]byte(sf.Content), &sshfile)
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
		DeployPath:     sshfile.Path,
		DeployFileName: sshfile.Name,
		DeployText:     sshfile.Content,
	}
	serverInfo, err := global.GlobalClient.Registry.Get("pilotgo-server")
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("http://%s:%s/api/v1/pluginapi/file_deploy", serverInfo.Address, serverInfo.Port)
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
		sfNode := SSHFile{
			UUID:           sf.UUID,
			ConfigInfoUUID: sf.ConfigInfoUUID,
			Path:           sf.Path,
			Name:           sf.Name,
			Content:        sf.Content,
			Version:        sf.Version,
			IsActive:       true,
			IsFromHost:     false,
			Hostuuid:       d.UUID,
			CreatedAt:      time.Now(),
		}

		// 返回执行失败的机器详情
		if d.Error != "" {
			sfNode.IsActive = false
			results = append(results, NodeResult{
				Type:     global.SSH,
				NodeUUID: d.UUID,
				Detail:   sshfile.Content,
				Result:   false,
				Err:      d.Error,
			})
		}
		err = sfNode.Add()
		if err != nil {
			results = append(results, NodeResult{
				Type:     global.SSH,
				NodeUUID: d.UUID,
				Detail:   "failed to collect ssh config to db",
				Result:   false,
				Err:      err.Error(),
			})
		}
	}

	// 全部下发成功直接修改数据库是否激活字段
	if results == nil {
		//下发成功修改数据库应用版本
		err = sf.UpdateByuuid()
		return nil, err
	}
	return results, errors.New("failed to apply SSHConfig")
}

func (sc *SSHConfig) Collect() ([]NodeResult, error) {
	ci, err := GetConfigByUUID(sc.ConfigInfoUUID)
	if err != nil {
		return nil, err
	}

	//发请求获取配置详情
	serverInfo, err := global.GlobalClient.Registry.Get("pilotgo-server")
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("http://%s:%s/api/v1/pluginapi/getnodefiles", serverInfo.Address, serverInfo.Port)
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
		Path:     sc.Path,
		FileName: sc.Name,
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
			sf := SSHFile{
				UUID:           uuid.New().String(),
				ConfigInfoUUID: sc.ConfigInfoUUID,
				Path:           sc.Path,
				Name:           sc.Name,
				Content:        file,
				Version:        fmt.Sprintf("v%s", time.Now().Format("2006-01-02-15-04-05")),
				IsFromHost:     true,
				IsActive:       true,
				Hostuuid:       v.UUID,
				CreatedAt:      time.Now(),
			}
			err = sf.Add()
			if err != nil {
				logger.Error("failed to add sshconfig: %s", err.Error())
				results = append(results, NodeResult{
					Type:     global.SSH,
					NodeUUID: v.UUID,
					Detail:   "failed to collect SSH config to db",
					Result:   false,
					Err:      err.Error()})
			}
		} else {
			results = append(results, NodeResult{
				Type:     global.SSH,
				NodeUUID: v.UUID,
				Detail:   "failed to collect SSH config:" + v.Data.(string),
				Result:   false,
				Err:      v.Error})
		}
	}
	if results != nil {
		return results, errors.New("failed to collect ssh config")
	}
	return nil, nil
}

func GetSSHFileByInfoUUID(uuid string, isindex interface{}) (SSHFile, error) {
	return internal.GetSSHFileByInfoUUID(uuid, isindex)
}

// 根据配置uuid获取所有配置文件
func GetSSHFilesByConfigUUID(uuid string) ([]SSHFile, error) {
	return internal.GetSSHFilesByConfigUUID(uuid)
}

// 查看某台机器某种类型的的历史配置信息
func GetSSHFilesByNode(nodeid string) ([]SSHConfig, error) {
	// 查找本台机器所属的配置uuid
	config_nodes, err := internal.GetConfigNodesByNode(nodeid)
	if err != nil {
		return nil, err
	}
	var scs []SSHConfig
	for _, v := range config_nodes {
		sf, err := internal.GetSSHFileByInfoUUID(v.ConfigInfoUUID, nil)
		if err != nil {
			return nil, err
		}
		sc := toSSHConfig(&sf)
		scs = append(scs, sc)
	}
	return scs, nil
}
