package clientmanager

import (
	"sync"

	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"gitee.com/openeuler/PilotGo-plugins/sdk/plugin/client"
)

const Version = "0.0.1"

var PluginInfo = &client.PluginInfo{
	Name:        "topo",
	Version:     Version,
	Description: "system application architecture perception",
	Author:      "wangjunqi",
	Email:       "wangjunqi@kylinos.cn",
	Url:         "http://192.168.75.100:9995/plugin/topo",
}

type Opsclient struct {
	Sdkmethod *client.Client
	AgentMap  sync.Map
}

var Galaops *Opsclient

type Agent struct {
	ID         uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	UUID       string `gorm:"not null;unique" json:"UUID"`
	IP         string `gorm:"not null" json:"IP"`
	Port       string `gorm:"not null" json:"port"`
	Department string `json:"Department"`
	State      int    `gorm:"not null" json:"State"` // true:running false:not running
}

/*******************************************************agentmanager*******************************************************/

func (o *Opsclient) AddAgent(a *Agent) {
	o.AgentMap.Store(a.UUID, a)
}

func (o *Opsclient) GetAgent(uuid string) *Agent {
	agent, ok := o.AgentMap.Load(uuid)
	if ok {
		return agent.(*Agent)
	}
	return nil
}

func (o *Opsclient) DeleteAgent(uuid string) {
	if _, ok := o.AgentMap.LoadAndDelete(uuid); !ok {
		logger.Warn("delete known agent:%s", uuid)
	}
}
