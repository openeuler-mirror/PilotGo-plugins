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
	IndexUrl    string `json:"index_url"`
	ManageUrl   string `json:"manage_url"`
}

type Client struct {
	Router *gin.Engine
}

var BaseInfo *PluginInfo

func InfoHandler(c *gin.Context) {

	c.JSON(http.StatusOK, BaseInfo)
}

func ReverseProxyHandler(c *gin.Context) {
	fmt.Println("reverse to grafana server")
	remote := "http://10.1.167.104:3000/" //转向的host
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

	router := gin.Default()
	mg := router.Group("plugin_manage/")
	{
		mg.GET("/info", InfoHandler)
	}

	pg := router.Group("plugin/" + desc.Name)
	{
		pg.Any("/*any", ReverseProxyHandler)
	}

	return &Client{
		Router: router,
	}
}

func (c *Client) Serve() {
	// TODO: 启动http服务
	c.Router.Run()
}
