package sdk

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"gitee.com/openeuler/PilotGo/sdk/common"
	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
	"gitee.com/openeuler/PilotGo/sdk/utils/httputils"
)

var plugin_client = client.GetClient()

type EventCallback func(e *common.EventMessage)

// 注册event事件监听
func ListenEvent(eventTypes []int, callbacks []EventCallback) error {
	var eventtypes []string
	for _, i := range eventTypes {
		eventtypes = append(eventtypes, strconv.Itoa(i))
	}

	eventServer, err := eventPluginServer()
	if err != nil {
		return err
	}

	url := eventServer + "/plugin/event/listener/register?eventTypes=" + strings.Join(eventtypes, ",")
	r, err := httputils.Put(url, &httputils.Params{
		Body: plugin_client.PluginInfo,
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
	for i, eventType := range eventTypes {
		registerEventCallback(eventType, callbacks[i])
	}
	return nil
}

// 取消注册event事件监听
func UnListenEvent(eventTypes []int) error {
	var eventtypes []string
	for _, i := range eventTypes {
		eventtypes = append(eventtypes, strconv.Itoa(i))
	}
	eventServer, err := eventPluginServer()
	if err != nil {
		return err
	}

	url := eventServer + "/api/v1/pluginapi/listener?eventTypes=" + strings.Join(eventtypes, ",")
	r, err := httputils.Delete(url, &httputils.Params{
		Body: plugin_client.PluginInfo,
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

	for _, eventType := range eventTypes {
		unregisterEventCallback(eventType)
	}
	return nil
}
