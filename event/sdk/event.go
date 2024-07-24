package sdk

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"gitee.com/openeuler/PilotGo/sdk/common"
	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/utils/httputils"
	"github.com/gin-gonic/gin"
)

func RegisterEventHandlers(router *gin.Engine) {

	api := router.Group("/plugin_manage/api/v1/")
	{
		api.POST("/event", eventHandler)
	}

	startEventProcessor()
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

// 发布event事件
func PublishEvent(msg common.EventMessage) error {
	eventServer, err := eventPluginServer()
	if err != nil {
		return err
	}
	url := eventServer + "/api/v1/pluginapi/publish_event"
	r, err := httputils.Put(url, &httputils.Params{
		Body: &msg,
	})
	if err != nil {
		return err
	}
	if r.StatusCode != http.StatusOK {
		return errors.New("server process error:" + strconv.Itoa(r.StatusCode))
	}

	resp := &common.CommonResult{}
	if err := json.Unmarshal(r.Body, resp); err != nil {
		return err
	}
	if resp.Code != http.StatusOK {
		return errors.New(resp.Message)
	}

	data := &struct {
		Status string `json:"status"`
		Error  string `json:"error"`
	}{}
	if err := resp.ParseData(data); err != nil {
		return err
	}
	return nil
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

func registerEventCallback(eventType int, callback EventCallback) {
	plugin_client.EventCallbackMap[eventType] = callback
}

func unregisterEventCallback(eventType int) {
	delete(plugin_client.EventCallbackMap, eventType)
}

func ProcessEvent(event *common.EventMessage) {
	plugin_client.EventChan <- event
}

func startEventProcessor() {
	go func() {
		for {
			e := <-plugin_client.EventChan

			// TODO: process event message
			cb, ok := plugin_client.EventCallbackMap[e.MessageType]
			if ok {
				cb(e)
			}
		}
	}()

}
