package main

import (
	"fmt"

	"openeuler.org/PilotGo/gala-ops-plugin/client"
	"openeuler.org/PilotGo/gala-ops-plugin/config"
)

func main() {
	fmt.Println("hello gala-ops")

	config.Init()

	client.StartClient(config.Config().Http)
}
