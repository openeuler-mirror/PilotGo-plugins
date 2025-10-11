package dangerousrule

import (
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/dangerous_rule/controller"
)

func DangerousRuleHandler(router *gin.RouterGroup) {
	api := router.Group("/dangerousRule")
	{
		api.POST("/add", controller.AddDangerousRuleHandler)
		api.GET("/list", controller.GetDangerousRulesHandler)
		api.PUT("/update", controller.UpdateDangerousRuleHandler)
		api.PUT("/changeStatus", controller.ChangeDangerousRuleStatusHandler)
		api.DELETE("/delete", controller.DeleteDangerousRuleHandler)
		api.POST("/detectRule", controller.DetectRealtimelyHandler)
	}

}
