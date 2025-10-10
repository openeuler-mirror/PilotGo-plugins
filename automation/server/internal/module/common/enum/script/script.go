package script

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum"
)

type ScriptType int

const (
	Shell  ScriptType = 1
	Perl   ScriptType = 2
	Python ScriptType = 3
	SQL    ScriptType = 4
)

var ScriptTypeMap = enum.EnumMap{
	int(Shell):  "Shell",
	int(Perl):   "Perl",
	int(Python): "Python",
	int(SQL):    "SQL",
}

func (s ScriptType) String() string {
	return ScriptTypeMap.String(int(s))
}

func ParseScriptType(s string) ScriptType {
	for k, v := range ScriptTypeMap {
		if v == s {
			return ScriptType(k)
		}
	}
	return 0
}

func (p ScriptType) MarshalJSON() ([]byte, error) {
	return json.Marshal(ScriptTypeMap[int(p)])
}

func (p *ScriptType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	for k, v := range ScriptTypeMap {
		if v == s {
			*p = ScriptType(k)
			return nil
		}
	}
	return fmt.Errorf("invalid ScriptType: '%s', allowed: %v", s, ScriptTypeMap)
}

func (p ScriptType) Value() (driver.Value, error) {
	return int64(p), nil
}

func (p *ScriptType) Scan(value interface{}) error {
	if value == nil {
		*p = 0
		return nil
	}
	if v, ok := value.(int64); ok {
		*p = ScriptType(int(v))
		return nil
	}
	return nil
}

func GetScriptType() []enum.Item {
	return ScriptTypeMap.ToItems()
}
