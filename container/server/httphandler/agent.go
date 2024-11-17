package httphandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeployDocker(ctx *gin.Context) {
	// 定义请求参数结构
	param := &struct {
		Batch []string `json:"batch" binding:"required"`
	}{}

	// 参数绑定和验证
	if err := ctx.BindJSON(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"status":  "parameter error",
			"message": err.Error(),
		})
		return // 添加return避免继续执行
	}

	// 参数校验
	if len(param.Batch) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"status":  "parameter error",
			"message": "batch cannot be empty",
		})
		return
	}

	// TODO: 实现Docker部署逻辑
	ret := ""

	// 返回成功响应
	ctx.JSON(http.StatusOK, gin.H{
		"code":   0,
		"status": "ok",
		"data":   ret,
	})
}
