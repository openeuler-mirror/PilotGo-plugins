package meta

import "fmt"

type Nodes struct {
	Lookup map[string]*Node
	Nodes  []*Node
}

type Node struct {
	ID      int
	Name    string
	Type    string
	Attrs   map[string]string
	Metrics map[string]string
}

func NewNodes() *Nodes {
	return &Nodes{make(map[string]*Node), make([]*Node, 0)}
}

func (ns *Nodes) Add(node *Node) {
	old_n, ok := ns.Lookup[node.Name]
	if ok {
		for k, v := range node.Attrs {
			if _, ok := old_n.Attrs[k]; !ok {
				old_n.Attrs[k] = v
			}
		}

		return
	}

	ns.Lookup[node.Name] = node
	ns.Nodes = append(ns.Nodes, node)
}

func (ns *Nodes) Remove(name string) error {
	for i := 0; i < len(ns.Nodes); i++ {
		if ns.Nodes[i].Name != name {
			continue
		}
		ns.Nodes = append(ns.Nodes[:i], ns.Nodes[i+1:]...)
		delete(ns.Lookup, name)
		return nil
	}
	return fmt.Errorf("node %s not found", name)
}
