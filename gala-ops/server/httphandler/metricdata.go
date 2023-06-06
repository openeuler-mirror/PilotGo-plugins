package httphandler

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"gitee.com/openeuler/PilotGo-plugins/sdk/utils"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"openeuler.org/PilotGo/gala-ops-plugin/client"
)

type Plugin struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Email       string `json:"email"`
	Url         string `json:"url"`
	Enabled     int    `json:"enabled"`
	Status      string `json:"status"`
}

func PrometheusAPI(URL string) (v1.API, error) {
	// 创建一个HTTP客户端
	client, err := api.NewClient(api.Config{
		Address: "http://" + URL,
	})

	if err != nil {
		return nil, err
	}

	// 创建一个API客户端
	v1api := v1.NewAPI(client)
	return v1api, nil
}

func PrometheusMetrics(ctx *gin.Context) {
	bs, err := utils.Request("GET", client.Client().Server+"plugins")
	if err != nil {
		logger.Error("faild to get plugin list: ", err)
	}
	plugins := &[]*Plugin{}
	err = json.Unmarshal(bs, plugins)
	if err != nil {
		logger.Error("unmarshal request plugin info error:%s", err.Error())
	}
	var Prometheus_addr string
	for _, p := range *plugins {
		if p.Name == "gala-ops" {
			Prometheus_addr = p.Url
		}
	}

	promAPI, err := PrometheusAPI(strings.Split(Prometheus_addr, "/")[2])
	if err != nil {
		logger.Error("failed to create prometheus api: ", err)
	}

	// 查询所有metrics列表
	result, warnings, err := promAPI.Query(ctx, "up", time.Now())

	if err != nil {
		logger.Error("failed to query prometheus: ", err)
		return
	}

	if len(warnings) > 0 {
		logger.Warn("Warnings:", warnings)
		return
	}

	fmt.Println("Metrics:")
	fmt.Println(result)
	// for _, metric := range result {
	// 	fmt.Println(metric)
	// }

}
