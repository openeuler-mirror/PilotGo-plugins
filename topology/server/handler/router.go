package handler

import (
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
	StaticRouter(engine)
	err := engine.Run(conf.Config().Topo.Server_addr)
	if err != nil {
		err = errors.Errorf("%s **fatal**2", err.Error()) // err top
		agentmanager.Topo.ErrCh <- err
		agentmanager.Topo.ErrGroup.Add(1)
		agentmanager.Topo.ErrGroup.Wait()
		close(agentmanager.Topo.ErrCh)
		os.Exit(1)
	}
}

func InitRouter(router *gin.Engine) {
	api := router.Group("/plugin/topology/api")
	{
		api.GET("/agentlist", AgentListHandle)

		api.GET("/single_host/:uuid", SingleHostHandle)
		api.GET("/single_host_tree/:uuid", SingleHostTreeHandle)

		api.GET("/multi_host", MultiHostHandle)
	}
}

func StaticRouter(router *gin.Engine) {
	static := router.Group("/plugin/topology")
	{
		// static.Static("/assets", "./frontend/assets")
		// static.StaticFile("/", "./frontend/index.html")
		static.Static("/assets", "./dist/assets")
		static.StaticFile("/", "./dist/index.html")

		// 解决页面刷新404的问题
		router.NoRoute(func(c *gin.Context) {
			if strings.HasPrefix(c.Request.RequestURI, "/plugin/topology") {
				// c.File("./frontend/index.html")
				c.File("./dist/index.html")
				return
			}
			c.AbortWithStatus(http.StatusNotFound)
		})
	}
}
