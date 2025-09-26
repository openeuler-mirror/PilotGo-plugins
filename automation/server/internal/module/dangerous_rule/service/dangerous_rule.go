package service

import (
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/global"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/dangerous_rule/dao"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/dangerous_rule/model"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/pkg/response"
)

func AddDangerousRule(rule *model.DangerousRule) error {
	if err := dao.AddDangerousRule(&model.DangerousRule{
		Expression:  rule.Expression,
		Description: rule.Description,
		ScriptTypes: rule.ScriptTypes,
		Action:      rule.Action,
		Creator:     rule.Creator,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		Status:      rule.Status,
	}); err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 { // ER_DUP_ENTRY
				return fmt.Errorf("规则已存在: %s", rule.Expression)
			}
		}
		return err
	}
	return LoadFromDB()
}

func GetDangerousRules(query *response.PagedQuery) ([]model.DangerousRule, int, error) {
	return dao.GetDangerousRules(query)
}

func UpdateDangerousRule(rule *model.DangerousRule) error {
	if err := dao.UpdateDangerousRule(&model.DangerousRule{
		ID:          rule.ID,
		Expression:  rule.Expression,
		Description: rule.Description,
		ScriptTypes: rule.ScriptTypes,
		Action:      rule.Action,
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		Status:      rule.Status,
	}); err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 { // ER_DUP_ENTRY
				return fmt.Errorf("规则已存在: %s", rule.Expression)
			}
		}
		return err
	}
	return LoadFromDB()
}

func ChangeDangerousRuleStatus(id int, status bool) error {
	if id == 0 {
		return fmt.Errorf("ID is required")
	}
	err := dao.ChangeDangerousRuleStatus(id, map[string]interface{}{
		"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		"status":     status,
	})
	if err != nil {
		return err
	}
	return LoadFromDB()
}

func DeleteDangerousRule(id []int) error {
	if err := dao.DeleteDangerousRule(id); err != nil {
		return err
	}
	return LoadFromDB()
}

/******************************更新redis**************************************/
const DangerousRuleKey = "dangerous_rules"

func LoadFromDB() error {
	var rules []model.DangerousRule
	if err := global.App.MySQL.Where("status = ?", 1).Find(&rules).Error; err != nil {
		return err
	}

	return global.App.Redis.Set(DangerousRuleKey, rules, 0)
}
