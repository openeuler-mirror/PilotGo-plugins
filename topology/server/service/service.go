package service

import (
	"gitee.com/openeuler/PilotGo-plugin-topology-server/meta"
	"gitee.com/openeuler/PilotGo-plugin-topology-server/processor"
	"github.com/pkg/errors"
)

func DataProcessService() ([]*meta.Node, []*meta.Edge, []error, []error) {
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

	return nodes.Nodes, edges.Edges, collect_errlist, process_errlist
}
