package controller

import (
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/dangerous_rule/model"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/dangerous_rule/service"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/pkg/response"
)

func AddDangerousRuleHandler(c *gin.Context) {
	var rule model.DangerousRule
	if err := c.ShouldBindJSON(&rule); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	if err := service.AddDangerousRule(&rule); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "success")
}

func GetDangerousRulesHandler(c *gin.Context) {
	query := &response.PagedQuery{}
	if err := c.ShouldBindQuery(query); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	rules, total, err := service.GetDangerousRules(query)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	response.DataPaged(c, rules, total, query)

}

func UpdateDangerousRuleHandler(c *gin.Context) {
	var rule model.DangerousRule
	if err := c.ShouldBindJSON(&rule); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	if err := service.UpdateDangerousRule(&rule); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "success")
}
func ChangeDangerousRuleStatusHandler(c *gin.Context) {
	var rule struct {
		ID     int    `json:"id"`
		User   string `json:"modify_user"`
		Status bool   `json:"status"`
	}
	if err := c.ShouldBindJSON(&rule); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	if err := service.ChangeDangerousRuleStatus(rule.ID, rule.User, rule.Status); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "success")
}

func DeleteDangerousRuleHandler(c *gin.Context) {
	var ids struct {
		ID []int `json:"id"`
	}
	if err := c.ShouldBindJSON(&ids); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	if err := service.DeleteDangerousRule(ids.ID); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "success")
}

func DetectRealtimelyHandler(c *gin.Context) {
	var req struct {
		Script     string `json:"script"`
		ScriptType string `json:"script_type"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	rules, err := service.DetectRealtimely(req.Script, req.ScriptType)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, rules, "success")
}
