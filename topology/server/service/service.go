package service

import (
	"gitee.com/openeuler/PilotGo-plugin-topology-server/meta"
	"gitee.com/openeuler/PilotGo-plugin-topology-server/processor"
	"github.com/pkg/errors"
)

func DataProcessService() ([]*meta.Node, []*meta.Edge, []error, []error) {
	dataprocesser := processor.CreateDataProcesser()
	nodes, edges, collect_errlist, process_errlist := dataprocesser.Process_data()

	for i, cerr := range collect_errlist {
		collect_errlist[i] = errors.Wrap(cerr, "**3")
	}

	for i, perr := range process_errlist {
		process_errlist[i] = errors.Wrap(perr, "**7")
	}

	return nodes.Nodes, edges.Edges, collect_errlist, process_errlist
}
