package service

import (
	"fmt"
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
	for _, node1 := range nodes.Nodes {
		if node1.UUID == uuid {
			repeat_node := false
			for _, node2 := range single_nodes {
				if node2.ID == node1.ID {
					repeat_node = true
				}
			}

			if !repeat_node {
				single_nodes = append(single_nodes, node1)
			}
		}
	}

	single_edges := []*meta.Edge{}
	for _, edge1 := range edges.Edges {
		if strings.Split(edge1.Src, "_")[0] == uuid {
			repeat_edge := false
			for _, edge2 := range single_edges {
				if edge2.ID == edge1.ID {
					repeat_edge = true
				}
			}

			if !repeat_edge {
				single_edges = append(single_edges, edge1)
			}
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

	hostids := []string{}
	multi_nodes := []*meta.Node{}
	for _, node := range nodes.Nodes {
		if node.Type == "host" {
			multi_nodes = append(multi_nodes, node)
			hostids = append(hostids, node.ID)
		}
	}

	// ttcode
	fmt.Printf("\033[32mhostids\033[0m: %v\n", hostids)

	multi_edges := []*meta.Edge{}
	for _, edge := range edges.Edges {
		if edge.Type == "tcp" || edge.Type == "udp" {
			multi_edges = append(multi_edges, edge)

			repeat_src := false
			repeat_dst := false
			for _, node := range multi_nodes {
				if node.ID == nodes.Lookup[edge.Src].ID {
					repeat_src = true
				}

				if node.ID == nodes.Lookup[edge.Dst].ID {
					repeat_dst = true
				}
			}

			if !repeat_src {
				multi_nodes = append(multi_nodes, nodes.Lookup[edge.Src])
			}

			if !repeat_dst {
				multi_nodes = append(multi_nodes, nodes.Lookup[edge.Dst])
			}
		} else if edge.Type == "server" || edge.Type == "client" {
			multi_edges = append(multi_edges, edge)

			repeat_src := false
			repeat_dst := false
			for _, node := range multi_nodes {
				if node.ID == nodes.Lookup[edge.Src].ID {
					repeat_src = true
				}

				if node.ID == nodes.Lookup[edge.Dst].ID {
					repeat_dst = true
				}
			}

			if !repeat_src {
				multi_nodes = append(multi_nodes, nodes.Lookup[edge.Src])
			}

			if !repeat_dst {
				multi_nodes = append(multi_nodes, nodes.Lookup[edge.Dst])
			}

			// 添加process-to-process、process-to-host edge
			start_nodeid := edge.Dst
			for {
				if nodes.Lookup[start_nodeid].Metrics["Ppid"] != "1" {
					// 添加 process-to-process edge
					edgeid := start_nodeid + "_belong_" + nodes.Lookup[start_nodeid].UUID + "_process_" + nodes.Lookup[start_nodeid].Metrics["Ppid"]
					edge1, ok := edges.Lookup.Load(edgeid)
					if !ok {
						fmt.Printf("%+v\n", errors.Errorf("faild to load edge from edges.lookup: %s**2", edgeid))
					}

					repeat_edge := false
					for _, edge2 := range multi_edges {
						if edge2.ID == edge1.(*meta.Edge).ID {
							repeat_edge = true
						}
					}

					if !repeat_edge {
						multi_edges = append(multi_edges, edge1.(*meta.Edge))
					}

					start_nodeid = nodes.Lookup[start_nodeid].UUID + "_process_" + nodes.Lookup[start_nodeid].Metrics["Ppid"]

					continue
				}

				// 添加 process-to-1 edge
				edgeid_to_1 := start_nodeid + "_belong_" + nodes.Lookup[start_nodeid].UUID + "_process_" + nodes.Lookup[start_nodeid].Metrics["Ppid"]
				edge_to_1, ok := edges.Lookup.Load(edgeid_to_1)
				if !ok {
					fmt.Printf("%+v\n", errors.Errorf("faild to load edge from edges.lookup: %s**2", edgeid_to_1))
				}

				repeat_edge_to_1 := false
				for _, edge2 := range multi_edges {
					if edge2.ID == edge_to_1.(*meta.Edge).ID {
						repeat_edge_to_1 = true
					}
				}

				if !repeat_edge_to_1 {
					multi_edges = append(multi_edges, edge_to_1.(*meta.Edge))
				}

				// 添加 1-to-host edge
				start_nodeid = nodes.Lookup[start_nodeid].UUID + "_process_" + nodes.Lookup[start_nodeid].Metrics["Ppid"]
				edgeid_to_host := ""
				for _, hostid := range hostids {
					if nodes.Lookup[start_nodeid].UUID == nodes.Lookup[hostid].UUID {
						edgeid_to_host = start_nodeid + "_belong_" + nodes.Lookup[hostid].ID
						break
					}
				}

				edge_to_host, ok := edges.Lookup.Load(edgeid_to_host)
				if !ok {
					fmt.Printf("%+v\n", errors.Errorf("faild to load edge from edges.lookup: %s**2", edgeid_to_host))
				}

				repeat_edge_to_host := false
				for _, edge2 := range multi_edges {
					if edge2.ID == edge_to_host.(*meta.Edge).ID {
						repeat_edge_to_host = true
					}
				}

				if !repeat_edge_to_host {
					multi_edges = append(multi_edges, edge_to_host.(*meta.Edge))
				}

				break
			}

		}
	}

	return multi_nodes, multi_edges, collect_errlist, process_errlist
}
