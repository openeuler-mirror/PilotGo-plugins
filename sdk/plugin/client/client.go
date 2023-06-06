package client

import (
	"github.com/gin-gonic/gin"
)

type Client struct {
	Server     string
	PluginInfo *PluginInfo
}

var BaseInfo *PluginInfo

func DefaultClient(desc *PluginInfo) *Client {
	BaseInfo = desc

	return &Client{
		PluginInfo: desc,
	}
}

// RegisterHandlers 注册一些插件标准的API接口，清单如下：
// GET /plugin_manage/info
func (c *Client) RegisterHandlers(router *gin.Engine) {
	// 提供插件基本信息
	mg := router.Group("plugin_manage/")
	{
		mg.GET("/info", InfoHandler)
	}

	// pg := router.Group("/plugin/" + desc.Name)
	// {
	// 	pg.Any("/*any", func(c *gin.Context) {
	// 		c.Set("__internal__reverse_dest", dest)
	// 		ReverseProxyHandler(c)
	// 	})
	// }
}
