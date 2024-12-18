/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: wubijie <wubijie@kylinos.cn>
 * Date: Fri Jan 5 15:17:58 2024 +0800
 */
package controller

import (
	"encoding/base64"
	"encoding/json"

	"gitee.com/openeuler/PilotGo/sdk/common"
	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"openeuler.org/PilotGo/configmanage-plugin/global"
	"openeuler.org/PilotGo/configmanage-plugin/service"
)

// AddConfigHandler 添加配置管理
func AddConfigHandler(c *gin.Context) {
	query := &struct {
		Type        string          `json:"type"`
		Description string          `json:"description"`
		BatchIds    []int           `json:"batchids"`
		DepartIds   []int           `json:"departids"`
		Nodes       []string        `json:"uuids"`
		Data        json.RawMessage `json:"data"`
	}{}
	err := c.ShouldBindJSON(query)
	if err != nil {
		response.Fail(c, "parameter error", err.Error())
		return
	}
	logger.Debug("add config")
	//存储配置管理类型
	ci := service.ConfigInstance{
		UUID:        uuid.New().String(),
		Type:        query.Type,
		Description: query.Description,
		BatchIds:    query.BatchIds,
		DepartIds:   query.DepartIds,
		Nodes:       query.Nodes,
	}
	err = ci.Add()
	if err != nil {
		logger.Error("failed to add configinstance: %s", err.Error())
		response.Fail(c, "failed to add configinstance:", err.Error())
		return
	}

	//解析对应配置管理的参数
	switch query.Type {
	case global.Repo:
		//解析参数
		repoconfig := &service.RepoConfig{}
		if err := json.Unmarshal(query.Data, repoconfig); err != nil {
			logger.Error("failed to parse repoconfig parameter: %s", err.Error())
			response.Fail(c, "failed to parse repoconfig parameter:", err.Error())
			return
		}

		files := []common.File{}
		if err := json.Unmarshal([]byte(repoconfig.Content), &files); err != nil {
			logger.Error("failed to parse file parameter: %s", err.Error())
			response.Fail(c, "failed to parse file parameter:", err.Error())
			return
		}
		for i, v := range files {
			files[i].Content = base64.StdEncoding.EncodeToString([]byte(v.Content))
		}

		repoconfig.UUID = uuid.New().String()
		repoconfig.ConfigInfoUUID = ci.UUID
		repoconfig.IsActive = false
		repoconfig.Content, err = json.Marshal(files)
		if err != nil {
			logger.Error("Error encoding JSON:: %s", err.Error())
			response.Fail(c, "Error encoding JSON:", err.Error())
			return
		}

		//将参数添加到数据库
		err = repoconfig.Record()
		if err != nil {
			logger.Error("failed to add repoconfig: %s", err.Error())
			response.Fail(c, "failed to add repoconfig:", err.Error())
			return
		}
		//收集机器的配置信息
		result, err := repoconfig.Collect()
		if err != nil {
			logger.Error("failed to collect repofile: %s", err.Error())
			response.Fail(c, result, err.Error())
			return
		}
		logger.Debug("add repoconfig success")
		response.Success(c, nil, "Add repo config success")

	case global.Host:
		// 解析参数
		hostconfig := &service.HostConfig{}
		if err := json.Unmarshal(query.Data, hostconfig); err != nil {
			logger.Error("failed to parse hostconfig parameter: %s", err.Error())
			response.Fail(c, "failed to parse hostconfig parameter:", err.Error())
			return
		}

		file := common.File{}
		if err := json.Unmarshal([]byte(hostconfig.Content), &file); err != nil {
			logger.Error("failed to parse file parameter: %s", err.Error())
			response.Fail(c, "failed to parse file parameter:", err.Error())
			return
		}
		file.Content = base64.StdEncoding.EncodeToString([]byte(file.Content))

		hostconfig.UUID = uuid.New().String()
		hostconfig.ConfigInfoUUID = ci.UUID
		hostconfig.IsActive = false
		hostconfig.Content, err = json.Marshal(file)
		if err != nil {
			logger.Error("Error encoding JSON:: %s", err.Error())
			response.Fail(c, "Error encoding JSON:", err.Error())
			return
		}

		// 将参数添加到数据库
		err = hostconfig.Record()
		if err != nil {
			logger.Error("failed to add hostconfig: %s", err.Error())
			response.Fail(c, "failed to add hostconfig:", err.Error())
			return
		}
		// 收集机器的配置信息
		result, err := hostconfig.Collect()
		if err != nil {
			logger.Error("failed to collect hostconfig: %s", err.Error())
			response.Fail(c, result, err.Error())
			return
		}
		logger.Debug("add hostconfig success")
		response.Success(c, nil, "Add host config success")

	case global.SSH:
		// 解析参数
		sshconfig := &service.SSHConfig{}
		if err := json.Unmarshal(query.Data, sshconfig); err != nil {
			logger.Error("failed to parse sshconfig parameter: %s", err.Error())
			response.Fail(c, "failed to parse sshconfig parameter:", err.Error())
			return
		}

		file := common.File{}
		if err := json.Unmarshal([]byte(sshconfig.Content), &file); err != nil {
			logger.Error("failed to parse file parameter: %s", err.Error())
			response.Fail(c, "failed to parse file parameter:", err.Error())
			return
		}
		file.Content = base64.StdEncoding.EncodeToString([]byte(file.Content))

		sshconfig.UUID = uuid.New().String()
		sshconfig.ConfigInfoUUID = ci.UUID
		sshconfig.IsActive = false
		sshconfig.Content, err = json.Marshal(file)
		if err != nil {
			logger.Error("Error encoding JSON:: %s", err.Error())
			response.Fail(c, "Error encoding JSON:", err.Error())
			return
		}

		// 将参数添加到数据库
		err = sshconfig.Record()
		if err != nil {
			logger.Error("failed to add sshconfig: %s", err.Error())
			response.Fail(c, "failed to add sshconfig:", err.Error())
			return
		}
		// 收集机器的配置信息
		result, err := sshconfig.Collect()
		if err != nil {
			logger.Error("failed to collect sshconfig: %s", err.Error())
			response.Fail(c, result, err.Error())
			return
		}
		logger.Debug("add sshconfig success")
		response.Success(c, nil, "Add sshconfig success")

	case global.SSHD:
		// 解析参数
		sshdconfig := &service.SSHDConfig{}
		if err := json.Unmarshal(query.Data, sshdconfig); err != nil {
			logger.Error("failed to parse sshdconfig parameter: %s", err.Error())
			response.Fail(c, "failed to parse sshdconfig parameter:", err.Error())
			return
		}

		file := common.File{}
		if err := json.Unmarshal([]byte(sshdconfig.Content), &file); err != nil {
			logger.Error("failed to parse file parameter: %s", err.Error())
			response.Fail(c, "failed to parse file parameter:", err.Error())
			return
		}
		file.Content = base64.StdEncoding.EncodeToString([]byte(file.Content))

		sshdconfig.UUID = uuid.New().String()
		sshdconfig.ConfigInfoUUID = ci.UUID
		sshdconfig.IsActive = false
		sshdconfig.Content, err = json.Marshal(file)
		if err != nil {
			logger.Error("Error encoding JSON:: %s", err.Error())
			response.Fail(c, "Error encoding JSON:", err.Error())
			return
		}

		// 将参数添加到数据库
		err = sshdconfig.Record()
		if err != nil {
			logger.Error("failed to add sshdconfig: %s", err.Error())
			response.Fail(c, "failed to add sshdconfig:", err.Error())
			return
		}
		// 收集机器的配置信息
		result, err := sshdconfig.Collect()
		if err != nil {
			logger.Error("failed to collect sshdconfig: %s", err.Error())
			response.Fail(c, result, err.Error())
			return
		}
		logger.Debug("add sshdconfig success")
		response.Success(c, nil, "Add sshdconfig success")

	case global.Sysctl:
		// 解析参数
		sysctlconfig := &service.SysctlConfig{}
		if err := json.Unmarshal(query.Data, sysctlconfig); err != nil {
			logger.Error("failed to parse sysctlconfig parameter: %s", err.Error())
			response.Fail(c, "failed to parse sysctlconfig parameter:", err.Error())
			return
		}

		file := common.File{}
		if err := json.Unmarshal([]byte(sysctlconfig.Content), &file); err != nil {
			logger.Error("failed to parse file parameter: %s", err.Error())
			response.Fail(c, "failed to parse file parameter:", err.Error())
			return
		}
		file.Content = base64.StdEncoding.EncodeToString([]byte(file.Content))

		sysctlconfig.UUID = uuid.New().String()
		sysctlconfig.ConfigInfoUUID = ci.UUID
		sysctlconfig.IsActive = false
		sysctlconfig.Content, err = json.Marshal(file)
		if err != nil {
			logger.Error("Error encoding JSON:: %s", err.Error())
			response.Fail(c, "Error encoding JSON:", err.Error())
			return
		}

		// 将参数添加到数据库
		err = sysctlconfig.Record()
		if err != nil {
			logger.Error("failed to add sysctlconfig: %s", err.Error())
			response.Fail(c, "failed to add sysctlconfig:", err.Error())
			return
		}
		// 收集机器的配置信息
		result, err := sysctlconfig.Collect()
		if err != nil {
			logger.Error("failed to collect sysctlconfig: %s", err.Error())
			response.Fail(c, result, err.Error())
			return
		}
		logger.Debug("add sysctlconfig success")
		response.Success(c, nil, "Add sysctlconfig success")

	case global.DNS:
		// 解析参数
		dnsconfig := &service.DNSConfig{}
		if err := json.Unmarshal(query.Data, dnsconfig); err != nil {
			logger.Error("failed to parse dnsconfig parameter: %s", err.Error())
			response.Fail(c, "failed to parse dnsconfig parameter:", err.Error())
			return
		}

		file := common.File{}
		if err := json.Unmarshal([]byte(dnsconfig.Content), &file); err != nil {
			logger.Error("failed to parse file parameter: %s", err.Error())
			response.Fail(c, "failed to parse file parameter:", err.Error())
			return
		}
		file.Content = base64.StdEncoding.EncodeToString([]byte(file.Content))

		dnsconfig.UUID = uuid.New().String()
		dnsconfig.ConfigInfoUUID = ci.UUID
		dnsconfig.IsActive = false
		dnsconfig.Content, err = json.Marshal(file)
		if err != nil {
			logger.Error("Error encoding JSON:: %s", err.Error())
			response.Fail(c, "Error encoding JSON:", err.Error())
			return
		}

		// 将参数添加到数据库
		err = dnsconfig.Record()
		if err != nil {
			logger.Error("failed to add dnsconfig: %s", err.Error())
			response.Fail(c, "failed to add dnsconfig:", err.Error())
			return
		}
		// 收集机器的配置信息
		result, err := dnsconfig.Collect()
		if err != nil {
			logger.Error("failed to collect dnsconfig: %s", err.Error())
			response.Fail(c, result, err.Error())
			return
		}
		logger.Debug("add dnsconfig success")
		response.Success(c, nil, "Add dnsconfig success")

	case global.PATH:
		// 解析参数
		pathconfig := &service.PathConfig{}
		if err := json.Unmarshal(query.Data, pathconfig); err != nil {
			logger.Error("failed to parse pathconfig parameter: %s", err.Error())
			response.Fail(c, "failed to parse pathconfig parameter:", err.Error())
			return
		}

		file := common.File{}
		if err := json.Unmarshal([]byte(pathconfig.Content), &file); err != nil {
			logger.Error("failed to parse file parameter: %s", err.Error())
			response.Fail(c, "failed to parse file parameter:", err.Error())
			return
		}
		file.Content = base64.StdEncoding.EncodeToString([]byte(file.Content))

		pathconfig.UUID = uuid.New().String()
		pathconfig.ConfigInfoUUID = ci.UUID
		pathconfig.IsActive = false
		pathconfig.Content, err = json.Marshal(file)
		if err != nil {
			logger.Error("Error encoding JSON:: %s", err.Error())
			response.Fail(c, "Error encoding JSON:", err.Error())
			return
		}

		// 将参数添加到数据库
		err = pathconfig.Record()
		if err != nil {
			logger.Error("failed to add pathconfig: %s", err.Error())
			response.Fail(c, "failed to add pathconfig:", err.Error())
			return
		}
		// 收集机器的配置信息
		result, err := pathconfig.Collect()
		if err != nil {
			logger.Error("failed to collect pathconfig: %s", err.Error())
			response.Fail(c, result, err.Error())
			return
		}
		logger.Debug("add pathconfig success")
		response.Success(c, nil, "Add pathconfig success")

	default:
		response.Fail(c, nil, "Unknown type:"+query.Type)
	}
}

