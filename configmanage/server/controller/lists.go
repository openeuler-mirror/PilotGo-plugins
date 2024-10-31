package controller

import (
	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/configmanage-plugin/global"
)

func ConfigTypeListHandler(c *gin.Context) {
	result := []string{global.Repo, global.Host, global.SSH, global.SSHD, global.Sysctl}
	response.Success(c, result, "get config type success")
}
