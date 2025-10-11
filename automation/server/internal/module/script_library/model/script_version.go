package model

import (
	"encoding/json"

	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum/script"
)

type ScriptParam struct { //脚本执行参数
	Name        string      `json:"name"`
	Type        string      `json:"type"` // string/int/bool 等
	Required    bool        `json:"required"`
	Default     interface{} `json:"default"`
	Description string      `json:"description"`
}

type ScriptVersion struct {
	ID          int                        `json:"id" gorm:"primaryKey;type:int;not null;comment:脚本版本ID"`
	ScriptID    string                     `json:"script_id" gorm:"type:varchar(36);not null;uniqueIndex:uniq_script_version;comment:脚本ID"`
	Params      json.RawMessage            `json:"params" gorm:"type:json;comment:脚本执行参数"` // 存 ScriptParam 数组
	Content     string                     `json:"content" gorm:"type:text;not null;comment:脚本内容"`
	Version     string                     `json:"version" gorm:"type:varchar(50);not null;uniqueIndex:uniq_script_version;comment:脚本版本号"`
	VersionDesc string                     `json:"version_desc" gorm:"type:varchar(500);comment:脚本版本描述"`
	Status      script.ScriptPublishStatus `json:"status" gorm:"default:1;comment:脚本版本状态"`
	ModifyUser  string                     `json:"modify_user" gorm:"type:varchar(100);not null;comment:最后修改者"`
	ModifyTime  string                     `json:"modify_time" gorm:"comment:最后修改时间"`
}

type ScriptVersionResponse struct {
	ID             string          `json:"id"`
	Name           string          `json:"name"`
	ScriptType     string          `json:"script_type"`
	Description    string          `json:"description"`
	Tag            Tag             `json:"tag"`
	ScriptVersions []ScriptVersion `json:"script_versions"`
}

type RawScriptVersion struct {
	// Script 字段
	ScriptID    string `json:"script_id"`
	Name        string `json:"name"`
	ScriptType  string `json:"script_type"`
	Description string `json:"description"`

	// JSON Script versions
	Versions string `json:"versions"`

	// JSON tag
	Tag string `json:"tag"`
}
