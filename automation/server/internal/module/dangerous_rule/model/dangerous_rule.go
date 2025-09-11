package model

import (
	"encoding/json"

	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum/rule"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum/script"
)

type DangerousRule struct {
	ID          int                  `json:"id" gorm:"primaryKey;autoIncrement;comment:规则ID"`
	Expression  string               `json:"expression" gorm:"type:varchar(255);uniqueIndex:uniq_expression;comment:语法检测表达式"`
	Description string               `json:"description" gorm:"type:varchar(255);comment:规则描述"`
	ScriptTypes script.ScriptTypeArr `json:"script_types" gorm:"type:json;comment:脚本类型"`
	Action      rule.ActionType      `json:"action" gorm:"comment:执行动作： 拦截（脚本不可保存、带参数时是否可执行）, 警告（用户二次确认）"`
	Creator     string               `json:"creator" gorm:"comment:创建人"`
	CreatedAt   string               `json:"created_at" gorm:"comment:创建时间"`
	UpdatedAt   string               `json:"updated_at" gorm:"comment:更新时间"`
	Status      bool                 `json:"status" gorm:"comment:规则启用、禁用"`
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

func (r *DangerousRule) UnmarshalJSON(data []byte) error {
	type Alias DangerousRule
	aux := &struct {
		Action      string   `json:"action"`
		ScriptTypes []string `json:"script_types"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	r.Action = rule.ParseActionType(aux.Action)
	r.ScriptTypes = script.NewScriptTypeArr(aux.ScriptTypes)
	return nil
}
