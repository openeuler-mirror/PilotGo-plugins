package model

import (
	"encoding/json"

	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum/workflow"
)

type TaskTemplate struct {
	ID            int                    `json:"id" gorm:"primaryKey;autoIncrement;comment:作业编排Id"`
	Name          string                 `json:"name" gorm:"type:varchar(255);comment:作业编排名称"`
	Description   string                 `json:"description" gorm:"type:varchar(255);comment:作业描述"`
	Tags          string                 `json:"tags" gorm:"comment:场景标签"`
	PublishStatus workflow.PublishStatus `json:"publish_status" gorm:"comment:发布状态"`
	FirstStepNum  int                    `json:"first_step_num"`
	LastStepNum   int                    `json:"last_step_num"`
	ModifyUser    string                 `json:"modify_user" gorm:"type:varchar(100);not null;comment:'最后修改者'"`
	ModifyTime    string                 `json:"modify_time" gorm:"comment:'最后修改时间'"`
}

type TaskTemplateParams struct {
	ID          int             `json:"id" gorm:"primaryKey;autoIncrement"`
	TemplateId  int             `json:"template_id" gorm:"comment:作业编排Id"`
	Name        string          `json:"name" gorm:"type:varchar(255);comment:输入参数名称"`
	Type        string          `json:"type" gorm:"comment:参数类型(string、number、host、json等)"`
	DefaultVaue json.RawMessage `json:"default_value" gorm:"comment:参数默认值"`
	Description string          `json:"description" gorm:"comment:输入参数描述"`
	IsRequired  bool            `json:"is_required" gorm:"comment:是否必需"`
}

type TaskTemplateOutputParams struct {
	ID          int    `json:"id" gorm:"primaryKey;autoIncrement"`
	TemplateId  int    `json:"template_id" gorm:"comment:作业编排Id"`
	Name        string `json:"name" gorm:"type:varchar(255);comment:输出参数名称"`
	Type        string `json:"type" gorm:"comment:参数类型string"`
	Description string `json:"description" gorm:"comment:输出参数描述"`
}

type TaskTemplateStep struct {
	ID              int               `json:"id" gorm:"primaryKey;autoIncrement"`
	TemplateId      int               `json:"template_id" gorm:"comment:作业编排Id"`
	StepType        workflow.StepType `json:"step_type" gorm:"comment:编排步骤类型"`
	Name            string            `json:"name"  gorm:"comment:步骤名称"`
	StepNum         int               `json:"step_num" gorm:"comment:作业编排步骤Id"`
	PreviousStepNum int               `json:"previous_step_num" gorm:"comment:作业编排上一步骤Id"`
	NextStepNum     int               `json:"next_step_num" gorm:"comment:作业编排下一步骤Id"`
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
	Template     TaskTemplate               `json:"template"`
	Params       []TaskTemplateParams       `json:"params"`
	OutputParams []TaskTemplateOutputParams `json:"output_params"`
	Steps        []TaskTemplateStep         `json:"steps"`
	Scripts      []TaskTemplateStepScript   `json:"scripts"`
}
