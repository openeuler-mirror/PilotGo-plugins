package controller

import (
	"encoding/json"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"openeuler.org/PilotGo/configmanage-plugin/global"
	"openeuler.org/PilotGo/configmanage-plugin/service"
)

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
		Type:        global.Repo,
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

		repoconfig.UUID = uuid.New().String()
		repoconfig.ConfigInfoUUID = ci.UUID
		repoconfig.IsActive = false

		//将参数添加到数据库
		err = repoconfig.Record()
		if err != nil {
			logger.Error("failed to add repoconfig: %s", err.Error())
			response.Fail(c, "failed to add repoconfig:", err.Error())
			return
		}
		//收集机器的配置信息
		err = repoconfig.Collect()
		if err != nil {
			logger.Error("failed to collect repofile: %s", err.Error())
			response.Fail(c, "failed to collect repofile:", err.Error())
			return
		}
		logger.Debug("add repoconfig success")
		response.Success(c, nil, "Add repo config success")

	case global.Host:

	case global.SSH:

	case global.SSHD:

	case global.Sysctl:

	default:
		response.Fail(c, nil, "Unknown type:"+query.Type)
	}
}

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
		//加载配置
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

	case global.SSH:

	case global.SSHD:

	case global.Sysctl:

	default:
		response.Fail(c, nil, "Unknown type of configinfo:"+query.UUID)
	}
}

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

	case global.SSH:

	case global.SSHD:

	case global.Sysctl:

	default:
		response.Fail(c, nil, "Unknown type of configinfo:"+query.UUID)
	}
}
