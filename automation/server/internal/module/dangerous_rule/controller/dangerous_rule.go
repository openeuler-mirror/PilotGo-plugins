package controller

import (
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/dangerous_rule/model"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/dangerous_rule/service"
)

func AddDangerousRuleHandler(c *gin.Context) {
	var rule model.DangerousRule
	if err := c.ShouldBindJSON(&rule); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := service.AddDangerousRule(&rule); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

func GetDangerousRulesHandler(c *gin.Context) {
	rules, err := service.GetDangerousRules()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": rules})
}

func UpdateDangerousRuleHandler(c *gin.Context) {
	var rule model.DangerousRule
	if err := c.ShouldBindJSON(&rule); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := service.UpdateDangerousRule(&rule); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

func DeleteDangerousRuleHandler(c *gin.Context) {
	var ids struct {
		ID []int `json:"id"`
	}
	if err := c.ShouldBindJSON(&ids); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := service.DeleteDangerousRule(ids.ID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}
