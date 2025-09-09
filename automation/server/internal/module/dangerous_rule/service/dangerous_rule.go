package service

import (
	"fmt"
	"time"

	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/go-sql-driver/mysql"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/dangerous_rule/dao"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/dangerous_rule/model"
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
	return nil
}

func GetDangerousRules(query *response.PaginationQ) ([]model.DangerousRule, int, error) {
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
	return nil
}

func ChangeDangerousRuleStatus(id int, status bool) error {
	if id == 0 {
		return fmt.Errorf("ID is required")
	}
	return dao.ChangeDangerousRuleStatus(id, map[string]interface{}{
		"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		"status":     status,
	})
}

func DeleteDangerousRule(id []int) error {
	return dao.DeleteDangerousRule(id)
}
