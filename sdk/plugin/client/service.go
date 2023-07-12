package client

import (
	"encoding/json"

	"gitee.com/openeuler/PilotGo-plugins/sdk/common"
	"gitee.com/openeuler/PilotGo-plugins/sdk/utils/httputils"
)

type ServiceResult struct {
	MachineUUID   string
	MachineIP     string
	RetCode       int
	ServiceStatus string
}

func (c *Client) ServiceStatus(batch *common.Batch, servicename string) ([]*ServiceResult, error) {
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
		Code    int              `json:"code"`
		Mseeage string           `json:"msg"`
		Data    []*ServiceResult `json:"data`
	}{}
	if err := json.Unmarshal(r.Body, res); err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (c *Client) StartService(batch *common.Batch, serviceName string) ([]*ServiceResult, error) {
	url := c.Server + "/api/v1/pluginapi/start_service"

	p := &struct {
		Batch       *common.Batch `json:batch`
		ServiceName string        `json:service`
	}{
		Batch:       batch,
		ServiceName: serviceName,
	}

	r, err := httputils.Put(url, &httputils.Params{
		Body: p,
	})
	if err != nil {
		return nil, err
	}

	res := &struct {
		Code    int              `json:"code"`
		Message string           `json:"msg"`
		Data    []*ServiceResult `json:"data"`
	}{}
	if err := json.Unmarshal(r.Body, res); err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (c *Client) StopService(batch *common.Batch, serviceName string) ([]*ServiceResult, error) {
	url := c.Server + "/api/v1/pluginapi/stop_service"

	p := &struct {
		Batch       *common.Batch `json:batch`
		ServiceName string        `json:service`
	}{
		Batch:       batch,
		ServiceName: serviceName,
	}

	r, err := httputils.Put(url, &httputils.Params{
		Body: p,
	})
	if err != nil {
		return nil, err
	}

	res := &struct {
		Code    int              `json:"code"`
		Message string           `json:"msg"`
		Data    []*ServiceResult `json:"data"`
	}{}
	if err := json.Unmarshal(r.Body, res); err != nil {
		return nil, err
	}

	return res.Data, nil
}
