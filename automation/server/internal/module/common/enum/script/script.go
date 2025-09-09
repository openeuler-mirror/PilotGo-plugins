package script

import "openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum"

type ScriptType int

const (
	Shell  ScriptType = 1
	Perl   ScriptType = 2
	Python ScriptType = 3
	SQL    ScriptType = 4
)

var ScriptTypeMap = enum.MapWrapper{
	int(Shell):  "Shell",
	int(Perl):   "Perl",
	int(Python): "Python",
	int(SQL):    "SQL",
}

func (s ScriptType) String() string {
	switch s {
	case Shell:
		return "Shell"
	case Perl:
		return "Perl"
	case Python:
		return "Python"
	case SQL:
		return "SQL"
	default:
		return "未支持"
	}
}
