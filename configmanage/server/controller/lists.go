/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: wubijie <wubijie@kylinos.cn>
 * Date: Thu Oct 31 16:45:12 2024 +0800
 */
package controller

import (
	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/configmanage-plugin/global"
	"openeuler.org/PilotGo/configmanage-plugin/service"
)

type PaginationQ struct {
	Ok             bool        `json:"ok"`
	Size           int         `form:"size" json:"size"`
	CurrentPageNum int         `form:"page" json:"page"`
	Data           interface{} `json:"data" comment:"muster be a pointer of slice gorm.Model"` // save pagination list
	TotalPage      int         `json:"total"`
}

func ConfigTypeListHandler(c *gin.Context) {
	result := []string{global.Repo, global.Host, global.SSH, global.SSHD, global.Sysctl, global.DNS, global.PATH}
	response.Success(c, result, "get config type success")
}

// 分页获取所有configinfo列表
func ConfigInfosHandler(c *gin.Context) {
	p := &PaginationQ{}
	// 将查询参数绑定到分页查询对象 p 中
	err := c.ShouldBindQuery(p)
	if err != nil {
		response.Fail(c, "parameter error", err.Error())
		return
	}
	num := p.Size * (p.CurrentPageNum - 1)
	total, data, err := service.GetInfos(num, p.Size)
	if err != nil {
		response.Fail(c, "parameter error", err.Error())
		return
	}
	p.Data = data
	p.TotalPage = total

	response.Success(c, p, "get config success")
}

// 根据列表中的configinfo_uuid获取某一个config信息
func ConfigInfoHandler(c *gin.Context) {
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
		logger.Error("failed to get configinfo: %s", err.Error())
		response.Fail(c, "get configinfo fail:", err.Error())
		return
	}
	// 获取对应配置管理的参数
	switch ci.Type {
	case global.Repo:
		// 获取有关配置的所有文件信息
		repofiles, err := service.GetRopeFilesByConfigUUID(ci.UUID)
		if err != nil {
			logger.Error("failed to get repoconfig files: %s", err.Error())
			response.Fail(c, "failed to get repoconfig files", err.Error())
			return
		}
		logger.Debug("load repoconfig success")
		response.Success(c, repofiles, "load repo config success")

	case global.Host:
		// 获取有关配置的所有文件信息
		hostfiles, err := service.GetHostFilesByConfigUUID(ci.UUID)
		if err != nil {
			logger.Error("failed to get hostconfig files: %s", err.Error())
			response.Fail(c, "failed to get hostconfig files", err.Error())
			return
		}
		logger.Debug("load hostconfig success")
		response.Success(c, hostfiles, "load hostconfig success")

	case global.SSH:
		// 获取有关配置的所有文件信息
		sshfiles, err := service.GetSSHFilesByConfigUUID(ci.UUID)
		if err != nil {
			logger.Error("failed to get sshconfig files: %s", err.Error())
			response.Fail(c, "failed to get sshconfig files", err.Error())
			return
		}
		logger.Debug("load sshconfig success")
		response.Success(c, sshfiles, "load sshconfig success")

	case global.SSHD:
		// 获取有关配置的所有文件信息
		sshdfiles, err := service.GetSSHDFilesByConfigUUID(ci.UUID)
		if err != nil {
			logger.Error("failed to get sshdconfig files: %s", err.Error())
			response.Fail(c, "failed to get sshdconfig files", err.Error())
			return
		}
		logger.Debug("load sshdconfig success")
		response.Success(c, sshdfiles, "load sshdconfig success")

	case global.Sysctl:
		// 获取有关配置的所有文件信息
		sysctlfiles, err := service.GetSysctlFilesByConfigUUID(ci.UUID)
		if err != nil {
			logger.Error("failed to get sysctlconfig files: %s", err.Error())
			response.Fail(c, "failed to get sysctlconfig files", err.Error())
			return
		}
		logger.Debug("load sysctlconfig success")
		response.Success(c, sysctlfiles, "load sysctlconfig success")

	case global.DNS:
		// 获取有关配置的所有文件信息GetDNSFilesByConfigUUID
		dnsfiles, err := service.GetDNSFilesByConfigUUID(ci.UUID)
		if err != nil {
			logger.Error("failed to get dnsfiles files: %s", err.Error())
			response.Fail(c, "failed to get dnsfiles files", err.Error())
			return
		}
		logger.Debug("load dnsfiles success")
		response.Success(c, dnsfiles, "load dnsfiles success")

	case global.PATH:
		// 获取有关配置的所有文件信息GetPathFilesByConfigUUID
		pathfiles, err := service.GetPathFilesByConfigUUID(ci.UUID)
		if err != nil {
			logger.Error("failed to get pathfiles files: %s", err.Error())
			response.Fail(c, "failed to get pathfiles files", err.Error())
			return
		}
		logger.Debug("load pathfiles success")
		response.Success(c, pathfiles, "load pathfiles success")

	default:
		response.Fail(c, nil, "Unknown type of configinfo:"+query.UUID)
	}
}

