package handler

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"gitee.com/openeuler/PilotGo-plugin-topology-server/agentmanager"
	"gitee.com/openeuler/PilotGo-plugin-topology-server/conf"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func InitWebServer() {
	engine := gin.Default()
	agentmanager.Topo.Sdkmethod.RegisterHandlers(engine)
	InitRouter(engine)
	err := engine.Run(conf.Config().Topo.Server_addr)
	if err != nil {
		fmt.Printf("%+v\n", errors.Errorf("%s**2", err.Error())) // err top
		os.Exit(1)
	}
}

func InitRouter(router *gin.Engine) {
	api := router.Group("/plugin/topology/api")
	{
		api.GET("/single_host/:uuid", SingleHostHandle)
		api.GET("/multi_host", MultiHostHandle)
	}

	static := router.Group("/plugin/topology/")
	{
		// static.Static("/assets", "./frontend/assets")
		// static.StaticFile("/", "./frontend/index.html")
		static.Static("/assets", "../web/dist/assets")
		static.StaticFile("/", "../web/dist/index.html")

		// 解决页面刷新404的问题
		router.NoRoute(func(c *gin.Context) {
			if !strings.HasPrefix(c.Request.RequestURI, "/plugin/topology/api/") {
				// c.File("./frontend/index.html")
				c.File("../web/dist/index.html")
				return
			}
			c.AbortWithStatus(http.StatusNotFound)
		})
	}
}
