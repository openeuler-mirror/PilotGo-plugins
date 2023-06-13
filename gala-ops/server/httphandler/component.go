package httphandler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"gitee.com/openeuler/PilotGo-plugins/sdk/plugin/client"
)

type Opsclient struct {
	Sdkmethod   *client.Client
	PromePlugin map[string]interface{}
}

var Galaops *Opsclient

func (o *Opsclient) UnixTimeStartandEnd(timerange time.Duration) (int64, int64) {
	now := time.Now()
	past5Minutes := now.Add(timerange * time.Minute)
	startOfPast5Minutes := time.Date(past5Minutes.Year(), past5Minutes.Month(), past5Minutes.Day(),
		past5Minutes.Hour(), past5Minutes.Minute(), 0, 0, past5Minutes.Location())
	timestamp := startOfPast5Minutes.Unix()
	return timestamp, now.Unix()
}

func (o *Opsclient) QueryMetric(endpoint string, querymethod string, param string) (interface{}, error) {
	ustr := endpoint + "/api/v1/" + querymethod + param
	u, err := url.Parse(ustr)
	if err != nil {
		return nil, err
	}
	u.RawQuery = u.Query().Encode()

	httpClient := &http.Client{Timeout: 10 * time.Second}
	resp, err := httpClient.Get(u.String())
	if err != nil {
		return nil, err
	}
	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data interface{}

	err = json.Unmarshal(bs, &data)
	if err != nil {
		logger.Error("unmarshal cpu usage rate error:%s", err.Error())
		return nil, fmt.Errorf("unmarshal cpu usage rate error:%s", err.Error())
	}
	return data, nil
}
