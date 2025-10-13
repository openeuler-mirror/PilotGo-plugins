package controller

import (
	"ant-agent/exec-script/model"
	"ant-agent/exec-script/service"

	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/pkg/response"
)

func ExecScript(c *gin.Context) {
	var script model.ScriptsRun
	if err := c.ShouldBindJSON(&script); err != nil {
		response.Fail(c, nil, "参数错误")
		return
	}
	result, err := service.ExecScript(&script)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, result, "脚本执行成功")
}
