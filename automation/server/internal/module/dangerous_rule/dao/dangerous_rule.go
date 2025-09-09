package dao

import (
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/global"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/dangerous_rule/model"
)

func AddDangerousRule(rule *model.DangerousRule) error {
	return global.App.MySQL.Save(rule).Error
}

func GetDangerousRules() ([]model.DangerousRule, error) {
	var rules []model.DangerousRule
	err := global.App.MySQL.Find(&rules).Error
	return rules, err
}

func UpdateDangerousRule(rule *model.DangerousRule) error {
	return global.App.MySQL.Model(&model.DangerousRule{}).Where("id = ?", rule.ID).Updates(rule).Error
}

func DeleteDangerousRule(id []int) error {
	return global.App.MySQL.Where("id IN ?", id).Delete(&model.DangerousRule{}).Unscoped().Error
}
