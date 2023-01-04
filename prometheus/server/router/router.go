package plugin

import (
	"fmt"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/plugin-sdk/plugin"
)

func DefaultClient(desc *plugin.PluginInfo) *plugin.Client {
	plugin.BaseInfo = desc
	dest := desc.ReverseDest

	router := gin.Default()
	mg := router.Group("plugin_manage/")
	{
		mg.GET("/info", plugin.InfoHandler)
	}

	pg := router.Group("/plugin/" + desc.Name)
	{
		pg.GET("/query", func(c *gin.Context) {
			c.Set("query", dest)
			Query(c)
		})
		pg.GET("/query_range", func(c *gin.Context) {
			c.Set("query_range", dest)
			QueryRange(c)
		})
		pg.GET("/targets", func(c *gin.Context) {
			c.Set("targets", dest)
			Targets(c)
		})

	}

	return &plugin.Client{
		Router: router,
	}
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
