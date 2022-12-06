package plugin

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

type PluginInfo struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Email       string `json:"email"`
	Url         string `json:"url"`
	ReverseDest string `json:"reverse_dest"`
}

type Client struct {
	Router *gin.Engine
}

var BaseInfo *PluginInfo

func InfoHandler(c *gin.Context) {

	c.JSON(http.StatusOK, BaseInfo)
}

func ReverseProxyHandler(c *gin.Context) {
	remote := c.GetString("__internal__reverse_dest")
	if remote == "" {
		fmt.Println("get reverse dest failed!")
		return
	}

	target, err := url.Parse(remote)
	if err != nil {
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	c.Request.URL.Path = strings.Replace(c.Request.URL.Path, "/plugin/grafana", "", 1) //请求API

	proxy.ServeHTTP(c.Writer, c.Request)
}

func DefaultClient(desc *PluginInfo) *Client {
	BaseInfo = desc
	dest := desc.ReverseDest

	router := gin.Default()
	mg := router.Group("plugin_manage/")
	{
		mg.GET("/info", InfoHandler)
	}

	pg := router.Group("/plugin/" + desc.Name)
	{
		pg.Any("/*any", func(c *gin.Context) {
			c.Set("__internal__reverse_dest", dest)
			ReverseProxyHandler(c)
		})
	}

	return &Client{
		Router: router,
	}
}

func (c *Client) Serve(url ...string) {
	// TODO: 启动http服务
	c.Router.Run(url...)
}
