//go:build production
// +build production

package handler

import (
	"embed"
	"io/fs"
	"mime"
	"net/http"
	"strings"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"github.com/gin-gonic/gin"
)

//go:embed assets index.html
var StaticFiles embed.FS

func StaticRouter(router *gin.Engine) {
	sf, err := fs.Sub(StaticFiles, "assets")
	if err != nil {
		logger.Warn(err.Error())
		return
	}

	mime.AddExtensionType(".js", "application/javascript")
	static := router.Group("/plugin/template")
	{
		static.StaticFS("/assets", http.FS(sf))
		static.GET("/", func(c *gin.Context) {
			c.FileFromFS("/", http.FS(StaticFiles))
		})

	}

	// 解决页面刷新404的问题
	router.NoRoute(func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.RequestURI, "/plugin/template/api") {
			c.FileFromFS("/", http.FS(StaticFiles))
			return
		}
		c.AbortWithStatus(http.StatusNotFound)
	})

}
