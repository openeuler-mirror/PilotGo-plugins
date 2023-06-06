package client

import (
	"encoding/json"
	"errors"

	"gitee.com/openeuler/PilotGo-plugins/sdk/utils"
)

type Event struct {
	ID       int
	MetaData interface{}
}

type EventCallback func(e *Event)

// 注册event事件监听
func (c *Client) ListenEvent(event Event, callback EventCallback) error {
	url := c.Server + "/api/v1/pluginapi/listener"
	data, err := utils.Request("PUT", url)
	if err != nil {
		return err
	}

	resp := &struct {
		Status string
		Error  string
	}{}
	if err := json.Unmarshal(data, resp); err != nil {
		return err
	}
	if resp.Status != "ok" {
		return errors.New(resp.Error)
	}

	// TODO: register event handler here
	return nil
}

// 取消注册event事件监听
func (c *Client) UnListenEvent(listenerID string) error {
	url := c.Server + "/api/v1/pluginapi/listener"
	data, err := utils.Request("DELETE", url)
	if err != nil {
		return err
	}

	resp := &struct {
		Status string
		Error  string
	}{}
	if err := json.Unmarshal(data, resp); err != nil {
		return err
	}
	if resp.Status != "ok" {
		return errors.New(resp.Error)
	}

	// TODO: unregister event handler here
	return nil
}
