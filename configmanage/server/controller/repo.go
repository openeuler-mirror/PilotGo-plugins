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
	query := &struct {
		Description string   `json:"description"`
		BatchIds    []uint   `json:"batchids"`
		DepartIds   []int    `json:"departids"`
		Nodes       []string `json:"uuids"`
		Name        string   `json:"name"`
		File        string   `json:"file"`
		Path        string   `json:"path"`
	}{}
	err := c.Bind(query)
	if err != nil {
		response.Fail(c, gin.H{"status": false}, err.Error())
		return
	}
	//文件信息存储
	configuuid := uuid.New().String()
	config := &service.RepoConfig{
		UUID: configuuid,
		Name: query.Name,
		File: query.File,
		Path: query.Path,
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

func HistoryRepoConfig(c *gin.Context) {
	//TODO:query 类型需要转变需要包含configuuid文件的uuid
	query := c.Query("configuuid")
	rcs, err := service.HistoryRepoConfig(query)
	if err != nil {
		response.Fail(c, gin.H{"status": false}, err.Error())
		return
	}
	response.Success(c, rcs, "Get repo last file success")
}

func RepoApply(c *gin.Context) {

	//TODO: query 类型需要转变
	query := &struct {
		Deploy_BatchIds  []int    `json:"deploy_batches"`
		Deploy_DepartIds []int    `json:"deploy_departs"`
		Deploy_NodeUUIds []string `json:"deploy_nodes"`
		DeployFile_UUID  string   `json:"file_uuid"` //文件uuid
	}{}
	err := c.Bind(query)
	if err != nil {
		response.Fail(c, gin.H{"status": false}, err.Error())
		return
	}

	/*configinfo, err := service.GetInfoByConfigUUID(query.FileBroadcast_UUID)
	if err != nil {
		response.Fail(c, gin.H{"status": false}, err.Error())
		return
	}*/
	repoconfigfile := &service.RepoConfig{
		UUID: query.DeployFile_UUID,
	}
	repoconfigfile.Load()
	de := service.Deploy{
		Deploy_BatchIds:  query.Deploy_BatchIds,
		Deploy_DepartIds: query.Deploy_DepartIds,
		Deploy_NodeUUIds: query.Deploy_NodeUUIds,
	}
	rcs, err := repoconfigfile.Apply(de)
	if err != nil {
		response.Fail(c, gin.H{"status": false}, err.Error())
		return
	}
	response.Success(c, rcs, "Get repo last file success")
}
