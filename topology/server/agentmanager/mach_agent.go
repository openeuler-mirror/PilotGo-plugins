package agentmanager

import (
	"gitee.com/openeuler/PilotGo-plugin-topology-server/meta"
	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
)

type Agent_m struct {
	ID         uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	UUID       string `gorm:"not null;unique" json:"uuid"`
	IP         string `gorm:"not null" json:"IP"`
	Port       string `gorm:"not null" json:"port"`
	Departid   string `json:"departid"`
	Departname string `json:"departname"`
	State      int    `gorm:"not null" json:"state"`
	TAState    int    `json:"TAstate"` // topo agent state: true(running) false(not runnings)

	Host_2           *meta.Host            `json:"host"`
	Processes_2      []*meta.Process       `json:"processes"`
	Netconnections_2 []*meta.Netconnection `json:"netconnections"`
}

func (t *Topoclient) AddAgent(a *Agent_m) {
	t.AgentMap.Store(a.UUID, a)
}

func (t *Topoclient) GetAgent(uuid string) *Agent_m {
	agent, ok := t.AgentMap.Load(uuid)
	if ok {
		return agent.(*Agent_m)
	}
	return nil
}

func (t *Topoclient) DeleteAgent(uuid string) {
	if _, ok := t.AgentMap.LoadAndDelete(uuid); !ok {
		logger.Warn("delete known agent:%s", uuid)
	}
}