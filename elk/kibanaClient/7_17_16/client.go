package kibanaClient

import (
	"context"
	"errors"

	"gitee.com/openeuler/PilotGo-plugin-elk/conf"
	"gitee.com/openeuler/PilotGo-plugin-elk/errormanager"
	"gitee.com/openeuler/PilotGo-plugin-elk/pluginclient"
	"github.com/elastic/elastic-agent-libs/kibana"
)

var Global_kibana *KibanaClient_v7

type KibanaClient_v7 struct {
	Client *kibana.Client
	Ctx    context.Context
}

func InitKibanaClient() {
	cfg := &kibana.ClientConfig{
		Protocol: "http",
		Host:     conf.Global_Config.Kibana.Addr,
		Username: conf.Global_Config.Kibana.Username,
		Password: conf.Global_Config.Kibana.Password,
	}

	ki_client, err := kibana.NewClientWithConfig(cfg, "", "", "", "")
	if err != nil {
		err = errors.New("failed to init kibana client **errstackfatal**0") // err top
		errormanager.ErrorTransmit(pluginclient.Global_Context, err, true)
		return
	}

	Global_kibana = &KibanaClient_v7{
		Client: ki_client,
		Ctx:    context.Background(),
	}
}
