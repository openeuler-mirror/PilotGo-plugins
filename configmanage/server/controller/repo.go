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
