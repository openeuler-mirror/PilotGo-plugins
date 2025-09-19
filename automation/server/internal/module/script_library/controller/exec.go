package controller

import (
	"gitee.com/openeuler/PilotGo/sdk/common"
	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/global"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/pkg/utils"
)

func ExecScriptHandler(c *gin.Context) {
	var sr struct {
		UUIDS         []string `json:"UUIDS"`
		ScriptType    string   `json:"script_type"`
		ScriptContent string   `json:"script_content"`
		Params        string   `json:"params"`
		TimeOutSec    int      `json:"timeoutSec"`
	}
	if err := c.ShouldBindJSON(&sr); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	_, err := global.App.Client.MachineList()
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	result, err := global.App.Client.AgentRunScripts(&common.Batch{MachineUUIDs: sr.UUIDS}, sr.ScriptType, utils.EncodeScriptContent(sr.ScriptContent), sr.Params, sr.TimeOutSec)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, result, "success")
}
