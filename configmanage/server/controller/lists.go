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
	result := []string{global.Repo, global.Host, global.SSH, global.SSHD, global.Sysctl}
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
		repofiles, err := service.GetRopeFilesByCinfigUUID(ci.UUID)
		if err != nil {
			logger.Error("failed to get repoconfig file:s %s", err.Error())
			response.Fail(c, "failed to get repoconfig files", err.Error())
			return
		}
		logger.Debug("load repoconfig success")
		response.Success(c, repofiles, "load repo config success")

	case global.Host:
		// 获取有关配置的所有文件信息
		repofiles, err := service.GetHostFilesByConfigUUID(ci.UUID)
		if err != nil {
			logger.Error("failed to get hostconfig file:s %s", err.Error())
			response.Fail(c, "failed to get hostconfig files", err.Error())
			return
		}
		logger.Debug("load hostconfig success")
		response.Success(c, repofiles, "load hostconfig success")
	case global.SSH:

	case global.SSHD:

	case global.Sysctl:

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
			logger.Error("failed to get repoconfig file:s %s", err.Error())
			response.Fail(c, "failed to get repoconfig files", err.Error())
			return
		}
		logger.Debug("load repoconfig success")
		response.Success(c, repofiles, "load repo config success")

	case global.Host:

	case global.SSH:

	case global.SSHD:

	case global.Sysctl:

	default:
		response.Fail(c, nil, "Unknown type of configinfo:"+query.UUID)
	}
}
