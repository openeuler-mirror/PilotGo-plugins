package enum

import (
	"encoding/json"
	"sort"
)

type Item struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

type MapWrapper map[int]string

func (m MapWrapper) MarshalJSON() ([]byte, error) {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	items := make([]Item, 0, len(keys))
	for _, k := range keys {
		items = append(items, Item{
			ID:   k,
			Type: m[k],
		})
	}
	return json.Marshal(items)
}
