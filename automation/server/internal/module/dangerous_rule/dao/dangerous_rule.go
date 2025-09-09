package dao

import (
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/global"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/dangerous_rule/model"
)

func AddDangerousRule(rule *model.DangerousRule) error {
	return global.App.MySQL.Save(rule).Error
}
