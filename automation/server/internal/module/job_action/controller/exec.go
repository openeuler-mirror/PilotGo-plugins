package controller

import (
	"fmt"

	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/job_action/service"
)

func ExecScriptHandler(c *gin.Context) {
	var sr struct {
		IPs        []string `json:"IPs"`
		ScriptID   string   `json:"script_id"`
		Params     string   `json:"params"`
		TimeOutSec int      `json:"timeoutSec"`
	}
	if err := c.ShouldBindJSON(&sr); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	err := service.ExecScript(sr.IPs, sr.ScriptID, sr.Params, sr.TimeOutSec)
	if err != nil {
		response.Fail(c, nil, fmt.Sprintf("下发脚本任务失败: %s", err.Error()))
		return
	}
	response.Success(c, nil, "success")
}
