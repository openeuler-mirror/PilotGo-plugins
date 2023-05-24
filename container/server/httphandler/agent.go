package httphandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeployDocker(ctx *gin.Context) {
	param := &struct {
		Batch []string
	}{}
	if err := ctx.BindJSON(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   -1,
			"status": "parameter error",
		})
	}

	ret := ""
	ctx.JSON(http.StatusOK, gin.H{
		"code":   0,
		"status": "ok",
		"data":   ret,
	})
}
