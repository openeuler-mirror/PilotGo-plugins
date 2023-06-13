package httphandler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/redis-plugin/global"
	"openeuler.org/PilotGo/redis-plugin/plugin"
)

// 安装运行
func InstallRedisExporter(ctx *gin.Context) {
	// TODO
	param := &struct {
		Batch []string
	}{}
	if err := ctx.BindJSON(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   -1,
			"status": "redis",
		})
	}

	cmd := "yum install -y redis_exporter && systemctl start redis_exporter"
	cmdResults, err := global.GlobalClient.RunScript(param.Batch, cmd)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   -1,
			"status": fmt.Sprintf("run redis exporter error:%s", err),
		})
	}
	ret := []interface{}{}
	monitorTargets := []string{}
	for _, result := range cmdResults {
		d := struct {
			MachineUUID   string
			MachineIP     string
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
		} else {
			// TODO: add redis exporter to prometheus monitor target here
			// default exporter port :9121
			monitorTargets = append(monitorTargets, result.MachineIP+":9121")
		}

		ret = append(ret, d)
	}
	err = plugin.MonitorTargets(monitorTargets)
	if err != nil {
		fmt.Println("error: failed to add redis exporter to prometheus monitor targets")
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":   0,
		"status": "ok",
		"data":   ret,
	})
}
