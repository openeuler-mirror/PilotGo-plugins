package utils

import "reflect"

func StructToMap(obj interface{}) map[string]string {
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}

	if objValue.Kind() != reflect.Struct {
		return nil
	}

	objType := objValue.Type()
	fieldCount := objType.NumField()

	m := make(map[string]string)
	for i := 0; i < fieldCount; i++ {
		field := objType.Field(i)
		fieldValue := objValue.Field(i)

		m[field.Name] = fieldValue.Interface().(string)
	}

	return m
}
