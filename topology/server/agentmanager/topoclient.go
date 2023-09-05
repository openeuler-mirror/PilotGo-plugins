package agentmanager

import (
	"encoding/json"
	"os"
	"sync"

	"gitee.com/openeuler/PilotGo-plugin-topology-server/conf"
	"gitee.com/openeuler/PilotGo-plugin-topology-server/handler"
	"gitee.com/openeuler/PilotGo-plugin-topology-server/utils"
	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"gitee.com/openeuler/PilotGo-plugins/sdk/plugin/client"
	"gitee.com/openeuler/PilotGo-plugins/sdk/utils/httputils"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
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
		filepath, line, funcname := utils.CallerInfo()
		logger.Error("\n\tfile: %s\n\tline: %d\n\tfunc: %s\n\terr: %s\n", filepath, line, funcname, err.Error())
		return
	}

	statuscode := resp.StatusCode
	if statuscode != 200 {
		filepath, line, funcname := utils.CallerInfo()
		logger.Error("\n\tfile: %s\n\tline: %d\n\tfunc: %s\n\terr: %s\n", filepath, line, funcname, string(resp.Body))
		return
	}

	result := &struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
	}{}

	err = json.Unmarshal(resp.Body, result)
	if err != nil {
		filepath, line, funcname := utils.CallerInfo()
		logger.Error("\n\tfile: %s\n\tline: %d\n\tfunc: %s\n\terr: %s\n", filepath, line, funcname, err.Error())
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
		filepath, line, funcname := utils.CallerInfo()
		logger.Error("\n\tfile: %s\n\tline: %d\n\tfunc: %s\n\terr: %s\n", filepath, line, funcname, err.Error())
		os.Exit(1)
	}
}

func (t *Topoclient) InitWebServer() {
	engine := gin.Default()
	Topo.Sdkmethod.RegisterHandlers(engine)
	handler.InitRouter(engine)
	err := engine.Run(conf.Config().Topo.Server_addr)
	if err != nil {
		filepath, line, funcname := utils.CallerInfo()
		logger.Fatal("\n\tfile: %s\n\tline: %d\n\tfunc: %s\n\terr: %s\n", filepath, line, funcname, err.Error())
	}
}

func (t *Topoclient) InitPluginClient() {
	PluginClient := client.DefaultClient(PluginInfo)
	PluginClient.Server = "http://" + conf.Config().PilotGo.Addr
	Topo = &Topoclient{
		Sdkmethod: PluginClient,
	}
}
