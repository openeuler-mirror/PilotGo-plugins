package client

import (
	"encoding/json"

	"gitee.com/openeuler/PilotGo-plugins/sdk/common"
	"gitee.com/openeuler/PilotGo-plugins/sdk/utils/httputils"
)

func (c *Client) ServiceStatus(batch *common.Batch, servicename string) ([]*CmdResult, error) {
	url := c.Server + "/api/v1/pluginapi/service/:name"
	p := &struct {
		Batch       *common.Batch `json:batch`
		ServiceName string        `json:service`
	}{
		Batch:       batch,
		ServiceName: servicename,
	}
	r, err := httputils.Put(url, &httputils.Params{
		Body: p,
	})
	if err != nil {
		return nil, err
	}

	res := &struct {
		Code    int          `json:"code"`
		Mseeage string       `json:"msg"`
		Data    []*CmdResult `json:"data`
	}{}
	if err := json.Unmarshal(r.Body, res); err != nil {
		return nil, err
	}
	return res.Data, nil
}
