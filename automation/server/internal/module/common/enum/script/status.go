package script

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum/common"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/pkg/response"
)

type ScriptPublishStatus int

const (
	Develop   ScriptPublishStatus = 1
	Published ScriptPublishStatus = 2
)

var ScriptPublishStatusMap = common.EnumMap{
	int(Develop):   "开发中",
	int(Published): "已发布",
}

func (p ScriptPublishStatus) String() string {
	return ScriptPublishStatusMap.String(int(p))
}
func ParseScriptPublishStatus(s string) ScriptPublishStatus {
	for k, v := range ScriptPublishStatusMap {
		if v == s {
			return ScriptPublishStatus(k)
		}
	}
	return 0
}

func (p ScriptPublishStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(ScriptPublishStatusMap[int(p)])
}

func (p *ScriptPublishStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		for k, v := range ScriptPublishStatusMap {
			if v == s {
				*p = ScriptPublishStatus(k)
				return nil
			}
		}
		return fmt.Errorf("invalid ScriptPublishStatus: '%s', allowed: %v", s, ScriptPublishStatusMap)
	}

	var num int
	if err := json.Unmarshal(data, &num); err == nil {
		if _, exists := ScriptPublishStatusMap[num]; exists {
			*p = ScriptPublishStatus(num)
			return nil
		}
		return fmt.Errorf("invalid ScriptPublishStatus: %d, allowed: %v", num, ScriptPublishStatusMap)
	}

	return fmt.Errorf("invalid ScriptPublishStatus, must be string or number")
}

func (p ScriptPublishStatus) IsValid() bool {
	_, exists := ScriptPublishStatusMap[int(p)]
	return exists
}

func (p ScriptPublishStatus) Value() (driver.Value, error) {
	if !p.IsValid() {
		return nil, fmt.Errorf("invalid ScriptPublishStatus: %d, allowed: %v", int(p), ScriptPublishStatusMap)
	}
	return int64(p), nil
}

func (p *ScriptPublishStatus) Scan(value interface{}) error {
	if value == nil {
		*p = 0
		return nil
	}
	if v, ok := value.(int64); ok {
		status := ScriptPublishStatus(int(v))
		if !status.IsValid() {
			return fmt.Errorf("invalid ScriptPublishStatus: %d, allowed: %v", int(v), ScriptPublishStatusMap)
		}
		*p = status
		return nil
	}
	return nil
}

func getScriptPublishStatus() []common.Item {
	return ScriptPublishStatusMap.ToItems()
}

func ScriptPublishStatusListHandler(c *gin.Context) {
	status := getScriptPublishStatus()
	response.Success(c, status, "success")
}
