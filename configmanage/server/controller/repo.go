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

	configuuid := uuid.New().String()
	config := &service.RepoConfig{
		UUID: configuuid,
		Name: "",
		File: "",
	}
	err = config.Record()
	if err != nil {
		response.Fail(c, gin.H{"status": false}, err.Error())
		return
	}

	ci := service.ConfigInstance{
		UUID:        uuid.New().String(),
		Type:        global.Repo,
		Description: "",
		BatchIds:    []uint{},
		DepartIds:   []int{},
		Nodes:       []string{},
		Config:      config,
	}
	err = ci.Add(configuuid)
	if err != nil {
		response.Fail(c, gin.H{"status": false}, err.Error())
		return
	}
}

func GetRepoConfig(c *gin.Context) {
	//TODO:query 类型需要转变需要包含uuid
	var query int
	err := c.ShouldBind(query)
	if err != nil {
		response.Fail(c, gin.H{"status": false}, err.Error())
		return
	}
	config := &service.RepoConfig{
		UUID: "configuuid",
	}
	err = config.Load()
	if err != nil {
		response.Fail(c, gin.H{"status": false}, err.Error())
		return
	}

}
