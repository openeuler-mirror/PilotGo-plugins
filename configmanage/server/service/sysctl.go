/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: wubijie <wubijie@kylinos.cn>
 * Date: Thu Nov 7 10:19:32 2024 +0800
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
修改内核参数的方法：1、sysctl 命令；2、挂载于 /proc/sys/ 目录的虚拟文件系统；3、/etc/sysctl.d/ 目录中的配置文件
采用1、2方法可以修改临时配置，重启后恢复原来配置；采用1、3方法可以永久修改配置

sysctl: 配置文件，sysctl.conf文件只有一个

一般方法：	1、在/etc/ysctl.conf中修改内容，/etc/sysctl.conf是系统控制文件，主要用于运行时配置内核参数等系统信息，它的内容全部对应于/proc/sys/目录及其子目录。

	2、有以下两种方法使其生效：

（1）重启机器reboot。但是一般不推荐，因为可能机器上正在运行的程序

（2）使用如下命令刷新配置，使其立即生效。

/sbin/sysctl -p

/sbin/sysctl -w net.ipv4.route.flush=1

考虑的问题：1、目前只考虑创建、下发/etc/sysctl.conf，不考虑其他子文件的，后续可继续完善。
*/
type SysctlFile = internal.SysctlFile
type SysctlConfig struct {
	UUID           string          `json:"uuid"`
	ConfigInfoUUID string          `json:"configinfouuid"`
	Content        json.RawMessage `json:"content"`
	Version        string          `json:"version"`
	Path           string          `json:"path"`
	Name           string          `json:"name"`
	//下发改变标志位
	IsActive bool `json:"isactive"`
}

func (sysc *SysctlConfig) toSysctlFile() SysctlFile {
	return SysctlFile{
		UUID:           sysc.UUID,
		ConfigInfoUUID: sysc.ConfigInfoUUID,
		Path:           sysc.Path,
		Name:           sysc.Name,
		Content:        sysc.Content,
		Version:        fmt.Sprintf("v%s", time.Now().Format("2006-01-02-15-04-05")),
		IsActive:       sysc.IsActive,
		IsFromHost:     false,
		CreatedAt:      time.Now(),
	}
}

func toSysctlConfig(sysf *SysctlFile) SysctlConfig {
	return SysctlConfig{
		UUID:           sysf.UUID,
		ConfigInfoUUID: sysf.ConfigInfoUUID,
		Path:           sysf.Path,
		Name:           sysf.Name,
		Content:        sysf.Content,
		Version:        sysf.Version,
		IsActive:       sysf.IsActive,
	}
}

func (sysc *SysctlConfig) Record() error {
	sysf := sysc.toSysctlFile()
	return sysf.Add()
}

func (sysc *SysctlConfig) Load() error {
	// 加载正在使用的某配置文件
	sysf, err := internal.GetSysctlFileByInfoUUID(sysc.ConfigInfoUUID, true)
	if err != nil {
		return err
	}
	sysc.UUID = sysf.UUID
	sysc.Path = sysf.Path
	sysc.Name = sysf.Name
	sysc.Content = sysf.Content
	sysc.Version = sysf.Version
	sysc.IsActive = sysf.IsActive
	return nil
}

