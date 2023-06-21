package httphandler

import (
	"gitee.com/openeuler/PilotGo-plugins/sdk/common"
	"gitee.com/openeuler/PilotGo-plugins/sdk/response"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/redis-plugin/service"
)

// 安装运行
func InstallRedisExporter(c *gin.Context) {
	// TODOs
	var param *common.Batch

	if err := c.BindJSON(param); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	ret, err := service.Install(param)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, ret, "安装成功")
}

func UnInstallRedisExporter(c *gin.Context) {
	var param *common.Batch
	if err := c.BindJSON(param); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	ret, err := service.UnInstall(param)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, ret, "卸载成功")
}
