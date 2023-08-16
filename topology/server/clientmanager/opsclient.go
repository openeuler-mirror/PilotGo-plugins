package clientmanager

import (
	"sync"

	"gitee.com/openeuler/PilotGo-plugins/sdk/plugin/client"
)

var Galaops *Opsclient

type Opsclient struct {
	Sdkmethod *client.Client
	AgentMap  sync.Map
}
