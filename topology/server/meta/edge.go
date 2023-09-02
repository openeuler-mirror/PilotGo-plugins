package meta

import "fmt"

type Edges struct {
	SrcToDsts map[string][]string
	DstToSrcs map[string][]string
	Edges     []*Edge
}

type Edge struct {
	ID    int
	Type  string
	Src   string
	Dst   string
	Dir   bool
	Attrs map[string]string
}

func NewEdges() *Edges {
	return &Edges{make(map[string][]string), make(map[string][]string), make([]*Edge, 0)}
}

func (e *Edges) Add(edge *Edge) {
	e.Edges = append(e.Edges, edge)
}

func (e *Edges) Remove(edge *Edge) error {
	for i := 0; i < len(e.Edges); i++ {
		if e.Edges[i].ID != edge.ID {
			continue
		}
		e.Edges = append(e.Edges[:i], e.Edges[i+1:]...)
		return nil
	}
	return fmt.Errorf("edge %+v not fount", edge)
}
