package meta

import (
	"strings"
	"sync"

	"github.com/pkg/errors"
)

type Edges struct {
	SrcToDsts map[string][]string
	DstToSrcs map[string][]string
	Lookup    sync.Map
	Edges     []*Edge
}

type Edge struct {
	ID      string
	Type    string
	Src     string `json:"source"`
	Dst     string `json:"target"`
	Dir     bool
	Metrics map[string]string
}

// 镜像id检测：多个goruntine并发添加、访问、修改相同的edge实例
func (e *Edges) Add(edge *Edge) {
	id_slice := strings.Split(edge.ID, "_")
	id_slice[0], id_slice[2] = id_slice[2], id_slice[0]

	mirror_id := strings.Join(id_slice, "_")

	if _, ok := e.Lookup.Load(mirror_id); ok {
		return
	}

	e.Lookup.Store(edge.ID, edge)
	e.Edges = append(e.Edges, edge)
}

func (e *Edges) Remove(id string) error {
	for i := 0; i < len(e.Edges); i++ {
		if e.Edges[i].ID != id {
			continue
		}
		e.Edges = append(e.Edges[:i], e.Edges[i+1:]...)
		if _, ok := e.Lookup.LoadAndDelete(id); !ok {
			return errors.Errorf("edge %+v not fount in sync.map**1", id)
		}

		return nil
	}

	return errors.Errorf("edge %+v not fount in slice**12", id)
}
