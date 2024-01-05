package controller

import (
	"fmt"
	"time"

	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/configmanage-plugin/service"
)

func AddRepoConfig(repoconfigs []service.RepoConfig, ciuuid string) error {
	version := fmt.Sprintf("v%s", time.Now().Format("2006-01-02-15-04-05"))
	for _, v := range repoconfigs {
		//文件信息存储
		config := &service.RepoConfig{
			UUID:    ciuuid,
			Name:    v.Name,
			File:    v.File,
			Path:    v.Path,
			Version: version,
		}
		err := config.Record()
		if err != nil {
			return err
		}
	}
	return nil
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

/*
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
*/
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
