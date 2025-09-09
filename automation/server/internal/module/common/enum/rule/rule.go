package rule

import "openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum"

type Severity int

const (
	SeverityBlock   Severity = 1
	SeverityWarning Severity = 2
)

var SeverityMap = enum.MapWrapper{
	int(SeverityBlock):   "拦截",
	int(SeverityWarning): "警告",
}

func (s Severity) String() string {
	switch s {
	case SeverityBlock:
		return "拦截"
	case SeverityWarning:
		return "警告"
	default:
		return "未定义"
	}
}
