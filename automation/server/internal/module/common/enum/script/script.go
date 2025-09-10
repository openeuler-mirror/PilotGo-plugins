package script

import (
	"database/sql/driver"

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

type ScriptTypeArr []ScriptType

func (a ScriptTypeArr) Strings() []string {
	intArr := make([]int, len(a))
	for i, v := range a {
		intArr[i] = int(v)
	}
	return enum.MultiEnum(intArr).Strings(enum.EnumMap(ScriptTypeMap))
}

func (a ScriptTypeArr) Value() (driver.Value, error) {
	intArr := make([]int, len(a))
	for i, v := range a {
		intArr[i] = int(v)
	}
	return enum.MultiEnum(intArr).Value()
}

func (a *ScriptTypeArr) Scan(value interface{}) error {
	var m enum.MultiEnum
	if err := m.Scan(value); err != nil {
		return err
	}
	*a = make([]ScriptType, len(m))
	for i, v := range m {
		(*a)[i] = ScriptType(v)
	}
	return nil
}

func NewScriptTypeArr(strs []string) ScriptTypeArr {
	res := make(ScriptTypeArr, 0, len(strs))
	for _, s := range strs {
		res = append(res, ParseScriptType(s))
	}
	return res
}
