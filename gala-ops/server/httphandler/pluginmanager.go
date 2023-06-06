package httphandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/gala-ops-plugin/client"
)

func PluginInfo(ctx *gin.Context) {
	plugin_info := client.Client().Plugin
	client.Client().Server = "http://" + ctx.Request.RemoteAddr
	
	ctx.JSON(http.StatusOK, gin.H{
		"code":   0,
		"status": "ok",
		"data":   plugin_info,
	})
}
