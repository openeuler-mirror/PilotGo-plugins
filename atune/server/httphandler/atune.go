package httphandler

import (
	"gitee.com/openeuler/PilotGo-plugins/sdk/response"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/atune-plugin/templete"
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
	tune := templete.GetTuneInfo(tuneName)
	if tune == nil {
		response.Fail(c, nil, "未找到可调优的业务名称")
		return
	}
	response.Success(c, tune, "获取到调优步骤信息")
}
