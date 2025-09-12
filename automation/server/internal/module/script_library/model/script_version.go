package model

type ScriptVersion struct {
	ID                  int    `json:"id" gorm:"primaryKey;type:int;not null;comment:'脚本版本ID'"`
	ScriptID            string `json:"script_id" gorm:"type:varchar(36);not null;uniqueIndex:uniq_script_version;comment:'脚本ID'"`
	Content             string `json:"content" gorm:"type:text;not null;comment:'脚本内容'"`
	Version             string `json:"version" gorm:"type:varchar(50);not null;uniqueIndex:uniq_script_version;comment:'脚本版本号'"`
	VersionDesc         string `json:"version_desc" gorm:"type:varchar(500);uniqueIndex:uniq_script_version;comment:'脚本版本描述'"`
	Status              bool   `json:"status" gorm:"type:boolean;not null;default:false;comment:'脚本版本状态,true表示上线,false表示禁用'"`
	Creator             string `json:"creator" gorm:"type:varchar(100);not null;comment:'创建者'"`
	CreatedAt           string `json:"created_at" gorm:"comment:'创建时间'"`
	LastModifyUser      string `json:"last_modify_user" gorm:"type:varchar(100);not null;comment:'最后修改者'"`
	LastModifyUpdatedAt string `json:"last_modify_updated_at" gorm:"comment:'最后修改时间'"`
}

type ScriptVersionResponse struct {
	ID             string          `json:"id"`
	Name           string          `json:"name"`
	ScriptType     string          `json:"script_type"`
	Description    string          `json:"description"`
	Tags           []Tag           `json:"tags"`
	IsPublic       bool            `json:"is_public"`
	ScriptVersions []ScriptVersion `json:"script_versions"`
}

type RawScriptVersion struct {
	// Script 字段
	ScriptID    string `json:"script_id"`
	Name        string `json:"name"`
	ScriptType  string `json:"script_type"`
	Description string `json:"description"`
	IsPublic    bool   `json:"is_public"`

	// JSON Script versions
	Versions string `json:"versions"`

	// JSON tags
	Tags string `json:"tags"`
}
