package httphandler

import (
	"log"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/gin-gonic/gin"
)

func GetContainerList(ctx *gin.Context) {
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

}

func ExecCmdInContainer(ctx *gin.Context) {

}

func MigrateContainer(ctx *gin.Context) {

}