// 加载数据库中存储的正在使用的配置文件信息
func LoadConfigHandler(c *gin.Context) {
	//TODO:修改请求的参数
	query := &struct {
		UUID string `json:"uuid"`
	}{}
	err := c.ShouldBindJSON(query)
	if err != nil {
		response.Fail(c, "parameter error", err.Error())
		return
	}
	logger.Debug("load config")

	//获取ConfigInstance
	ci, err := service.GetConfigByUUID(query.UUID)
	if err != nil {
		logger.Error("failed to get configinfo file: %s", err.Error())
		response.Fail(c, "get configinfo fail:", err.Error())
		return
	}

	//获取对应配置管理的参数
	switch ci.Type {
	case global.Repo:
		repoconfig := &service.RepoConfig{
			ConfigInfoUUID: ci.UUID,
		}
		// 加载正在使用的配置
		err = repoconfig.Load()
		if err != nil {
			logger.Error("failed to get repofile file: %s", err.Error())
			response.Fail(c, "failed to get repofile file:", err.Error())
			return
		}
		ci.Config = repoconfig
		logger.Debug("load repoconfig success")
		response.Success(c, ci, "load repo config success")

	case global.Host:
		hostconfig := &service.HostConfig{
			ConfigInfoUUID: ci.UUID,
		}
		// 加载正在使用的配置
		err = hostconfig.Load()
		if err != nil {
			logger.Error("failed to get hostconfig file: %s", err.Error())
			response.Fail(c, "failed to get hostconfig file:", err.Error())
			return
		}
		ci.Config = hostconfig
		logger.Debug("load hostconfig success")
		response.Success(c, ci, "load hostconfig success")

	case global.SSH:
		sshconfig := &service.SSHConfig{
			ConfigInfoUUID: ci.UUID,
		}
		// 加载正在使用的配置
		err = sshconfig.Load()
		if err != nil {
			logger.Error("failed to get sshconfig file: %s", err.Error())
			response.Fail(c, "failed to get sshconfig file:", err.Error())
			return
		}
		ci.Config = sshconfig
		logger.Debug("load sshconfig success")
		response.Success(c, ci, "load sshconfig success")

	case global.SSHD:
		sshdconfig := &service.SSHDConfig{
			ConfigInfoUUID: ci.UUID,
		}
		// 加载正在使用的配置
		err = sshdconfig.Load()
		if err != nil {
			logger.Error("failed to get sshdconfig file: %s", err.Error())
			response.Fail(c, "failed to get sshdconfig file:", err.Error())
			return
		}
		ci.Config = sshdconfig
		logger.Debug("load sshdconfig success")
		response.Success(c, ci, "load sshdconfig success")

	case global.Sysctl:
		sysctlconfig := &service.SysctlConfig{
			ConfigInfoUUID: ci.UUID,
		}
		// 加载正在使用的配置
		err = sysctlconfig.Load()
		if err != nil {
			logger.Error("failed to get sysctlconfig file: %s", err.Error())
			response.Fail(c, "failed to get sysctlconfig file:", err.Error())
			return
		}
		ci.Config = sysctlconfig
		logger.Debug("load sysctlconfig success")
		response.Success(c, ci, "load sysctlconfig success")

	case global.DNS:
		dnsconfig := &service.DNSConfig{
			ConfigInfoUUID: ci.UUID,
		}
		// 加载正在使用的配置
		err = dnsconfig.Load()
		if err != nil {
			logger.Error("failed to get dnsconfig file: %s", err.Error())
			response.Fail(c, "failed to get dnsconfig file:", err.Error())
			return
		}
		ci.Config = dnsconfig
		logger.Debug("load dnsconfig success")
		response.Success(c, ci, "load dnsconfig success")

	case global.PATH:
		pathconfig := &service.PathConfig{
			ConfigInfoUUID: ci.UUID,
		}
		// 加载正在使用的配置
		err = pathconfig.Load()
		if err != nil {
			logger.Error("failed to get pathconfig file: %s", err.Error())
			response.Fail(c, "failed to get pathconfig file:", err.Error())
			return
		}
		ci.Config = pathconfig
		logger.Debug("load pathconfig success")
		response.Success(c, ci, "load pathconfig success")

	default:
		response.Fail(c, nil, "Unknown type of configinfo:"+query.UUID)
	}
}

