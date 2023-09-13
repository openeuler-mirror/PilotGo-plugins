package handler

import (
	"fmt"
	"net/http"

	"gitee.com/openeuler/PilotGo-plugin-topology-server/agentmanager"
	"gitee.com/openeuler/PilotGo-plugin-topology-server/service"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func SingleHostHandle(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	nodes, edges, collect_errlist, process_errlist := service.SingleHostService(uuid)

	if len(collect_errlist) != 0 || len(process_errlist) != 0 {
		for i, cerr := range collect_errlist {
			collect_errlist[i] = errors.Wrap(cerr, "**4") // err top
			fmt.Printf("%+v\n", collect_errlist[i])
			// errors.EORE(collect_errlist[i])

			// // ttcode
			// agentmanager.Topo.ErrGroup.Add(1)
			// agentmanager.Topo.ErrGroup.Wait()
			// agentmanager.Topo.ErrCh <- collect_errlist[i]
		}

		for i, perr := range process_errlist {
			process_errlist[i] = errors.Wrap(perr, "**10") // err top
			fmt.Printf("%+v\n", process_errlist[i])
			// errors.EORE(process_errlist[i])

			// // ttcode
			// agentmanager.Topo.ErrGroup.Add(1)
			// agentmanager.Topo.ErrGroup.Wait()
			// agentmanager.Topo.ErrCh <- process_errlist[i]
		}
	}

	if len(nodes) == 0 || len(edges) == 0 {
		err := errors.New("nodes list is null or edges list is null**0") // err top
		fmt.Printf("%+v\n", err)
		// errors.EORE(err)

		// // ttcode
		// agentmanager.Topo.ErrGroup.Add(1)
		// agentmanager.Topo.ErrGroup.Wait()
		// agentmanager.Topo.ErrCh <- err

		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  -1,
			"error": err.Error(),
			"data":  nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  0,
		"error": nil,
		"data": map[string]interface{}{
			"nodes": nodes,
			"edges": edges,
		},
	})
}

func SingleHostTreeHandle(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	nodes, collect_errlist, process_errlist := service.SingleHostTreeService(uuid)

	if len(collect_errlist) != 0 || len(process_errlist) != 0 {
		for i, cerr := range collect_errlist {
			collect_errlist[i] = errors.Wrap(cerr, "**4") // err top
			fmt.Printf("%+v\n", collect_errlist[i])
			// errors.EORE(collect_errlist[i])
		}

		for i, perr := range process_errlist {
			process_errlist[i] = errors.Wrap(perr, "**10") // err top
			fmt.Printf("%+v\n", perr)
			// errors.EORE(process_errlist[i])
		}
	}

	if nodes == nil {
		err := errors.New("node tree is null**0") // err top
		fmt.Printf("%+v\n", err)
		// errors.EORE(err)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  -1,
			"error": err.Error(),
			"data":  nil,
		})
		return
	}

	agentmap := make(map[string]string)
	agentmanager.Topo.AgentMap.Range(func(key, value any) bool {
		agent := value.(*agentmanager.Agent_m)
		if agent.Host_2 != nil {
			agentmap[agent.UUID] = agent.IP + ":" + agent.Port
		}

		return true
	})

	ctx.JSON(http.StatusOK, gin.H{
		"code":  0,
		"error": nil,
		"data": map[string]interface{}{
			"tree":      nodes,
			"agentlist": agentmap,
		},
	})
}

func MultiHostHandle(ctx *gin.Context) {
	nodes, edges, collect_errlist, process_errlist := service.MultiHostService()

	if len(collect_errlist) != 0 || len(process_errlist) != 0 {
		for i, cerr := range collect_errlist {
			collect_errlist[i] = errors.Wrap(cerr, "**4") // err top
			fmt.Printf("%+v\n", collect_errlist[i])
			// errors.EORE(collect_errlist[i])
		}

		for i, perr := range process_errlist {
			process_errlist[i] = errors.Wrap(perr, "**10") // err top
			fmt.Printf("%+v\n", perr)
			// errors.EORE(process_errlist[i])
		}
	}

	if len(nodes) == 0 || len(edges) == 0 {
		err := errors.New("nodes list is null or edges list is null**0") // err top
		fmt.Printf("%+v\n", err)
		// errors.EORE(err)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  -1,
			"error": err.Error(),
			"data":  nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  0,
		"error": nil,
		"data": map[string]interface{}{
			"nodes": nodes,
			"edges": edges,
		},
	})
}
