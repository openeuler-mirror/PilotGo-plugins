package httphandler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/gala-ops-plugin/pluginclient"
)

func InstallGopher(ctx *gin.Context) {
	// TODO
	param := &struct {
		Batch []string
	}{}
	if err := ctx.BindJSON(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   -1,
			"status": "parameter error",
		})
	}

	cmd := "yum install -y gala-gopher"
	cmdResults, err := pluginclient.Client().RunScript(param.Batch, cmd)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   -1,
			"status": fmt.Sprintf("run remote script error:%s", err),
		})
	}

	ret := []interface{}{}
	for _, result := range cmdResults {
		d := struct {
			MachineUUID   string
			InstallStatus string
			Error         string
		}{
			MachineUUID:   result.MachineUUID,
			InstallStatus: "ok",
			Error:         "",
		}

		if result.Code != 0 {
			d.InstallStatus = "error"
			d.Error = result.Stderr
		}

		ret = append(ret, d)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":   0,
		"status": fmt.Sprintf("run remote script error:%s", err),
		"data":   ret,
	})
}

func UpgradeGopher(ctx *gin.Context) {
	// TODO
}

func UninstallGopher(ctx *gin.Context) {
	// TODO
}