// TODO： 考虑文件下发和执行命令使用配置
func ApplyConfigHandler(c *gin.Context) {
	//TODO:修改请求的参数
	query := &struct {
		ConfigInfoUUID string `json:"configinfouuid"`
		UUID           string `json:"uuid"`
	}{}
	err := c.Bind(query)
	if err != nil {
		response.Fail(c, "parameter error", err.Error())
		return
	}
	logger.Debug("start configuration apply")

	//获取ConfigInstance
	ci, err := service.GetInfoByUUID(query.ConfigInfoUUID)
	if err != nil {
		logger.Error("failed to get configinfo file: %s", err.Error())
		response.Fail(c, "failed to get configinfo file:", err.Error())
		return
	}

	//解析对应配置管理的参数
	switch ci.Type {
	case global.Repo:
		repoconfig := &service.RepoConfig{
			UUID:           query.UUID,
			ConfigInfoUUID: ci.UUID,
		}
		_, err := repoconfig.Apply()
		if err != nil {
			logger.Error("failed to apply repoconfig file: %s", err.Error())
			response.Fail(c, "failed to apply repofile:", err.Error())
			return
		}
		response.Success(c, nil, "apply repo config success")

	case global.Host:
		hostconfig := &service.HostConfig{
			UUID:           query.UUID,
			ConfigInfoUUID: ci.UUID,
		}
		_, err := hostconfig.Apply()
		if err != nil {
			logger.Error("failed to apply hostconfig file: %s", err.Error())
			response.Fail(c, "failed to apply hostconfig:", err.Error())
			return
		}
		response.Success(c, nil, "apply hostconfig success")

	case global.SSH:
		sshconfig := &service.SSHConfig{
			UUID:           query.UUID,
			ConfigInfoUUID: ci.UUID,
		}
		_, err := sshconfig.Apply()
		if err != nil {
			logger.Error("failed to apply sshconfig file: %s", err.Error())
			response.Fail(c, "failed to apply sshconfig:", err.Error())
			return
		}
		response.Success(c, nil, "apply sshconfig success")

	case global.SSHD:
		sshdconfig := &service.SSHDConfig{
			UUID:           query.UUID,
			ConfigInfoUUID: ci.UUID,
		}
		_, err := sshdconfig.Apply()
		if err != nil {
			logger.Error("failed to apply sshdconfig file: %s", err.Error())
			response.Fail(c, "failed to apply sshdconfig:", err.Error())
			return
		}
		response.Success(c, nil, "apply sshdconfig success")

	case global.Sysctl:
		sysctlconfig := &service.SysctlConfig{
			UUID:           query.UUID,
			ConfigInfoUUID: ci.UUID,
		}
		_, err := sysctlconfig.Apply()
		if err != nil {
			logger.Error("failed to apply sysctlconfig file: %s", err.Error())
			response.Fail(c, "failed to apply sysctlconfig:", err.Error())
			return
		}
		response.Success(c, nil, "apply sysctlconfig success")

	case global.DNS:
		dnsconfig := &service.DNSConfig{
			UUID:           query.UUID,
			ConfigInfoUUID: ci.UUID,
		}
		_, err := dnsconfig.Apply()
		if err != nil {
			logger.Error("failed to apply dnsconfig file: %s", err.Error())
			response.Fail(c, "failed to apply dnsconfig:", err.Error())
			return
		}
		response.Success(c, nil, "apply dnsconfig success")

	case global.PATH:
		pathconfig := &service.PathConfig{
			UUID:           query.UUID,
			ConfigInfoUUID: ci.UUID,
		}
		_, err := pathconfig.Apply()
		if err != nil {
			logger.Error("failed to apply pathconfig file: %s", err.Error())
			response.Fail(c, "failed to apply pathconfig:", err.Error())
			return
		}
		response.Success(c, nil, "apply pathconfig success")

	default:
		response.Fail(c, nil, "Unknown type of configinfo:"+query.UUID)
	}
}

