package meta

import (
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
	Src     string
	Dst     string
	UUID    string
	Dir     bool
	Metrics map[string]string
}

func (e *Edges) Add(edge *Edge) {
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
