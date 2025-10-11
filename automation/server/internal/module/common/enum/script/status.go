package script

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum"
)

type ScriptVersionStatus int

const (
	Develop   ScriptVersionStatus = 1
	Published ScriptVersionStatus = 2
)

var ScriptVersionStatusMap = enum.EnumMap{
	int(Develop):   "开发中",
	int(Published): "已发布",
}

func (p ScriptVersionStatus) String() string {
	return ScriptVersionStatusMap.String(int(p))
}
func ParseScriptVersionStatus(s string) ScriptVersionStatus {
	for k, v := range ScriptVersionStatusMap {
		if v == s {
			return ScriptVersionStatus(k)
		}
	}
	return 0
}

func (p ScriptVersionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(ScriptVersionStatusMap[int(p)])
}

func (p *ScriptVersionStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		for k, v := range ScriptVersionStatusMap {
			if v == s {
				*p = ScriptVersionStatus(k)
				return nil
			}
		}
		return fmt.Errorf("invalid ScriptVersionStatus: '%s', allowed: %v", s, ScriptVersionStatusMap)
	}

	var num int
	if err := json.Unmarshal(data, &num); err == nil {
		if _, exists := ScriptVersionStatusMap[num]; exists {
			*p = ScriptVersionStatus(num)
			return nil
		}
		return fmt.Errorf("invalid ScriptVersionStatus: %d, allowed: %v", num, ScriptVersionStatusMap)
	}

	return fmt.Errorf("invalid ScriptVersionStatus, must be string or number")
}

func (p ScriptVersionStatus) Value() (driver.Value, error) {
	return int64(p), nil
}

func (p *ScriptVersionStatus) Scan(value interface{}) error {
	if value == nil {
		*p = 0
		return nil
	}
	if v, ok := value.(int64); ok {
		*p = ScriptVersionStatus(int(v))
		return nil
	}
	return nil
}

func GetScriptVersionStatus() []enum.Item {
	return ScriptVersionStatusMap.ToItems()
}
