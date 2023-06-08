package router

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/prometheus-plugin/global"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(logger.LoggerDebug())
	router.Use(gin.Recovery())

	return router
}

func RegisterAPIs(router *gin.Engine) {
	global.GlobalClient.RegisterHandlers(router)

	pg := router.Group("/plugin/" + global.GlobalClient.PluginInfo.Name)
	{
		pg.GET("/query", func(c *gin.Context) {
			c.Set("query", global.GlobalClient.PluginInfo.ReverseDest)
			Query(c)
		})
		pg.GET("/query_range", func(c *gin.Context) {
			c.Set("query_range", global.GlobalClient.PluginInfo.ReverseDest)
			QueryRange(c)
		})
		pg.GET("/targets", func(c *gin.Context) {
			c.Set("targets", global.GlobalClient.PluginInfo.ReverseDest)
			Targets(c)
		})

	}
	router.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })
}

func Query(c *gin.Context) {
	remote := c.GetString("query")
	if remote == "" {
		fmt.Println("get reverse dest failed!")
		return
	}
	url, err := url.Parse(remote)
	if err != nil {
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	c.Request.URL.Path = "api/v1/query" //请求API
	proxy.ServeHTTP(c.Writer, c.Request)
}

func QueryRange(c *gin.Context) {
	remote := c.GetString("query_range")
	if remote == "" {
		fmt.Println("get reverse dest failed!")
		return
	}
	url, err := url.Parse(remote)
	if err != nil {
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	c.Request.URL.Path = "api/v1/query_range" //请求API
	proxy.ServeHTTP(c.Writer, c.Request)
}

func Targets(c *gin.Context) {
	remote := c.GetString("targets")
	if remote == "" {
		fmt.Println("get reverse dest failed!")
		return
	}
	url, err := url.Parse(remote)
	if err != nil {
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	c.Request.URL.Path = "api/v1/targets" //请求API
	proxy.ServeHTTP(c.Writer, c.Request)
}
