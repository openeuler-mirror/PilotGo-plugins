package agentmanager

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"gitee.com/openeuler/PilotGo-plugin-topology-server/conf"
	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"gitee.com/openeuler/PilotGo-plugins/sdk/plugin/client"
	"gitee.com/openeuler/PilotGo-plugins/sdk/utils/httputils"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

var Topo *Topoclient

type Topoclient struct {
	Sdkmethod *client.Client
	AgentMap  sync.Map
}

func (t *Topoclient) InitMachineList() {
	url := Topo.Sdkmethod.Server + "/api/v1/pluginapi/machine_list"

	resp, err := httputils.Get(url, nil)
	if err != nil {
		err = errors.Errorf("%s**2", err)
		fmt.Printf("%+v\n", err) // err top
		// errors.EORE(err)
		os.Exit(1)
	}

	statuscode := resp.StatusCode
	if statuscode != 200 {
		err = errors.New("http返回状态码异常**2")
		fmt.Printf("%+v\n", err) // err top
		// errors.EORE(err)
	}

	result := &struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
	}{}

	err = json.Unmarshal(resp.Body, result)
	if err != nil {
		err = errors.Errorf("%s**2", err.Error())
		fmt.Printf("%+v\n", err) // err top
		// errors.EORE(err)
	}

	for _, m := range result.Data.([]interface{}) {
		p := &Agent_m{}
		mapstructure.Decode(m, p)
		p.TAState = 0
		Topo.AddAgent(p)
	}
}

func (t *Topoclient) InitLogger() {
	err := logger.Init(conf.Config().Logopts)
	if err != nil {
		err = errors.Errorf("%s**2", err.Error())
		fmt.Printf("%+v\n", err) // err top
		// errors.EORE(err)
		os.Exit(1)
	}
}

func (t *Topoclient) InitPluginClient() {
	PluginInfo.Url = "http://" + conf.Config().Topo.Server_addr + "/plugin/topology"
	PluginClient := client.DefaultClient(PluginInfo)
	PluginClient.Server = "http://" + conf.Config().PilotGo.Addr
	Topo = &Topoclient{
		Sdkmethod: PluginClient,
	}
}

func (t *Topoclient) InitArangodb() {

}
