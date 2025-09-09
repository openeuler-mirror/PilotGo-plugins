package model

import (
	"encoding/json"

	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum/rule"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum/script"
)

type DangerousRule struct {
	ID          int               `json:"id" gorm:"primaryKey;autoIncrement"`
	Expression  string            `json:"expression"`  // 语法检测表达式
	Description string            `json:"description"` // 规则描述
	ScriptType  script.ScriptType `json:"script_type"` // 脚本类型
	Severity    rule.Severity     `json:"severity"`    // 风险等级： 拦截（脚本不可保存、带参数时是否可执行）, 警告（用户二次确认）
	Creator     string            `json:"creator"`     // 创建人
	CreatedAt   string            `json:"created_at"`  // 创建时间
	UpdatedAt   string            `json:"updated_at"`  // 更新时间
	Status      bool              `json:"status"`      // 规则启用、禁用
}

func (r DangerousRule) MarshalJSON() ([]byte, error) {
	type Alias DangerousRule
	return json.Marshal(&struct {
		Severity   string `json:"severity"`
		ScriptType string `json:"script_type"`
		Alias
	}{
		Severity:   r.Severity.String(), // 把数字转成文字
		ScriptType: r.ScriptType.String(),
		Alias:      (Alias)(r),
	})
}
