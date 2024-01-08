package controller

import (
	"encoding/json"
	"fmt"

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
		response.Fail(c, gin.H{"status": false}, err.Error())
		return
	}

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
		response.Fail(c, gin.H{"status": false}, err.Error())
		return
	}

	//解析对应配置管理的参数
	switch query.Type {
	case global.Repo:
		//解析参数
		var repoconfigs []service.RepoConfig
		if err := json.Unmarshal(query.Data, &repoconfigs); err != nil {
			response.Fail(c, gin.H{"status": false}, err.Error())
			return
		}

		//将参数添加到数据库
		err = AddRepoConfig(repoconfigs, ci.UUID)
		if err != nil {
			response.Fail(c, gin.H{"status": false}, err.Error())
			return
		}
		response.Success(c, nil, "Add repo config success")

	case global.Host:

	case global.SSH:

	case global.SSHD:

	case global.Sysctl:

	default:
		fmt.Println("Unknown type:", query.Type)
	}
}

func ApplyConfigHandler(c *gin.Context) {
	query := &struct {
		Deploy_Type      string          `json:"deploy_type"`
		Deploy_BatchIds  []int           `json:"deploy_batches"`
		Deploy_DepartIds []int           `json:"deploy_departs"`
		Deploy_NodeUUIds []string        `json:"deploy_nodes"`
		Deploy_Data      json.RawMessage `json:"deploy_data"`
	}{}
	err := c.ShouldBindJSON(query)
	if err != nil {
		response.Fail(c, gin.H{"status": false}, err.Error())
		return
	}

	//解析对应配置管理的参数
	switch query.Deploy_Type {
	case global.Repo:
		//TODO:解析参数,data可以传输文件uuid,其余信息从数据库查询
		var repoconfig service.RepoConfig
		if err := json.Unmarshal(query.Deploy_Data, &repoconfig); err != nil {
			response.Fail(c, gin.H{"status": false}, err.Error())
			return
		}

		de := service.Deploy{
			Deploy_BatchIds:  query.Deploy_BatchIds,
			Deploy_DepartIds: query.Deploy_DepartIds,
			Deploy_NodeUUIds: query.Deploy_NodeUUIds,
			Deploy_Path:      repoconfig.Path,
			Deploy_FileName:  repoconfig.Name,
			Deploy_Text:      repoconfig.File,
		}
		//将参数添加到数据库
		result, err := repoconfig.Apply(de)
		if err != nil {
			response.Fail(c, result, err.Error())
			return
		}
		response.Success(c, nil, "Add repo config success")

	case global.Host:

	case global.SSH:

	case global.SSHD:

	case global.Sysctl:

	default:
		fmt.Println("Unknown type:", query.Deploy_Type)
	}
}
