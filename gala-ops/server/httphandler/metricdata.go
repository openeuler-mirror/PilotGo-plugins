package httphandler

import (
	"encoding/json"
	"net/http"

	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"gitee.com/openeuler/PilotGo-plugins/sdk/utils"
	"github.com/gin-gonic/gin"
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

func TargetsList(ctx *gin.Context) {
	// 查询prometheus监控对象列表
	bs, err := utils.Request("GET", "http://192.168.75.100:8090/plugin/Prometheu/api/v1/query?query=up")
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
