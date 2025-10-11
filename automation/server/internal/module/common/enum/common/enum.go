package common

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"sort"
)

type EnumMap map[int]string

func (m EnumMap) String(id int) string {
	if v, ok := m[id]; ok {
		return v
	}
	return "未知"
}

type Item struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

func (m EnumMap) ToItems() []Item {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	items := make([]Item, 0, len(keys))
	for _, k := range keys {
		items = append(items, Item{ID: k, Type: m[k]})
	}
	return items
}

type MultiEnum []int

func (a MultiEnum) Strings(m EnumMap) []string {
	out := make([]string, 0, len(a))
	for _, v := range a {
		out = append(out, m.String(v))
	}
	return out
}

func (a MultiEnum) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *MultiEnum) Scan(value interface{}) error {
	if value == nil {
		*a = nil
		return nil
	}
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("cannot scan %T into MultiEnum", value)
	}
	return json.Unmarshal(bytes, a)
}
