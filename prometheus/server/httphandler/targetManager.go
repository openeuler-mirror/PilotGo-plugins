package httphandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/prometheus-plugin/httphandler/service"
)

// 将监控target添加到prometheus插件db
func AddPrometheusTarget(c *gin.Context) {
	var target service.PrometheusTarget
	if err := c.BindJSON(&target); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"data": nil,
			"msg":  err.Error()})
		return
	}
	err := service.AddPrometheusTarget(target)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"data": nil,
			"msg":  err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": nil,
		"msg":  "添加成功"})
}

// 将监控target从prometheus插件db中删除
func DeletePrometheusTarget(c *gin.Context) {
	var target service.PrometheusTarget
	if err := c.BindJSON(&target); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"data": nil,
			"msg":  err.Error()})
		return
	}
	err := service.DeletePrometheusTarget(target)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"data": nil,
			"msg":  err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": nil,
		"msg":  "删除成功"})
}
