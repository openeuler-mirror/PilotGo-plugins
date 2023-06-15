package httphandler

import (
	"net/http"

	"gitee.com/openeuler/PilotGo-plugins/sdk/response"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/prometheus-plugin/httphandler/service"
)

func DBTargets(c *gin.Context) {
	targets, err := service.GetPrometheusTarget()
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	objs := []PrometheusObject{
		{
			Targets: targets,
		},
	}
	c.JSON(http.StatusOK, objs)
}

type PrometheusObject struct {
	Targets []string `json:"targets"`
}
