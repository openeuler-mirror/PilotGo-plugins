package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"

	"gitee.com/openeuler/PilotGo-plugin-elk/conf"
	"gitee.com/openeuler/PilotGo-plugin-elk/errormanager"
	"gitee.com/openeuler/PilotGo-plugin-elk/pluginclient"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func InitWebServer() {
	if pluginclient.Global_Client == nil {
		err := errors.New("Global_Client is nil **errstackfatal**2") // err top
		errormanager.ErrorTransmit(pluginclient.Global_Context, err, true)
		return
	}

	go func() {
		engine := gin.Default()
		gin.SetMode(gin.ReleaseMode)
		pluginclient.Global_Client.RegisterHandlers(engine)
		InitRouter(engine)
		StaticRouter(engine)

		if conf.Global_Config.Elk.Https_enabled {
			err := engine.RunTLS(conf.Global_Config.Elk.Addr, conf.Global_Config.Elk.Public_certificate, conf.Global_Config.Elk.Private_key)
			if err != nil {
				err = errors.Errorf("%s **errstackfatal**2", err.Error()) // err top
				errormanager.ErrorTransmit(pluginclient.Global_Context, err, true)
			}
		} else {
			err := engine.Run(conf.Global_Config.Elk.Addr)
			if err != nil {
				err = errors.Errorf("%s **errstackfatal**2", err.Error()) // err top
				errormanager.ErrorTransmit(pluginclient.Global_Context, err, true)
			}
		}
	}()
}

func InitRouter(router *gin.Engine) {
	api := router.Group("/plugin/elk/api")
	{
		api.POST("/create_policy", CreatePolicyHandle)
		api.POST("/search", SearchHandle)
	}

	timeoutapi := router.Group("/plugin/elk/api")
	timeoutapi.Use(TimeoutMiddleware2(15 * time.Second))
	{

	}
}

func TimeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(12*time.Second),
		timeout.WithHandler(func(ctx *gin.Context) {
			ctx.Next()
		}),
		timeout.WithResponse(func(ctx *gin.Context) {
			ctx.JSON(http.StatusGatewayTimeout, gin.H{
				"code": http.StatusGatewayTimeout,
				"msg":  "server response timeout",
				"data": nil,
			})
		}),
	)
}

// 服务器响应超时中间件
func TimeoutMiddleware2(timeout time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer func() {
			if !c.GetBool("write") && ctx.Err() == context.DeadlineExceeded {
				c.Writer.WriteHeader(http.StatusGatewayTimeout)
				c.Abort()
			}

			cancel()
		}()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
