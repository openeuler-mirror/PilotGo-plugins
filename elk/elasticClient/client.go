package elasticClient

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"gitee.com/openeuler/PilotGo-plugin-elk/conf"
	"gitee.com/openeuler/PilotGo-plugin-elk/errormanager"
	"gitee.com/openeuler/PilotGo-plugin-elk/pluginclient"
	elastic "github.com/elastic/go-elasticsearch/v7"
)

var Global_elastic *ElasticClient_v7

type ElasticClient_v7 struct {
	Client *elastic.Client
	Ctx    context.Context
}

func InitElasticClient() {
	cfg := elastic.Config{
		Addresses: []string{
			fmt.Sprintf("http://%s", conf.Global_Config.Elasticsearch.Addr),
		},
		Username: conf.Global_Config.Elasticsearch.Username,
		Password: conf.Global_Config.Elasticsearch.Password,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
		},
	}

	es_client, err := elastic.NewClient(cfg)
	if err != nil {
		err = errors.New("failed to init kibana client **errstackfatal**0") // err top
		errormanager.ErrorTransmit(pluginclient.Global_Context, err, true)
		return
	}

	Global_elastic = &ElasticClient_v7{
		Client: es_client,
		Ctx:    context.Background(),
	}
}
