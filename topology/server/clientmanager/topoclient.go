package clientmanager

import (
	"sync"

	"gitee.com/openeuler/PilotGo-plugins/sdk/plugin/client"
)

var Galaops *Topoclient

type Topoclient struct {
	Sdkmethod *client.Client
	AgentMap  sync.Map
}
