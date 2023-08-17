package handler

import (
	"fmt"
	"net/http"

	"gitee.com/openeuler/PilotGo-plugin-topology-agent/collector"
	"gitee.com/openeuler/PilotGo-plugin-topology-agent/service"
	"gitee.com/openeuler/PilotGo-plugin-topology-agent/utils"
	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"github.com/gin-gonic/gin"
)

func Raw_metric_data(ctx *gin.Context) {
	data, err := service.DataCollectorService()
	if err != nil {
		logger.Error("(Raw_metric_data->DataCollectorService: %s)", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   -1,
			"status": fmt.Errorf("(Raw_metric_data->DataCollectorService: %s)", err),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":   0,
		"status": "ok",
		"size":   utils.GetSize(*data.(*collector.PsutilCollector)),
		"data":   data,
	})
}
