package handler

import (
	"fmt"
	"net/http"

	"gitee.com/openeuler/PilotGo-plugin-topology-server/service"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func SingleHostHandle(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	nodes, edges, collect_errlist, process_errlist := service.SingleHostService(uuid)

	if len(collect_errlist) != 0 || len(process_errlist) != 0 {
		for i, cerr := range collect_errlist {
			collect_errlist[i] = errors.Wrap(cerr, "**4")
			fmt.Printf("%+v\n", collect_errlist[i])
			// errors.EORE(collect_errlist[i])
		}

		for i, perr := range process_errlist {
			process_errlist[i] = errors.Wrap(perr, "**10")
			fmt.Printf("%+v\n", perr)
			// errors.EORE(process_errlist[i])
		}
	}

	if len(nodes) == 0 || len(edges) == 0 {
		err := errors.New("nodes list is null or edges list is null**0")
		fmt.Printf("%+v\n", err)
		// errors.EORE(err)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  -1,
			"error": err,
			"data":  nil,
		})
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

func MultiHostHandle(ctx *gin.Context) {
	nodes, edges, collect_errlist, process_errlist := service.MultiHostService()

	if len(collect_errlist) != 0 || len(process_errlist) != 0 {
		for i, cerr := range collect_errlist {
			collect_errlist[i] = errors.Wrap(cerr, "**4")
			fmt.Printf("%+v\n", collect_errlist[i])
			// errors.EORE(collect_errlist[i])
		}

		for i, perr := range process_errlist {
			process_errlist[i] = errors.Wrap(perr, "**10")
			fmt.Printf("%+v\n", perr)
			// errors.EORE(process_errlist[i])
		}
	}

	if len(nodes) == 0 || len(edges) == 0 {
		err := errors.New("nodes list is null or edges list is null**0")
		fmt.Printf("%+v\n", err)
		// errors.EORE(err)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  -1,
			"error": err,
			"data":  nil,
		})
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
