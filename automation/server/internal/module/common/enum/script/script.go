package script

import (
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

func GetScriptType() []enum.Item {
	return ScriptTypeMap.ToItems()
}
