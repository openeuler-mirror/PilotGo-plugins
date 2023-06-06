package httphandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/container-plugin/client"
)

func PluginInfo(ctx *gin.Context) {
	plugin_info := client.Client().Plugin
	client.Client().Server = "http://" + ctx.Request.RemoteAddr

	ctx.JSON(http.StatusOK, plugin_info)
}
