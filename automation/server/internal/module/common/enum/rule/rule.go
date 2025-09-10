package rule

import "openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum"

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

func GetActions() []enum.Item {
	return ActionMap.ToItems()
}
