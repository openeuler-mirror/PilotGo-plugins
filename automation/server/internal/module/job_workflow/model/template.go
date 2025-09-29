package model

import "encoding/json"

type TaskTemplate struct {
	ID           int    `json:"id" gorm:"primaryKey;autoIncrement;comment:作业编排Id"`
	Name         string `json:"name" gorm:"type:varchar(255);comment:作业编排名称"`
	Description  string `json:"description" gorm:"type:varchar(255);comment:作业描述"`
	Tags         string `json:"tags" gorm:"comment:场景标签"`
	FirstStepNum int    `json:"first_step_num"`
	LastStepNum  int    `json:"last_step_num"`
	ModifyUser   string `json:"modify_user" gorm:"type:varchar(100);not null;comment:'最后修改者'"`
	ModifyTime   string `json:"modify_time" gorm:"comment:'最后修改时间'"`
}

type TaskTemplateVariable struct {
	ID           int             `json:"id" gorm:"primaryKey;autoIncrement"`
	TemplateId   int             `json:"template_id" gorm:"comment:作业编排Id"`
	Name         string          `json:"name" gorm:"type:varchar(255);comment:变量名称"`
	Type         string          `json:"type" gorm:"comment:变量类型(字符串、命名空间、数组等)"`
	DefaultVaue  json.RawMessage `json:"default_value" gorm:"comment:变量默认值"`
	Description  string          `json:"description" gorm:"comment:变量描述"`
	IsChangeable bool            `json:"is_changeable" gorm:"comment:赋值可变"`
	IsRequired   bool            `json:"is_required" gorm:"comment:是否必需"`
}

type TaskTemplateStep struct {
	ID              int    `json:"id" gorm:"primaryKey;autoIncrement"`
	TemplateId      int    `json:"template_id" gorm:"comment:作业编排Id"`
	StepType        string `json:"step_type" gorm:"comment:编排步骤类型(开始、脚本、人工干预、结束等)"`
	Name            string `json:"name"  gorm:"comment:步骤名称"`
	StepNum         int    `json:"step_num" gorm:"comment:作业编排步骤Id"`
	PreviousStepNum int    `json:"previous_step_num" gorm:"comment:作业编排上一步骤Id"`
	NextStepNum     int    `json:"next_step_num" gorm:"comment:作业编排下一步骤Id"`
}

type TaskTemplateStepScript struct {
	ID                  int             `json:"id" gorm:"primaryKey;autoIncrement"`
	TemplateId          int             `json:"template_id" gorm:"comment:作业编排Id"`
	StepNum             int             `json:"step_num" gorm:"comment:作业编排步骤Id"`
	ScriptType          string          `json:"script_type"  gorm:"comment:脚本类型"`
	ScriptId            string          `json:"script_id"  gorm:"comment:引用脚本Id"`
	ScriptVersionId     string          `json:"script_version_id"  gorm:"comment:引用脚本版本Id"`
	ScriptContent       string          `json:"script_content"  gorm:"comment:脚本内容"`
	ScriptParam         string          `json:"script_param"  gorm:"comment:脚本执行参数"`
	ScriptTimeout       int             `json:"script_timeout"  gorm:"comment:脚本超时"`
	DestinationHostList json.RawMessage `json:"destination_host_list" gorm:"comment:远程执行主机列表"`
}

type TaskTemplateDTO struct {
	Template  TaskTemplate             `json:"template"`
	Variables []TaskTemplateVariable   `json:"variables"`
	Steps     []TaskTemplateStep       `json:"steps"`
	Scripts   []TaskTemplateStepScript `json:"scripts"`
}
