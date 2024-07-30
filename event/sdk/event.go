package sdk

import (
	"encoding/json"
	"errors"
	"io"

	"gitee.com/openeuler/PilotGo/sdk/common"
	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
	"github.com/gin-gonic/gin"
)

func RegisterEventHandlers(router *gin.Engine, c *client.Client) {

	api := router.Group("/plugin_manage/api/v1/")
	{
		api.POST("/event", eventHandler)
	}
	plugin_client = c
	startEventProcessor(c)
}

func eventHandler(c *gin.Context) {
	j, err := io.ReadAll(c.Request.Body) // 接收数据
	if err != nil {
		logger.Error("没获取到：%s", err.Error())
		return
	}
	var msg common.EventMessage
	if err := json.Unmarshal(j, &msg); err != nil {
		logger.Error("反序列化结果失败%s", err.Error())
		return
	}

	ProcessEvent(&msg)
}

func eventPluginServer() (string, error) {
	plugins, err := plugin_client.GetPlugins()
	if err != nil {
		return "", err
	}

	var eventServer string
	for _, p := range plugins {
		if p.Name == "event" {
			eventServer = p.Url
			break
		}
	}

	if eventServer == "" {
		return "", errors.New("event plugin not found")
	}

	return eventServer, nil
}

func registerEventCallback(eventType int, callback common.EventCallback) {
	plugin_client.EventCallbackMap[eventType] = callback
}

func unregisterEventCallback(eventType int) {
	delete(plugin_client.EventCallbackMap, eventType)
}

func ProcessEvent(event *common.EventMessage) {
	plugin_client.EventChan <- event
}

func startEventProcessor(c *client.Client) {
	go func(c *client.Client) {
		for {
			e := <-c.EventChan

			// TODO: process event message
			cb, ok := c.EventCallbackMap[e.MessageType]
			if ok {
				cb(e)
			}
		}
	}(c)

}
