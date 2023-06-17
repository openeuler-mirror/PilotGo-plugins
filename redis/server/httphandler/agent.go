package httphandler

import (
	"net/http"

	"gitee.com/openeuler/PilotGo-plugins/sdk/common"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/redis-plugin/service"
)

// 安装运行
func InstallRedisExporter(ctx *gin.Context) {
	// TODOs
	var param *common.Batch

	if err := ctx.BindJSON(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   -1,
			"status": err,
		})
	}

	ret, err := service.Install(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   -1,
			"status": err,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":   0,
		"status": "ok",
		"data":   ret,
	})
}
