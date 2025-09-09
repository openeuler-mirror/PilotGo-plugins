package service

import (
	"time"

	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/dangerous_rule/dao"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/dangerous_rule/model"
)

func AddDangerousRule(rule *model.DangerousRule) error {
	if err := dao.AddDangerousRule(&model.DangerousRule{
		Expression:  rule.Expression,
		Description: rule.Description,
		ScriptType:  rule.ScriptType,
		Severity:    rule.Severity,
		Creator:     rule.Creator,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		Status:      rule.Status,
	}); err != nil {
		return err
	}
	return nil
}

func GetDangerousRules() ([]model.DangerousRule, error) {
	return dao.GetDangerousRules()
}

func UpdateDangerousRule(rule *model.DangerousRule) error {
	rule.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	return dao.UpdateDangerousRule(rule)
}

func DeleteDangerousRule(id []int) error {
	return dao.DeleteDangerousRule(id)
}
