package httphandler

import (
	"log"

	"github.com/docker/docker/client"
	docker "github.com/fsouza/go-dockerclient"
	"github.com/gin-gonic/gin"
)

func ContainerList(ctx *gin.Context) {
	client, err := docker.NewClient("tcp://127.0.0.1:2375")
	if err != nil {
		log.Fatal(err)
	}

	containers, err := client.ListContainers(docker.ListContainersOptions{All: true})
	if err != nil {
		log.Fatal(err)
	}

	for _, container := range containers {
		log.Printf("ID: %s, Image: %s, Status: %s\n", container.ID[:12], container.Image, container.Status)
	}
}

func RunContainer(ctx *gin.Context) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

}

func ExecCmdInContainer(ctx *gin.Context) {

}

func MigrateContainer(ctx *gin.Context) {

}
