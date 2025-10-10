package rule

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum"
)

type ActionType int

const (
	Block   ActionType = 1
	Warning ActionType = 2
)

var ActionMap = enum.EnumMap{
	int(Block):   "拦截",
	int(Warning): "警告",
}

func (s ActionType) String() string {
	return ActionMap.String(int(s))
}

func ParseActionType(s string) ActionType {
	for k, v := range ActionMap {
		if v == s {
			return ActionType(k)
		}
	}
	return 0
}

func (p ActionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(ActionMap[int(p)])
}

func (p *ActionType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	for k, v := range ActionMap {
		if v == s {
			*p = ActionType(k)
			return nil
		}
	}
	return fmt.Errorf("invalid ActionType: '%s', allowed: %v", s, ActionMap)
}

func (p ActionType) Value() (driver.Value, error) {
	return int64(p), nil
}

func (p *ActionType) Scan(value interface{}) error {
	if value == nil {
		*p = 0
		return nil
	}
	if v, ok := value.(int64); ok {
		*p = ActionType(int(v))
		return nil
	}
	return nil
}

func GetActions() []enum.Item {
	return ActionMap.ToItems()
}
