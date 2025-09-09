package model

import (
	"encoding/json"

	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum/rule"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum/script"
)

type DangerousRule struct {
	ID          int                  `json:"id" gorm:"primaryKey;autoIncrement"`
	Expression  string               `gorm:"type:varchar(255);uniqueIndex:uniq_expression" json:"expression"` // 语法检测表达式
	Description string               `json:"description"`                                                     // 规则描述
	ScriptTypes script.ScriptTypeArr `gorm:"type:json" json:"script_types"`                                   // 脚本类型
	Action      rule.ActionType      `json:"action"`                                                          // 执行动作： 拦截（脚本不可保存、带参数时是否可执行）, 警告（用户二次确认）
	Creator     string               `json:"creator"`                                                         // 创建人
	CreatedAt   string               `json:"created_at"`                                                      // 创建时间
	UpdatedAt   string               `json:"updated_at"`                                                      // 更新时间
	Status      bool                 `json:"status"`                                                          // 规则启用、禁用
}

func (r DangerousRule) MarshalJSON() ([]byte, error) {
	type Alias DangerousRule
	return json.Marshal(&struct {
		Action      string   `json:"action"`
		ScriptTypes []string `json:"script_types"`
		Alias
	}{
		Action:      r.Action.String(),
		ScriptTypes: r.ScriptTypes.Strings(),
		Alias:       (Alias)(r),
	})
}
