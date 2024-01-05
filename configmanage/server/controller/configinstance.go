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
