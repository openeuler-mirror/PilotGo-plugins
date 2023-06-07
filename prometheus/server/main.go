package main

import (
	"fmt"
	"os"

	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"gitee.com/openeuler/PilotGo-plugins/sdk/plugin/client"
	"openeuler.org/PilotGo/prometheus-plugin/config"
	"openeuler.org/PilotGo/prometheus-plugin/plugin"
	"openeuler.org/PilotGo/prometheus-plugin/router"
)

func main() {
	fmt.Println("hello prometheus")

	config.Init()

	if err := logger.Init(&config.Config().Logopts); err != nil {
		fmt.Printf("logger init failed, please check the config file: %s", err)
		os.Exit(-1)
	}

	server := router.InitRouter()

	client := client.DefaultClient(plugin.Init(config.Config().Prometheus))
	client.RegisterHandlers(server)

	if err := server.Run(config.Config().Http.Addr); err != nil {
		logger.Fatal("failed to run server")
	}
}
