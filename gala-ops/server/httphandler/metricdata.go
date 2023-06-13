package httphandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"gitee.com/openeuler/PilotGo-plugins/sdk/utils"
	"github.com/gin-gonic/gin"
)

func TargetsList(ctx *gin.Context) {
	// 查询prometheus监控对象列表
	promurl := Galaops.PromePlugin["url"].(string)
	promsq := "/api/v1/query?query=up"
	bs, err := utils.Request("GET", promurl+promsq)
	if err != nil {
		logger.Error("faild to get prometheus targets: ", err)
	}

	var data interface{}

	err = json.Unmarshal(bs, &data)
	if err != nil {
		logger.Error("unmarshal prometheus targets error:%s", err.Error())
	}
	ctx.JSON(http.StatusOK, data)
}

func CPUusagerate(ctx *gin.Context) {
	promurl := Galaops.PromePlugin["url"].(string)
	start, end := Galaops.UnixTimeStartandEnd(-5)
	param := map[string]string{
		//"query": `avg%20by(job)%20(gala_gopher_cpu_total_used_per%7Bjob%3D~%22192.168.75.132%22%7D)`,
		"query": `avg by(job) (gala_gopher_cpu_total_used_per{job=~"192.168.75.132"})`,
		"start": fmt.Sprint(start),
		"end":   fmt.Sprint(end),
		"step":  "15s",
	}

	urlparam := fmt.Sprintf("?query=%v&start=%v&end=%v&step=%v", param["query"], param["start"], param["end"], param["step"])
	logger.Debug(urlparam)
	data, err := Galaops.QueryMetric(promurl, "query_range", urlparam)
	if err != nil {
		logger.Error("faild to querymetric from prometheus: ", err)
	}
	ctx.JSON(http.StatusOK, data)
}
