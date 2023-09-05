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

func (d *DataCollector) GetCollectDataFromTopoAgent(a *agentmanager.Agent_m) error {
	url := "http://" + a.IP + ":" + a.Port + "/plugin/api/rawdata"

	r, err := httputils.Get(url, nil)
	if err != nil {
		filepath, line, funcname := utils.CallerInfo()
		logger.Error("\n\tfile: %s\n\tline: %d\n\tfunc: %s\n\terr: %s\n", filepath, line, funcname, err.Error())
		return fmt.Errorf("file: %s, line: %d, func: %s, err -> %s", filepath, line, funcname, err.Error())
	}

	results := &struct {
		Code   int         `json:"code"`
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
	}{}

	err = json.Unmarshal(r.Body, &results)
	if err != nil {
		filepath, line, funcname := utils.CallerInfo()
		logger.Error("\n\tfile: %s\n\tline: %d\n\tfunc: %s\n\terr: %s\n", filepath, line, funcname, err.Error())
		return fmt.Errorf("file: %s, line: %d, func: %s, err -> %s", filepath, line, funcname, err.Error())
	}

	type collectdata struct {
		Host_1           *meta.Host
		Processes_1      []*meta.Process
		Netconnections_1 []*meta.Netconnection
	}

	if results.Code == -1 {
		return fmt.Errorf(results.Status)
	}

	var collectdata_entity collectdata
	mapstructure.Decode(results.Data, &collectdata_entity)

	// mapstructure.decode映射之后collectdata.Netconnections_1中的网络连接地址字符串为空:
	// mapstructure.decode(input, output): input中的字段为源结构体字段的json别名，当两者不一致时会映射失败（大小写可不一致）
	// data := results.Data.(map[string]interface{})
	// temp_netconnections := []*meta.Netconnection{}
	// for _, net := range data["Netconnections_1"].([]interface{}) {
	// 	net_struct := &meta.Netconnection{}
	// 	net_map := net.(map[string]interface{})
	// 	for k, v := range net_map {
	// 		switch k {
	// 		case "family":
	// 			net_struct.Family = uint32(v.(float64))
	// 		case "fd":
	// 			net_struct.Fd = uint32(v.(float64))
	// 		case "localaddr":
	// 			net_struct.Laddr = v.(string)
	// 		case "remoteaddr":
	// 			net_struct.Raddr = v.(string)
	// 		case "pid":
	// 			net_struct.Pid = int32(v.(float64))
	// 		case "status":
	// 			net_struct.Status = v.(string)
	// 		case "type":
	// 			net_struct.Type = uint32(v.(float64))
	// 		case "uids":
	// 			for _, vuid := range v.([]interface{}) {
	// 				net_struct.Uids = append(net_struct.Uids, int32(vuid.(float64)))
	// 			}
	// 		}
	// 	}
	// 	temp_netconnections = append(temp_netconnections, net_struct)
	// }

	a.Host_2 = collectdata_entity.Host_1
	a.Processes_2 = collectdata_entity.Processes_1
	a.Netconnections_2 = collectdata_entity.Netconnections_1

	return nil
}
