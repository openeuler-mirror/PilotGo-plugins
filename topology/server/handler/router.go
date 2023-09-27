package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/timeout"

	"gitee.com/openeuler/PilotGo-plugin-topology-server/agentmanager"
	"gitee.com/openeuler/PilotGo-plugin-topology-server/conf"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func InitWebServer() {
	engine := gin.Default()
	engine.Use(TimeoutMiddleware())
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

func TimeoutResponse(ctx *gin.Context) {
	ctx.JSON(http.StatusGatewayTimeout, gin.H{
		"code":  http.StatusGatewayTimeout,
		"error": "timeout",
		"data":  nil,
	})
}

func TimeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(600*time.Second),
		timeout.WithHandler(func(ctx *gin.Context) {
			ctx.Next()
		}),
		timeout.WithResponse(TimeoutResponse),
	)
}
