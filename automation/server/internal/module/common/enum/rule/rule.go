package rule

import "openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum"

type Severity int

const (
	SeverityBlock   Severity = 1
	SeverityWarning Severity = 2
)

var SeverityMap = enum.EnumMap{
	int(SeverityBlock):   "拦截",
	int(SeverityWarning): "警告",
}

func (s Severity) String() string {
	return SeverityMap.String(int(s))
}

func GetSeverities() []enum.Item {
	return SeverityMap.ToItems()
}
