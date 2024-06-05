package kibanaClient

import (
	"github.com/tidwall/gjson"
)

/*
@varssrc: vars字段值的json数据

将varssrc转换为[]map[string]interface{}数据
*/
func VarsProcess(varssrc gjson.Result) []map[string]interface{} {
	varsdst := []map[string]interface{}{}
	varssrc.ForEach(func(key gjson.Result, value gjson.Result) bool {
		_var := map[string]interface{}{}
		for k, v := range value.Map() {
			if v.IsArray() {
				_arr := []string{}
				for _, s := range v.Array() {
					_arr = append(_arr, s.String())
				}
				_var[k] = _arr
			} else if v.IsBool() {
				_var[k] = v.Bool()
			} else {
				_var[k] = v.String()
			}
		}
		varsdst = append(varsdst, _var)
		return true
	})
	return varsdst
}