func (sysc *SysctlConfig) Apply() ([]NodeResult, error) {
	// 从数据库获取下发的信息
	sysf, err := internal.GetSysctlFileByUUID(sysc.UUID)
	if err != nil {
		return nil, err
	}
	if sysf.ConfigInfoUUID != sysc.ConfigInfoUUID || sysf.UUID != sysc.UUID {
		return nil, errors.New("数据库不存在此配置")
	}

	batchids, err := internal.GetConfigBatchByUUID(sysc.ConfigInfoUUID)
	if err != nil {
		return nil, err
	}
	departids, err := internal.GetConfigDepartByUUID(sysc.ConfigInfoUUID)
	if err != nil {
		return nil, err
	}
	nodes, err := internal.GetConfigNodesByUUID(sysc.ConfigInfoUUID)
	if err != nil {
		return nil, err
	}

	// 从hc中解析下发的文件内容，逐一进行下发
	sysctlfile := common.File{}
	err = json.Unmarshal([]byte(sysf.Content), &sysctlfile)
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
		DeployPath:     sysctlfile.Path,
		DeployFileName: sysctlfile.Name,
		DeployText:     sysctlfile.Content,
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
		sysfNode := SysctlFile{
			UUID:           sysf.UUID,
			ConfigInfoUUID: sysf.ConfigInfoUUID,
			Path:           sysf.Path,
			Name:           sysf.Name,
			Content:        sysf.Content,
			Version:        sysf.Version,
			IsActive:       true,
			IsFromHost:     false,
			Hostuuid:       d.UUID,
			CreatedAt:      time.Now(),
		}

		// 返回执行失败的机器详情
		if d.Error != "" {
			sysfNode.IsActive = false
			results = append(results, NodeResult{
				Type:     global.Sysctl,
				NodeUUID: d.UUID,
				Detail:   sysctlfile.Content,
				Result:   false,
				Err:      d.Error,
			})
		}
		err = sysfNode.Add()
		if err != nil {
			results = append(results, NodeResult{
				Type:     global.Sysctl,
				NodeUUID: d.UUID,
				Detail:   "failed to collect sysctl config to db",
				Result:   false,
				Err:      err.Error(),
			})
		}
	}

	// 全部下发成功直接修改数据库是否激活字段
	if results == nil {
		//下发成功修改数据库应用版本
		err = sysf.UpdateByuuid()
		return nil, err
	}
	return results, errors.New("failed to apply SysctlConfig")
}

func (sysc *SysctlConfig) Collect() ([]NodeResult, error) {
	ci, err := GetConfigByUUID(sysc.ConfigInfoUUID)
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
		Path:     sysc.Path,
		FileName: sysc.Name,
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
			sysf := SysctlFile{
				UUID:           uuid.New().String(),
				ConfigInfoUUID: sysc.ConfigInfoUUID,
				Path:           sysc.Path,
				Name:           sysc.Name,
				Content:        file,
				Version:        fmt.Sprintf("v%s", time.Now().Format("2006-01-02-15-04-05")),
				IsFromHost:     true,
				IsActive:       true,
				Hostuuid:       v.UUID,
				CreatedAt:      time.Now(),
			}
			err = sysf.Add()
			if err != nil {
				logger.Error("failed to add sysctl config: %s", err.Error())
				results = append(results, NodeResult{
					Type:     global.Sysctl,
					NodeUUID: v.UUID,
					Detail:   "failed to collect sysctl config to db",
					Result:   false,
					Err:      err.Error()})
			}
		} else {
			results = append(results, NodeResult{
				Type:     global.Sysctl,
				NodeUUID: v.UUID,
				Detail:   "failed to collect sysctl config:" + v.Data.(string),
				Result:   false,
				Err:      v.Error})
		}
	}
	if results != nil {
		return results, errors.New("failed to collect sysctl config")
	}
	return nil, nil
}

func GetSysctlFileByInfoUUID(uuid string, isindex interface{}) (SysctlFile, error) {
	return internal.GetSysctlFileByInfoUUID(uuid, isindex)
}

// 根据配置uuid获取所有配置文件
func GetSysctlFilesByConfigUUID(uuid string) ([]SysctlFile, error) {
	return internal.GetSysctlFilesByConfigUUID(uuid)
}

// 查看某台机器某种类型的的历史配置信息
func GetSysctlFilesByNode(nodeid string) ([]SysctlConfig, error) {
	// 查找本台机器所属的配置uuid
	config_nodes, err := internal.GetConfigNodesByNode(nodeid)
	if err != nil {
		return nil, err
	}
	var syscs []SysctlConfig
	for _, v := range config_nodes {
		sysf, err := internal.GetSysctlFileByInfoUUID(v.ConfigInfoUUID, nil)
		if err != nil {
			return nil, err
		}
		sysc := toSysctlConfig(&sysf)
		syscs = append(syscs, sysc)
	}
	return syscs, nil
}
