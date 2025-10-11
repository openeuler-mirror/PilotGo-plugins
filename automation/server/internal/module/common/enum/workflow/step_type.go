package workflow

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum/common"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/pkg/response"
)

type StepType int

const (
	ExecTask       StepType = 1
	ManualReview   StepType = 2
	ProcessControl StepType = 3
)

var StepTypeMap = common.EnumMap{
	int(ExecTask):       "任务",
	int(ManualReview):   "人工处理",
	int(ProcessControl): "流程控制",
}

func (p StepType) String() string {
	return StepTypeMap.String(int(p))
}
func ParseStepType(s string) StepType {
	for k, v := range StepTypeMap {
		if v == s {
			return StepType(k)
		}
	}
	return 0
}

func (p StepType) MarshalJSON() ([]byte, error) {
	return json.Marshal(StepTypeMap[int(p)])
}

func (p *StepType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	for k, v := range StepTypeMap {
		if v == s {
			*p = StepType(k)
			return nil
		}
	}
	return fmt.Errorf("invalid StepType: '%s', allowed: %v", s, StepTypeMap)
}

func (p StepType) Value() (driver.Value, error) {
	return int64(p), nil
}

func (p *StepType) Scan(value interface{}) error {
	if value == nil {
		*p = 0
		return nil
	}
	if v, ok := value.(int64); ok {
		*p = StepType(int(v))
		return nil
	}
	return nil
}

func getStepType() []common.Item {
	return StepTypeMap.ToItems()
}

func StepTypeListHandler(c *gin.Context) {
	stepTypes := getStepType()
	response.Success(c, stepTypes, "success")
}
