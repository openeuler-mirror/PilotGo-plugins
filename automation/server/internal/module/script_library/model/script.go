package model

type Script struct {
	ID                  string `json:"id" gorm:"primaryKey;type:varchar(36);not null;comment:'脚本ID'"`
	Name                string `json:"name" gorm:"type:varchar(255);not null;unique;comment:'脚本名称'"`
	ScriptType          string `json:"script_type" gorm:"type:varchar(100);not null;comment:'脚本类型'"`
	Description         string `json:"description" gorm:"type:varchar(500);comment:'脚本描述'"`
	Tags                string `json:"tags" gorm:"comment:'场景标签'"`
	IsPublic            bool   `json:"is_public" gorm:"type:boolean;not null;comment:'是否公开'"`
	Creator             string `json:"creator" gorm:"type:varchar(100);not null;comment:'创建者'"`
	CreatedAt           string `json:"created_at" gorm:"type:datetime;not null;comment:'创建时间'"`
	LastModifyUser      string `json:"last_modify_user" gorm:"type:varchar(100);not null;comment:'最后修改者'"`
	LastModifyUpdatedAt string `json:"last_modify_updated_at" gorm:"type:datetime;not null;comment:'最后修改时间'"`
}

type ScriptVersion struct {
	ID                  int    `json:"id" gorm:"primaryKey;type:int;not null;comment:'脚本版本ID'"`
	ScriptID            string `json:"script_id" gorm:"type:varchar(36);not null;comment:'脚本ID'"`
	Content             string `json:"content" gorm:"type:text;not null;comment:'脚本内容'"`
	Version             string `json:"version" gorm:"type:varchar(50);not null;comment:'脚本版本号'"`
	VersionDesc         string `json:"version_desc" gorm:"type:varchar(500);comment:'脚本版本描述'"`
	Status              bool   `json:"status" gorm:"type:boolean;not null;default:false;comment:'脚本版本状态,true表示上线,false表示禁用'"`
	Creator             string `json:"creator" gorm:"type:varchar(100);not null;comment:'创建者'"`
	CreatedAt           string `json:"created_at" gorm:"type:datetime;not null;comment:'创建时间'"`
	LastModifyUser      string `json:"last_modify_user" gorm:"type:varchar(100);not null;comment:'最后修改者'"`
	LastModifyUpdatedAt string `json:"last_modify_updated_at" gorm:"type:datetime;not null;comment:'最后修改时间'"`
}

type ScriptWithVersion struct {
	Name           string `json:"name"`
	ScriptType     string `json:"script_type"`
	Description    string `json:"description"`
	Tags           string `json:"tags"`
	Content        string `json:"content"`
	Version        string `json:"version"`
	VersionDesc    string `json:"version_desc"`
	IsPublic       bool   `json:"is_public"`
	LastModifyUser string `json:"last_modify_user"`
	Creator        string `json:"creator"`
}
