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
			c.Set("__internal__reverse_dest", dest)
			QueryRange(c)
		})
		pg.GET("/query_range", func(c *gin.Context) {
			c.Set("__internal__reverse_dest", dest)
			QueryRange(c)
		})
		pg.GET("/rules", func(c *gin.Context) {
			c.Set("__internal__reverse_dest", dest)
			QueryRange(c)
		})

	}

	return &plugin.Client{
		Router: router,
	}
}

func Query(c *gin.Context) {
	remote := c.GetString("__internal__reverse_dest")
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
	remote := c.GetString("__internal__reverse_dest")
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

func Rules(c *gin.Context) {
	remote := c.GetString("__internal__reverse_dest")
	if remote == "" {
		fmt.Println("get reverse dest failed!")
		return
	}
	url, err := url.Parse(remote)
	if err != nil {
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	c.Request.URL.Path = "api/v1/rules" //请求API
	proxy.ServeHTTP(c.Writer, c.Request)
}
