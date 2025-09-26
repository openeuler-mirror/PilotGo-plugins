package dao

import (
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/global"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/dangerous_rule/model"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/pkg/response"
)

func AddDangerousRule(rule *model.DangerousRule) error {
	return global.App.MySQL.Save(rule).Error
}

func GetDangerousRules(query *response.PagedQuery) ([]model.DangerousRule, int, error) {
	var rules []model.DangerousRule
	var total int64

	q := global.App.MySQL.Model(&model.DangerousRule{}).Limit(query.PageSize).Offset((query.CurrentPage - 1) * query.PageSize)
	qc := global.App.MySQL.Model(&model.DangerousRule{})

	if err := q.Order("id desc").Find(&rules).Error; err != nil {
		return rules, 0, err
	}

	if err := qc.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return rules, int(total), nil
}

func UpdateDangerousRule(rule *model.DangerousRule) error {
	return global.App.MySQL.Model(&model.DangerousRule{}).Where("id = ?", rule.ID).Updates(rule).Error
}

func ChangeDangerousRuleStatus(id int, data interface{}) error {
	return global.App.MySQL.Model(&model.DangerousRule{}).Where("id = ?", id).Updates(data).Error
}

func DeleteDangerousRule(id []int) error {
	return global.App.MySQL.Where("id IN ?", id).Delete(&model.DangerousRule{}).Unscoped().Error
}
