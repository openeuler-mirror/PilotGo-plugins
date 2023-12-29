package controller

import (
	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/atune-plugin/model"
	"openeuler.org/PilotGo/atune-plugin/service"
	"openeuler.org/PilotGo/atune-plugin/template"
)

func GetAtuneAll(c *gin.Context) {
	allData := []string{
		"compress",
		"compress_Except",
		"ffmpeg",
		"fio",
		"gcc_compile",
		"go_gc",
		"graphicsmagick",
		"iozone",
		"key_parameters_select",
		"key_parameters_select_variant",
		"mariadb",
		"memcached",
		"memory",
		"mysql_sysbench",
		"nginx",
		"openGauss",
		"redis",
		"spark",
		"tensorflow_train",
		"tidb",
		"tomcat"}
	response.Success(c, allData, "获取到全部的可调优业务名称")
}
func GetAtuneInfo(c *gin.Context) {
	tuneName := c.Query("name")
	tune := template.GetTuneInfo(tuneName)
	if tune == nil {
		response.Fail(c, nil, "未找到可调优的业务名称")
		return
	}
	response.Success(c, tune, "获取到调优步骤信息")
}

func QueryTunes(c *gin.Context) {
	query := &response.PaginationQ{}
	err := c.ShouldBindQuery(query)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	data, total, err := service.QueryTunes(query)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	response.DataPagination(c, data, total, query)
}

func SaveTune(c *gin.Context) {
	var t model.Tunes
	if err := c.Bind(&t); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	if err := service.SaveTune(t); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "已保存调优对象模板")
}

func UpdateTune(c *gin.Context) {
	var t model.Tunes
	if err := c.Bind(&t); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	if err := service.UpdateTune(t); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "已更新调优对象模板")
}

func DeleteTune(c *gin.Context) {
	tunedel := struct {
		TuneID []int `json:"ids"`
	}{}
	if err := c.Bind(&tunedel); err != nil {
		response.Fail(c, nil, "parameter error")
		return
	}

	if err := service.DeleteTune(tunedel.TuneID); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "已删除调优对象模板")
}

func SearchTune(c *gin.Context) {
	search := c.Query("search")

	query := &response.PaginationQ{}
	if err := c.ShouldBindQuery(query); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	data, total, err := service.SearchTune(search, query)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.DataPagination(c, data, total, query)
}
