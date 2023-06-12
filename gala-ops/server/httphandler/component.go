package httphandler

import "gitee.com/openeuler/PilotGo-plugins/sdk/plugin/client"

type Opsclient struct {
	Sdkmethod   *client.Client
	PromePlugin map[string]interface{}
}

var Galaops *Opsclient
