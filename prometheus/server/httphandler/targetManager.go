package httphandler

import (
	"gitee.com/openeuler/PilotGo-plugins/sdk/response"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/prometheus-plugin/httphandler/service"
)

// 将监控target添加到prometheus插件db
func AddPrometheusTarget(c *gin.Context) {
	var target service.PrometheusTarget
	if err := c.BindJSON(&target); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	err := service.AddPrometheusTarget(target)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "添加成功")
}

// 将监控target从prometheus插件db中删除
func DeletePrometheusTarget(c *gin.Context) {
	var target service.PrometheusTarget
	if err := c.BindJSON(&target); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	err := service.DeletePrometheusTarget(target)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "删除成功")
}