// 修改某配置信息
func UpdateConfigHandler(c *gin.Context) {
	// 某配置的uuid和类型不可以修改，具体配置中的uuid不可修改
	query := &struct {
		UUID        string          `json:"uuid"`
		Type        string          `json:"type"`
		Description string          `json:"description"`
		BatchIds    []int           `json:"batchids"`
		DepartIds   []int           `json:"departids"`
		Nodes       []string        `json:"uuids"`
		Data        json.RawMessage `json:"data"`
	}{}
	err := c.ShouldBindJSON(query)
	if err != nil {
		response.Fail(c, "parameter error", err.Error())
		return
	}
	logger.Debug("load config")

	// 获取ConfigInstance
	ci, err := service.GetConfigByUUID(query.UUID)
	if err != nil {
		logger.Error("failed to get configinfo file: %s", err.Error())
		response.Fail(c, "get configinfo fail:", err.Error())
		return
	}

	// 判断类型是否一致，或者数据库
	if ci == nil || ci.Type != query.Type {
		response.Fail(c, "configinfo not exist", "")
		return
	}

	// 更新配置信息
	ci.Description = query.Description
	ci.BatchIds = query.BatchIds
	ci.DepartIds = query.DepartIds
	ci.Nodes = query.Nodes
	err = ci.Add()
	if err != nil {
		logger.Error("failed to update configinstance: %s", err.Error())
		response.Fail(c, "failed to update configinstance:", err.Error())
		return
	}

	// 获取对应配置管理的参数
	switch ci.Type {
	case global.Repo:
		// 解析参数
		repoconfig := &service.RepoConfig{}
		if err := json.Unmarshal(query.Data, repoconfig); err != nil {
			logger.Error("failed to parse repoconfig parameter: %s", err.Error())
			response.Fail(c, "failed to parse repoconfig parameter:", err.Error())
			return
		}

		files := []common.File{}
		if err := json.Unmarshal([]byte(repoconfig.Content), &files); err != nil {
			logger.Error("failed to parse file parameter: %s", err.Error())
			response.Fail(c, "failed to parse file parameter:", err.Error())
			return
		}
		for i, v := range files {
			files[i].Content = base64.StdEncoding.EncodeToString([]byte(v.Content))
		}

		// 查询此配置的内容
		rf, err := service.GetRepoFileByInfoUUID(ci.UUID, nil)
		if err != nil {
			logger.Error("failed to get repoconfig parameter: %s", err.Error())
			response.Fail(c, "failed to get repoconfig parameter:", err.Error())
			return
		}

		// 更新参数
		repoconfig.UUID = rf.UUID
		repoconfig.ConfigInfoUUID = ci.UUID
		repoconfig.IsActive = false
		repoconfig.Content, err = json.Marshal(files)
		if err != nil {
			logger.Error("Error encoding JSON:: %s", err.Error())
			response.Fail(c, "Error encoding JSON:", err.Error())
			return
		}

		// 将参数添加到数据库
		err = repoconfig.Record()
		if err != nil {
			logger.Error("failed to update repoconfig: %s", err.Error())
			response.Fail(c, "failed to update repoconfig:", err.Error())
			return
		}
		logger.Debug("update repoconfig success")
		response.Success(c, nil, "update repo config success")

	case global.Host:
		// 解析参数
		hostconfig := &service.HostConfig{}
		if err := json.Unmarshal(query.Data, hostconfig); err != nil {
			logger.Error("failed to parse hostconfig parameter: %s", err.Error())
			response.Fail(c, "failed to parse hostconfig parameter:", err.Error())
			return
		}

		file := common.File{}
		if err := json.Unmarshal([]byte(hostconfig.Content), &file); err != nil {
			logger.Error("failed to parse file parameter: %s", err.Error())
			response.Fail(c, "failed to parse file parameter:", err.Error())
			return
		}
		file.Content = base64.StdEncoding.EncodeToString([]byte(file.Content))

		// 查询此配置的内容
		hf, err := service.GetHostFileByInfoUUID(ci.UUID, nil)
		if err != nil {
			logger.Error("failed to get hostconfig parameter: %s", err.Error())
			response.Fail(c, "failed to get hostconfig parameter:", err.Error())
			return
		}

		// 更新参数
		hostconfig.UUID = hf.UUID
		hostconfig.ConfigInfoUUID = ci.UUID
		hostconfig.IsActive = false
		hostconfig.Content, err = json.Marshal(file)
		if err != nil {
			logger.Error("Error encoding JSON:: %s", err.Error())
			response.Fail(c, "Error encoding JSON:", err.Error())
			return
		}

		// 将参数添加到数据库
		err = hostconfig.Record()
		if err != nil {
			logger.Error("failed to update hostconfig: %s", err.Error())
			response.Fail(c, "failed to update hostconfig:", err.Error())
			return
		}

		logger.Debug("update hostconfig success")
		response.Success(c, nil, "update host config success")

	case global.SSH:
		// 解析参数
		sshconfig := &service.SSHConfig{}
		if err := json.Unmarshal(query.Data, sshconfig); err != nil {
			logger.Error("failed to parse sshconfig parameter: %s", err.Error())
			response.Fail(c, "failed to parse sshconfig parameter:", err.Error())
			return
		}

		file := common.File{}
		if err := json.Unmarshal([]byte(sshconfig.Content), &file); err != nil {
			logger.Error("failed to parse file parameter: %s", err.Error())
			response.Fail(c, "failed to parse file parameter:", err.Error())
			return
		}
		file.Content = base64.StdEncoding.EncodeToString([]byte(file.Content))

		// 查询此配置的内容
		sf, err := service.GetSSHFileByInfoUUID(ci.UUID, nil)
		if err != nil {
			logger.Error("failed to get sshconfig parameter: %s", err.Error())
			response.Fail(c, "failed to get sshconfig parameter:", err.Error())
			return
		}

		// 更新参数
		sshconfig.UUID = sf.UUID
		sshconfig.ConfigInfoUUID = ci.UUID
		sshconfig.IsActive = false
		sshconfig.Content, err = json.Marshal(file)
		if err != nil {
			logger.Error("Error encoding JSON:: %s", err.Error())
			response.Fail(c, "Error encoding JSON:", err.Error())
			return
		}

		// 将参数添加到数据库
		err = sshconfig.Record()
		if err != nil {
			logger.Error("failed to update sshconfig: %s", err.Error())
			response.Fail(c, "failed to update sshconfig:", err.Error())
			return
		}

		logger.Debug("update sshconfig success")
		response.Success(c, nil, "update sshconfig success")

	case global.SSHD:
		// 解析参数
		sshdconfig := &service.SSHDConfig{}
		if err := json.Unmarshal(query.Data, sshdconfig); err != nil {
			logger.Error("failed to parse sshdconfig parameter: %s", err.Error())
			response.Fail(c, "failed to parse sshdconfig parameter:", err.Error())
			return
		}

		file := common.File{}
		if err := json.Unmarshal([]byte(sshdconfig.Content), &file); err != nil {
			logger.Error("failed to parse file parameter: %s", err.Error())
			response.Fail(c, "failed to parse file parameter:", err.Error())
			return
		}
		file.Content = base64.StdEncoding.EncodeToString([]byte(file.Content))

		// 查询此配置的内容
		sdf, err := service.GetSSHDFileByInfoUUID(ci.UUID, nil)
		if err != nil {
			logger.Error("failed to get sshdconfig parameter: %s", err.Error())
			response.Fail(c, "failed to get sshdconfig parameter:", err.Error())
			return
		}

		// 更新参数
		sshdconfig.UUID = sdf.UUID
		sshdconfig.ConfigInfoUUID = ci.UUID
		sshdconfig.IsActive = false
		sshdconfig.Content, err = json.Marshal(file)
		if err != nil {
			logger.Error("Error encoding JSON:: %s", err.Error())
			response.Fail(c, "Error encoding JSON:", err.Error())
			return
		}

		// 将参数添加到数据库
		err = sshdconfig.Record()
		if err != nil {
			logger.Error("failed to update sshdconfig: %s", err.Error())
			response.Fail(c, "failed to update sshdconfig:", err.Error())
			return
		}

		logger.Debug("update sshdconfig success")
		response.Success(c, nil, "update sshdconfig success")

	case global.Sysctl:
		// 解析参数
		sysctlconfig := &service.SysctlConfig{}
		if err := json.Unmarshal(query.Data, sysctlconfig); err != nil {
			logger.Error("failed to parse sysctlconfig parameter: %s", err.Error())
			response.Fail(c, "failed to parse sysctlconfig parameter:", err.Error())
			return
		}

		file := common.File{}
		if err := json.Unmarshal([]byte(sysctlconfig.Content), &file); err != nil {
			logger.Error("failed to parse file parameter: %s", err.Error())
			response.Fail(c, "failed to parse file parameter:", err.Error())
			return
		}
		file.Content = base64.StdEncoding.EncodeToString([]byte(file.Content))

		// 查询此配置的内容
		sysf, err := service.GetSysctlFileByInfoUUID(ci.UUID, nil)
		if err != nil {
			logger.Error("failed to get sysctlconfig parameter: %s", err.Error())
			response.Fail(c, "failed to get sysctlconfig parameter:", err.Error())
			return
		}

		// 更新参数
		sysctlconfig.UUID = sysf.UUID
		sysctlconfig.ConfigInfoUUID = ci.UUID
		sysctlconfig.IsActive = false
		sysctlconfig.Content, err = json.Marshal(file)
		if err != nil {
			logger.Error("Error encoding JSON:: %s", err.Error())
			response.Fail(c, "Error encoding JSON:", err.Error())
			return
		}

		// 将参数添加到数据库
		err = sysctlconfig.Record()
		if err != nil {
			logger.Error("failed to update sysctlconfig: %s", err.Error())
			response.Fail(c, "failed to update sysctlconfig:", err.Error())
			return
		}

		logger.Debug("update sysctlconfig success")
		response.Success(c, nil, "update sysctlconfig success")

	case global.DNS:
		// 解析参数
		dnsconfig := &service.DNSConfig{}
		if err := json.Unmarshal(query.Data, dnsconfig); err != nil {
			logger.Error("failed to parse dnsconfig parameter: %s", err.Error())
			response.Fail(c, "failed to parse dnsconfig parameter:", err.Error())
			return
		}

		file := common.File{}
		if err := json.Unmarshal([]byte(dnsconfig.Content), &file); err != nil {
			logger.Error("failed to parse file parameter: %s", err.Error())
			response.Fail(c, "failed to parse file parameter:", err.Error())
			return
		}
		file.Content = base64.StdEncoding.EncodeToString([]byte(file.Content))

		// 查询此配置的内容
		df, err := service.GetDNSFileByInfoUUID(ci.UUID, nil)
		if err != nil {
			logger.Error("failed to get dnsconfig parameter: %s", err.Error())
			response.Fail(c, "failed to get dnsconfig parameter:", err.Error())
			return
		}

		// 更新参数
		dnsconfig.UUID = df.UUID
		dnsconfig.ConfigInfoUUID = ci.UUID
		dnsconfig.IsActive = false
		dnsconfig.Content, err = json.Marshal(file)
		if err != nil {
			logger.Error("Error encoding JSON:: %s", err.Error())
			response.Fail(c, "Error encoding JSON:", err.Error())
			return
		}

		// 将参数添加到数据库
		err = dnsconfig.Record()
		if err != nil {
			logger.Error("failed to update dnsconfig: %s", err.Error())
			response.Fail(c, "failed to update dnsconfig:", err.Error())
			return
		}

		logger.Debug("update dnsconfig success")
		response.Success(c, nil, "update dnsconfig success")

	case global.PATH:
		// 解析参数
		pathconfig := &service.PathConfig{}
		if err := json.Unmarshal(query.Data, pathconfig); err != nil {
			logger.Error("failed to parse pathconfig parameter: %s", err.Error())
			response.Fail(c, "failed to parse pathconfig parameter:", err.Error())
			return
		}

		file := common.File{}
		if err := json.Unmarshal([]byte(pathconfig.Content), &file); err != nil {
			logger.Error("failed to parse file parameter: %s", err.Error())
			response.Fail(c, "failed to parse file parameter:", err.Error())
			return
		}
		file.Content = base64.StdEncoding.EncodeToString([]byte(file.Content))

		// 查询此配置的内容
		df, err := service.GetPathFileByInfoUUID(ci.UUID, nil)
		if err != nil {
			logger.Error("failed to get pathconfig parameter: %s", err.Error())
			response.Fail(c, "failed to get pathconfig parameter:", err.Error())
			return
		}

		// 更新参数
		pathconfig.UUID = df.UUID
		pathconfig.ConfigInfoUUID = ci.UUID
		pathconfig.IsActive = false
		pathconfig.Content, err = json.Marshal(file)
		if err != nil {
			logger.Error("Error encoding JSON:: %s", err.Error())
			response.Fail(c, "Error encoding JSON:", err.Error())
			return
		}

		// 将参数添加到数据库
		err = pathconfig.Record()
		if err != nil {
			logger.Error("failed to update pathconfig: %s", err.Error())
			response.Fail(c, "failed to update pathconfig:", err.Error())
			return
		}

		logger.Debug("update pathconfig success")
		response.Success(c, nil, "update pathconfig success")

	default:
		response.Fail(c, nil, "Unknown type of configinfo:"+query.UUID)
	}
}
