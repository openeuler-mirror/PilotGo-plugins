package httphandler

import (
	"fmt"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

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

func Alerts(c *gin.Context) {
	remote := c.GetString("alerts")
	if remote == "" {
		fmt.Println("get reverse dest failed!")
		return
	}
	url, err := url.Parse(remote)
	if err != nil {
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	c.Request.URL.Path = "api/v1/alerts" //请求API
	proxy.ServeHTTP(c.Writer, c.Request)
}
