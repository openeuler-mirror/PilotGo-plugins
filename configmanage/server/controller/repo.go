package controller

import (
	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"openeuler.org/PilotGo/configmanage-plugin/global"
	"openeuler.org/PilotGo/configmanage-plugin/service"
)

func AddRepoHandler(c *gin.Context) {
	//TODO:query 类型需要转变
	var query int
	err := c.ShouldBind(query)
	if err != nil {
		response.Fail(c, gin.H{"status": false}, err.Error())
		return
	}

	uuid := uuid.New().String()
	config := &service.RepoConfig{
		UUID: uuid,
	}
	err = config.Record()
	if err != nil {
		response.Fail(c, gin.H{"status": false}, err.Error())
		return
	}

	ci := service.ConfigInstance{
		UUID:        uuid,
		Type:        global.Repo,
		Description: "",
		BatchIds:    []uint{},
		DepartIds:   []int{},
		UUIDS:       []string{},
		Config:      config,
	}
	err = ci.Record()
	if err != nil {
		response.Fail(c, gin.H{"status": false}, err.Error())
		return
	}
}
