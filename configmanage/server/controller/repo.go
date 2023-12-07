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
	query := struct {
		Description string   `json:"description"`
		BatchIds    []uint   `json:"batchids"`
		DepartIds   []int    `json:"departids"`
		Nodes       []string `json:"uuids"`
		Name        string   `json:"name"`
		File        string   `json:"file"`
	}{}
	err := c.ShouldBind(query)
	if err != nil {
		response.Fail(c, gin.H{"status": false}, err.Error())
		return
	}

	configuuid := uuid.New().String()
	config := &service.RepoConfig{
		UUID: configuuid,
		Name: query.Name,
		File: query.File,
	}
	err = config.Record()
	if err != nil {
		response.Fail(c, gin.H{"status": false}, err.Error())
		return
	}

	ci := service.ConfigInstance{
		UUID:        uuid.New().String(),
		Type:        global.Repo,
		Description: query.Description,
		BatchIds:    query.BatchIds,
		DepartIds:   query.DepartIds,
		Nodes:       query.Nodes,
		Config:      config,
	}
	err = ci.Add(configuuid)
	if err != nil {
		response.Fail(c, gin.H{"status": false}, err.Error())
		return
	}
	response.Success(c, nil, "Add repo config success")
}

func GetRepoConfig(c *gin.Context) {
	//TODO:query 类型需要转变需要包含configuuid
	query := c.Query("configuuid")
	config := &service.RepoConfig{
		UUID: query,
	}
	err := config.Load()
	if err != nil {
		response.Fail(c, gin.H{"status": false}, err.Error())
		return
	}
	response.Success(c, config, "Get repo file success")
}

func UpdateRepoConfig(c *gin.Context) {
	//TODO:query 类型需要转变需要包含configuuid
	query := service.RepoConfig{}
	err := c.ShouldBind(query)
	if err != nil {
		response.Fail(c, gin.H{"status": false}, err.Error())
		return
	}
	configuuid := query.UUID
	query.UUID = uuid.New().String()
	err = query.Record()
	if err != nil {
		response.Fail(c, gin.H{"status": false}, err.Error())
		return
	}

	err = query.UpdateRepoConfig(configuuid)
	if err != nil {
		response.Fail(c, gin.H{"status": false}, err.Error())
		return
	}
	response.Success(c, query, "Update repo file success")
}
