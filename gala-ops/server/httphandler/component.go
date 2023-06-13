package httphandler

import (
	"time"

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

