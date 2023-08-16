package clientmanager

import (
	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
)

type Agent struct {
	ID         uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	UUID       string `gorm:"not null;unique" json:"UUID"`
	IP         string `gorm:"not null" json:"IP"`
	Port       string `gorm:"not null" json:"port"`
	Department string `json:"Department"`
	State      int    `gorm:"not null" json:"State"` // true:running false:not running
}

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
