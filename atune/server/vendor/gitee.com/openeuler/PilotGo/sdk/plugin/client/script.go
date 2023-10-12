package client

import (
	"encoding/base64"
	"encoding/json"

	"gitee.com/openeuler/PilotGo/sdk/common"
	"gitee.com/openeuler/PilotGo/sdk/utils/httputils"
)

type RunCommandCallback func([]*common.CmdResult)

func (c *Client) RunCommand(batch *common.Batch, cmd string) ([]*common.CmdResult, error) {
	url := c.Server + "/api/v1/pluginapi/run_command"

	p := &common.CmdStruct{
		Batch:   batch,
		Command: base64.StdEncoding.EncodeToString([]byte(cmd)),
	}

	r, err := httputils.Post(url, &httputils.Params{
		Body: p,
	})
	if err != nil {
		return nil, err
	}

	res := &struct {
		Code    int                 `json:"code"`
		Message string              `json:"msg"`
		Data    []*common.CmdResult `json:"data"`
	}{}
	if err := json.Unmarshal(r.Body, res); err != nil {
		return nil, err
	}

	return res.Data, nil
}

type ScriptStruct struct {
	Batch  *common.Batch `json:"batch"`
	Script string        `json:"script"`
	Params []string      `json:"params"`
}

func (c *Client) RunScript(batch *common.Batch, script string, params []string) ([]*common.CmdResult, error) {
	url := c.Server + "/api/v1/pluginapi/run_script"

	p := &ScriptStruct{
		Batch:  batch,
		Script: base64.StdEncoding.EncodeToString([]byte(script)),
		Params: params,
	}

	r, err := httputils.Post(url, &httputils.Params{
		Body: p,
	})
	if err != nil {
		return nil, err
	}

	res := &struct {
		Code    int                 `json:"code"`
		Message string              `json:"msg"`
		Data    []*common.CmdResult `json:"data"`
	}{}
	if err := json.Unmarshal(r.Body, res); err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (c *Client) RunCommandAsync(batch *common.Batch, cmd string, callback RunCommandCallback) error {
	url := c.Server + "/api/v1/pluginapi/run_command_async"

	p := &common.CmdStruct{
		Batch:   batch,
		Command: base64.StdEncoding.EncodeToString([]byte(cmd)),
	}

	r, err := httputils.Post(url, &httputils.Params{
		Body: p,
	})
	if err != nil {
		return err
	}

	res := &struct {
		TaskID string `json:"task_id"`
	}{}
	if err := json.Unmarshal(r.Body, res); err != nil {
		return err
	}

	taskID := res.TaskID
	c.registerCommandResultCallback(taskID, callback)

	return nil
}

func (c *Client) startCommandResultProcessor() {
	// TODO：do exit
	for {
		d := <-c.asyncCmdResultChan

		cb, ok := c.cmdProcessorCallbackMap[d.TaskID]
		if !ok {
			continue
		}

		// 注意：map并发安全
		cb(d.Result)
		delete(c.cmdProcessorCallbackMap, d.TaskID)
	}
}

func (c *Client) registerCommandResultCallback(taskID string, callback RunCommandCallback) {
	c.cmdProcessorCallbackMap[taskID] = callback
}
