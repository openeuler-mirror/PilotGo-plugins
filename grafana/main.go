package main

import (
	"fmt"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"gitee.com/openeuler/PilotGo-plugins/sdk/plugin/client"
	"github.com/gin-gonic/gin"
)

const Version = "0.0.1"

var PluginInfo = &client.PluginInfo{
	Name:        "grafana",
	Version:     Version,
	Description: "grafana可视化工具支持",
	Author:      "guozhengxin",
	Email:       "guozhengxin@kylinos.cn",
	Url:         "http://10.41.121.71:9999/plugin/grafana",
	ReverseDest: "http://10.41.121.71:3000/",
}

func main() {
	fmt.Println("hello grafana")

	InitLogger()

	server := gin.Default()

	client := client.DefaultClient(PluginInfo)
	client.RegisterHandlers(server)

	if err := server.Run(":9999"); err != nil {
		logger.Fatal("failed to run server")
	}
}

func InitLogger() {
	// TODO: use config in file
	conf := &logger.LogOpts{
		Level:   "debug",
		Driver:  "stdio",
		Path:    "./log",
		MaxFile: 10,
		MaxSize: 1024 * 1024 * 30,
	}

	if err := logger.Init(conf); err != nil {
		fmt.Printf("logger init failed, please check the config file: %s", err)
		os.Exit(-1)
	}
}

func InitRouter(router *gin.Engine) {
	// 所有grafana请求反向代理到grafana服务器
	pg := router.Group("/plugin/" + PluginInfo.Name)
	{
		pg.Any("/*any", func(c *gin.Context) {
			c.Set("__internal__reverse_dest", PluginInfo.ReverseDest)
			ReverseProxyHandler(c)
		})
	}
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
