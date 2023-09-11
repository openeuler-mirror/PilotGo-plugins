package service

import (
	"strings"

	"gitee.com/openeuler/PilotGo-plugin-topology-server/meta"
	"gitee.com/openeuler/PilotGo-plugin-topology-server/processor"
	"github.com/pkg/errors"
)

func SingleHostService(uuid string) ([]*meta.Node, []*meta.Edge, []error, []error) {
	dataprocesser := processor.CreateDataProcesser()
	nodes, edges, collect_errlist, process_errlist := dataprocesser.Process_data()
	if len(collect_errlist) != 0 || len(process_errlist) != 0 {
		for i, cerr := range collect_errlist {
			collect_errlist[i] = errors.Wrap(cerr, "**3")
		}

		for i, perr := range process_errlist {
			process_errlist[i] = errors.Wrap(perr, "**7")
		}
	}

	single_nodes := []*meta.Node{}
	for _, node := range nodes.Nodes {
		if node.UUID == uuid {
			single_nodes = append(single_nodes, node)
		}
	}

	single_edges := []*meta.Edge{}
	for _, edge := range edges.Edges {
		if strings.Split(edge.Src, "_")[0] == uuid {
			single_edges = append(single_edges, edge)
		}
	}

	return single_nodes, single_edges, collect_errlist, process_errlist
	// if len(collect_errlist) != 0 && len(process_errlist) != 0 {
	// 	for i, cerr := range collect_errlist {
	// 		collect_errlist[i] = errors.Wrap(cerr, "**3")
	// 	}

	// 	for i, perr := range process_errlist {
	// 		process_errlist[i] = errors.Wrap(perr, "**7")
	// 	}

	// 	return nil, nil, collect_errlist, process_errlist
	// } else if len(collect_errlist) != 0 && len(process_errlist) == 0 {
	// 	for i, cerr := range collect_errlist {
	// 		collect_errlist[i] = errors.Wrap(cerr, "**3")
	// 	}

	// 	return nil, nil, collect_errlist, nil
	// } else if len(collect_errlist) == 0 && len(process_errlist) != 0 {
	// 	for i, perr := range process_errlist {
	// 		process_errlist[i] = errors.Wrap(perr, "**7")
	// 	}

	// 	return nil, nil, nil, process_errlist
	// }
}

func MultiHostService() ([]*meta.Node, []*meta.Edge, []error, []error) {
	dataprocesser := processor.CreateDataProcesser()
	nodes, edges, collect_errlist, process_errlist := dataprocesser.Process_data()
	if len(collect_errlist) != 0 || len(process_errlist) != 0 {
		for i, cerr := range collect_errlist {
			collect_errlist[i] = errors.Wrap(cerr, "**3")
		}

		for i, perr := range process_errlist {
			process_errlist[i] = errors.Wrap(perr, "**7")
		}
	}

	multi_nodes := []*meta.Node{}
	for _, node := range nodes.Nodes {
		if node.Type == "host" {
			multi_nodes = append(multi_nodes, node)
		}
	}

	multi_edges := []*meta.Edge{}
	for _, edge := range edges.Edges {
		if edge.Type == "server" || edge.Type == "client" || edge.Type == "tcp" || edge.Type == "udp" {
			multi_edges = append(multi_edges, edge)

			multi_nodes = append(multi_nodes, nodes.Lookup[edge.Src])
			multi_nodes = append(multi_nodes, nodes.Lookup[edge.Dst])
		}
	}

	return multi_nodes, multi_edges, collect_errlist, process_errlist
}
