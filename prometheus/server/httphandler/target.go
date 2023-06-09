package httphandler

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/prometheus-plugin/httphandler/service"
)

func DBTargets(c *gin.Context) {
	targets, err := service.GetPrometheusTarget()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"data": nil,
			"msg":  err.Error()})
		return
	}

	objs := []PrometheusObject{
		{
			Targets: targets,
		},
	}
	c.JSON(http.StatusOK, objs)
}

func PrometheusAPITargets(c *gin.Context) {
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

type PrometheusObject struct {
	Targets []string `json:"targets"`
}
