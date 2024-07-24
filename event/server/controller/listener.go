package controller

import (
	"strconv"
	"strings"

	"gitee.com/openeuler/PilotGo/sdk/common"
	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-event/service"
)

func RegisterListenerHandler(c *gin.Context) {
	p := client.PluginInfo{}
	if err := c.ShouldBind(&p); err != nil {
		response.Fail(c, gin.H{"status": false}, err.Error())
		return
	}

	l := &service.Listener{
		Name: p.Name,
		URL:  p.Url,
	}
	service.AddListener(l)

	eventtypes := strings.Split(c.Query("eventTypes"), ",")
	for _, v := range eventtypes {
		eventtype, err := strconv.Atoi(v)
		if err != nil {
			response.Fail(c, gin.H{"status": false}, err.Error())
			return
		}
		service.AddEventMap(eventtype, l)
	}
	response.Success(c, gin.H{"status": "ok"}, "注册eventType成功")
}

func UnregisterListenerHandler(c *gin.Context) {
	p := client.PluginInfo{}
	if err := c.ShouldBind(&p); err != nil {
		response.Fail(c, gin.H{"status": false}, err.Error())
		return
	}
	l := &service.Listener{
		Name: p.Name,
		URL:  p.Url,
	}
	eventtypes := strings.Split(c.Query("eventTypes"), ",")
	for _, v := range eventtypes {
		eventtype, err := strconv.Atoi(v)
		if err != nil {
			response.Fail(c, gin.H{"status": false}, err.Error())
			return
		}
		service.RemoveEventMap(eventtype, l)
	}

	if !service.IsExitEventMap(l) {
		service.RemoveListener(l)
	}
	response.Success(c, gin.H{"status": "ok"}, "删除eventType成功")
}

func PublishEventHandler(c *gin.Context) {
	msg := &common.EventMessage{}
	if err := c.ShouldBind(msg); err != nil {
		response.Fail(c, gin.H{"status": false}, err.Error())
		return
	}

	service.PublishEvent(msg)
	response.Success(c, gin.H{"status": "ok"}, "publishEvent成功")
}
