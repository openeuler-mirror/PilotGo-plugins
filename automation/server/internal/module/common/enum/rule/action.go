package rule

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum/common"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/pkg/response"
)

type ActionType int

const (
	Block   ActionType = 1
	Warning ActionType = 2
)

var ActionMap = common.EnumMap{
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

func (p ActionType) IsValid() bool {
	_, exists := ActionMap[int(p)]
	return exists
}

func (p ActionType) Value() (driver.Value, error) {
	if !p.IsValid() {
		return nil, fmt.Errorf("invalid ActionType: %d, allowed: %v", int(p), ActionMap)
	}
	return int64(p), nil
}

func (p *ActionType) Scan(value interface{}) error {
	if value == nil {
		*p = 0
		return nil
	}
	if v, ok := value.(int64); ok {
		status := ActionType(int(v))
		if !status.IsValid() {
			return fmt.Errorf("invalid ActionType: %d, allowed: %v", int(v), ActionMap)
		}
		*p = status
		return nil
	}
	return nil
}

func getRuleActions() []common.Item {
	return ActionMap.ToItems()
}

func RuleActionListHandler(c *gin.Context) {
	actions := getRuleActions()
	response.Success(c, actions, "success")
}
