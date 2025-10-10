package model

import "openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum/rule"

type DangerousRule struct {
	ID          int             `json:"id" gorm:"primaryKey;autoIncrement;comment:规则ID"`
	Expression  string          `json:"expression" gorm:"type:varchar(255);uniqueIndex:uniq_expression;comment:语法检测表达式"`
	Description string          `json:"description" gorm:"type:varchar(255);comment:规则描述"`
	ScriptTypes string          `json:"script_types" gorm:"comment:脚本类型"`
	Action      rule.ActionType `json:"action" gorm:"comment:执行动作: 拦截(脚本不可保存、带参数时是否可执行), 警告(用户二次确认)"`
	ModifyUser  string          `json:"modify_user" gorm:"type:varchar(100);not null;comment:最后修改者"`
	ModifyTime  string          `json:"modify_time" gorm:"comment:最后修改时间"`
	Status      bool            `json:"status" gorm:"comment:规则启用、禁用"`
}
