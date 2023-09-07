package collector

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"gitee.com/openeuler/PilotGo-plugin-topology-server/agentmanager"
	"gitee.com/openeuler/PilotGo-plugin-topology-server/conf"
	"gitee.com/openeuler/PilotGo-plugin-topology-server/meta"
	"gitee.com/openeuler/PilotGo-plugin-topology-server/utils"
	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"gitee.com/openeuler/PilotGo-plugins/sdk/utils/httputils"
	"github.com/mitchellh/mapstructure"
)

type DataCollector struct{}

func CreateDataCollector() *DataCollector {
	return &DataCollector{}
}

func (d *DataCollector) Collect_instant_data() error {
	start := time.Now()
	var wg sync.WaitGroup

	agentmanager.Topo.AgentMap.Range(
		func(key, value any) bool {
			wg.Add(1)
			go func() {
				defer wg.Done()
				agent := value.(*agentmanager.Agent_m)
				agent.Port = conf.Config().Topo.Agent_port
				err := d.GetCollectDataFromTopoAgent(agent)
				if err != nil {
					filepath, line, funcname := utils.CallerInfo()
					logger.Error("\n\tfile: %s\n\tline: %d\n\tfunc: %s\n", filepath, line, funcname)
				}
				agentmanager.Topo.AddAgent(agent)
			}()
			return true
		},
	)

	wg.Wait()

	elapse := time.Since(start)
	logger.Debug("\033[32mtopo server 采集数据获取时间\033[0m: %v\n", elapse)

	return nil
}

func (d *DataCollector) GetCollectDataFromTopoAgent(agent *agentmanager.Agent_m) error {
	url := "http://" + agent.IP + ":" + agent.Port + "/plugin/api/rawdata"

	resp, err := httputils.Get(url, nil)
	if err != nil {
		filepath, line, funcname := utils.CallerInfo()
		logger.Error("\n\tfile: %s\n\tline: %d\n\tfunc: %s\n\terr: %s\n", filepath, line, funcname, err.Error())
		return fmt.Errorf("file: %s, line: %d, func: %s, err -> %s", filepath, line, funcname, err.Error())
	}

	results := &struct {
		Code  int         `json:"code"`
		Error string      `json:"error"`
		Data  interface{} `json:"data"`
	}{}

	err = json.Unmarshal(resp.Body, &results)
	if err != nil {
		filepath, line, funcname := utils.CallerInfo()
		logger.Error("\n\tfile: %s\n\tline: %d\n\tfunc: %s\n\terr: %s\n", filepath, line, funcname, err.Error())
		return fmt.Errorf("file: %s, line: %d, func: %s, err -> %s", filepath, line, funcname, err.Error())
	}

	statuscode := results.Code
	if statuscode != 0 {
		filepath, line, funcname := utils.CallerInfo()
		logger.Error("\n\tfile: %s\n\tline: %d\n\tfunc: %s\n\terr: %s\n", filepath, line, funcname, err.Error())
		return fmt.Errorf(results.Error)
	}

	collectdata := &struct {
		Host_1           *meta.Host
		Processes_1      []*meta.Process
		Netconnections_1 []*meta.Netconnection
	}{}
	mapstructure.Decode(results.Data, collectdata)

	agent.Host_2 = collectdata.Host_1
	agent.Processes_2 = collectdata.Processes_1
	agent.Netconnections_2 = collectdata.Netconnections_1

	return nil
}