// 查看某台机器某种类型的的历史配置信息
func ConfigHistoryHandler(c *gin.Context) {
	//TODO:修改请求的参数
	query := &struct {
		UUID string `json:"node_uuid"`
		Type string `json:"type"`
	}{}
	err := c.ShouldBindJSON(query)
	if err != nil {
		response.Fail(c, "parameter error", err.Error())
		return
	}
	logger.Debug("load config")

	// 获取对应配置管理的参数
	switch query.Type {
	case global.Repo:
		// 获取有关本台机器配置的所有文件信息
		repofiles, err := service.GetRopeFilesByNode(query.UUID)
		if err != nil {
			logger.Error("failed to get repoconfig files: %s", err.Error())
			response.Fail(c, "failed to get repoconfig files", err.Error())
			return
		}
		logger.Debug("load repoconfig success")
		response.Success(c, repofiles, "load repo config success")

	case global.Host:
		// 获取有关本台机器配置的所有文件信息
		hostconfig, err := service.GetHostFilesByNode(query.UUID)
		if err != nil {
			logger.Error("failed to get hostconfig files: %s", err.Error())
			response.Fail(c, "failed to get hostconfig files", err.Error())
			return
		}
		logger.Debug("load hostconfig success")
		response.Success(c, hostconfig, "load hostconfig success")

	case global.SSH:
		// 获取有关本台机器配置的所有文件信息
		sshconfig, err := service.GetSSHFilesByNode(query.UUID)
		if err != nil {
			logger.Error("failed to get sshconfig files: %s", err.Error())
			response.Fail(c, "failed to get sshconfig files", err.Error())
			return
		}
		logger.Debug("load sshconfig success")
		response.Success(c, sshconfig, "load sshconfig success")

	case global.SSHD:
		// 获取有关本台机器配置的所有文件信息
		sshdconfig, err := service.GetSSHDFilesByNode(query.UUID)
		if err != nil {
			logger.Error("failed to get sshdconfig files: %s", err.Error())
			response.Fail(c, "failed to get sshdconfig files", err.Error())
			return
		}
		logger.Debug("load sshdconfig success")
		response.Success(c, sshdconfig, "load sshdconfig success")

	case global.Sysctl:
		// 获取有关本台机器配置的所有文件信息
		sysctlfiles, err := service.GetSysctlFilesByNode(query.UUID)
		if err != nil {
			logger.Error("failed to get sysctlconfig files: %s", err.Error())
			response.Fail(c, "failed to get sysctlconfig files", err.Error())
			return
		}
		logger.Debug("load sysctlconfig success")
		response.Success(c, sysctlfiles, "load sysctlconfig success")

	case global.DNS:
		// 获取有关本台机器配置的所有文件信息
		dnsconfig, err := service.GetDNSFilesByNode(query.UUID)
		if err != nil {
			logger.Error("failed to get dnsconfig files: %s", err.Error())
			response.Fail(c, "failed to get dnsconfig files", err.Error())
			return
		}
		logger.Debug("load dnsconfig success")
		response.Success(c, dnsconfig, "load dnsconfig success")

	case global.PATH:
		// 获取有关本台机器配置的所有文件信息
		pathconfig, err := service.GetPathFilesByNode(query.UUID)
		if err != nil {
			logger.Error("failed to get pathconfig files: %s", err.Error())
			response.Fail(c, "failed to get pathconfig files", err.Error())
			return
		}
		logger.Debug("load pathconfig success")
		response.Success(c, pathconfig, "load pathconfig success")

	default:
		response.Fail(c, nil, "Unknown type of configinfo:"+query.UUID)
	}
}
