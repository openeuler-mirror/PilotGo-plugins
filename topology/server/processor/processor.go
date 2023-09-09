package processor

import (
	"fmt"
	"strconv"
	"sync"

	"gitee.com/openeuler/PilotGo-plugin-topology-server/agentmanager"
	"gitee.com/openeuler/PilotGo-plugin-topology-server/collector"
	"gitee.com/openeuler/PilotGo-plugin-topology-server/meta"
	"gitee.com/openeuler/PilotGo-plugin-topology-server/utils"
	"github.com/pkg/errors"
)

type DataProcesser struct{}

func CreateDataProcesser() *DataProcesser {
	return &DataProcesser{}
}

func (d *DataProcesser) Process_data() (*meta.Nodes, *meta.Edges, []error, []error) {
	var Nodes *meta.Nodes
	var Edges *meta.Edges
	var wg sync.WaitGroup
	var collect_errorlist []error
	var process_errorlist []error

	datacollector := collector.CreateDataCollector()
	collect_errorlist = datacollector.Collect_instant_data()
	if len(collect_errorlist) != 0 {
		for i, err := range collect_errorlist {
			collect_errorlist[i] = errors.Wrap(err, "**7")
		}
	}

	agentmanager.Topo.AgentMap.Range(
		func(key, value any) bool {
			wg.Add(1)

			go func() {
				defer wg.Done()
				agent := value.(*agentmanager.Agent_m)

				err := d.Create_node_entities(agent, Nodes)
				if err != nil {
					process_errorlist = append(process_errorlist, errors.Wrap(err, "**2"))
				}

				err = d.Create_edge_entities(agent, Edges, Nodes)
				if err != nil {
					process_errorlist = append(process_errorlist, errors.Wrap(err, "**2"))
				}
			}()

			return true
		},
	)
	wg.Wait()

	return Nodes, Edges, collect_errorlist, process_errorlist
}

func (d *DataProcesser) Create_node_entities(agent *agentmanager.Agent_m, nodes *meta.Nodes) error {
	host_node := &meta.Node{
		ID:      fmt.Sprintf("%s-%s-%s", agent.UUID, meta.NODE_HOST, agent.IP),
		Name:    agent.UUID,
		Type:    meta.NODE_HOST,
		UUID:    agent.UUID,
		Metrics: *utils.HostToMap(agent.Host_2, &agent.AddrInterfaceMap_2),
	}

	nodes.Add(host_node)

	for _, process := range agent.Processes_2 {
		proc_node := &meta.Node{
			ID:      fmt.Sprintf("%s-%s-%d", agent.UUID, meta.NODE_PROCESS, process.Pid),
			Name:    process.ExeName,
			Type:    meta.NODE_PROCESS,
			UUID:    agent.UUID,
			Metrics: *utils.ProcessToMap(process),
		}

		nodes.Add(proc_node)

		for _, thread := range process.Threads {
			thre_node := &meta.Node{
				ID:      fmt.Sprintf("%s-%s-%d", agent.UUID, meta.NODE_THREAD, thread.Tid),
				Name:    strconv.Itoa(int(thread.Tid)),
				Type:    meta.NODE_THREAD,
				UUID:    agent.UUID,
				Metrics: *utils.ThreadToMap(&thread),
			}

			nodes.Add(thre_node)
		}

		// for _, net := range process.NetIOCounters {
		// 	net_node := &meta.Node{
		// 		ID:      fmt.Sprintf("%s-%s-%d", agent.UUID, meta.NODE_NET, process.Pid),
		// 		Name:    net.Name,
		// 		Type:    meta.NODE_NET,
		// 		UUID:    agent.UUID,
		// 		Metrics: *utils.NetToMap(&net, &agent.AddrInterfaceMap_2),
		// 	}

		// 	nodes.Add(net_node)
		// }
	}

	// 临时定义不含网络流量metric的网络节点
	for _, net := range agent.Netconnections_2 {
		net_node := &meta.Node{
			ID:      fmt.Sprintf("%s-%s-%d", agent.UUID, meta.NODE_NET, net.Pid),
			Name:    net.Laddr,
			Type:    meta.NODE_NET,
			UUID:    agent.UUID,
			Metrics: *utils.NetToMap(net),
		}

		nodes.Add(net_node)
	}

	return nil
}

func (d *DataProcesser) Create_edge_entities(agent *agentmanager.Agent_m, edges *meta.Edges, nodes *meta.Nodes) error {
	nodes_map := map[string][]*meta.Node{}

	for _, node := range nodes.Nodes {
		switch node.Type {
		case meta.NODE_HOST:
			nodes_map[meta.NODE_HOST] = append(nodes_map[meta.NODE_HOST], node)
		case meta.NODE_PROCESS:
			nodes_map[meta.NODE_PROCESS] = append(nodes_map[meta.NODE_PROCESS], node)
		case meta.NODE_THREAD:
			nodes_map[meta.NODE_THREAD] = append(nodes_map[meta.NODE_THREAD], node)
		case meta.NODE_NET:
			nodes_map[meta.NODE_NET] = append(nodes_map[meta.NODE_NET], node)
		}
	}

	for _, obj := range nodes_map[meta.NODE_HOST] {
		for _, sub := range nodes_map[meta.NODE_PROCESS] {
			if sub.Metrics["Ppid"] == "1" && sub.UUID == obj.UUID {
				belong_edge := &meta.Edge{
					ID:   fmt.Sprintf("%s_%s_%s", sub.ID, meta.EDGE_BELONG, obj.ID),
					Type: meta.EDGE_BELONG,
					Src:  sub.ID,
					Dst:  obj.ID,
					Dir:  true,
				}

				edges.Add(belong_edge)
			}
		}
	}

	return nil
}
