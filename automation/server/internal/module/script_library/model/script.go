package model

import "openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum/script"

type Script struct {
	ID          string            `json:"id" gorm:"primaryKey;type:varchar(36);not null;comment:脚本ID"`
	Name        string            `json:"name" gorm:"type:varchar(255);not null;uniqueIndex:uniq_script_name;comment:脚本名称"`
	ScriptType  script.ScriptType `json:"script_type" gorm:"type:varchar(100);not null;comment:脚本类型"`
	Description string            `json:"description" gorm:"type:varchar(500);comment:脚本描述"`
	Tag         string            `json:"tag" gorm:"comment:场景标签"`
	UsageType   string            `json:"usage_type" gorm:"type:varchar(50);not null;uniqueIndex:uniq_script_name;comment:脚本业务类型"`
	ModifyUser  string            `json:"modify_user" gorm:"type:varchar(100);not null;comment:最后修改者"`
	ModifyTime  string            `json:"modify_time" gorm:"comment:最后修改时间"`
}

type ScriptResponse struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	ScriptType  script.ScriptType `json:"script_type"`
	UsageType   string            `json:"usage_type"`
	Description string            `json:"description"`
	Tag         Tag               `json:"tag"`
	ModifyUser  string            `json:"modify_user"`
	ModifyTime  string            `json:"modify_time"`
}

type AddScriptDTO struct {
	Script       Script        `json:"script"`
	FirstVersion ScriptVersion `json:"first_version"`
}
