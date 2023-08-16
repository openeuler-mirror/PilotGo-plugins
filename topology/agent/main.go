package main

import (
	"fmt"
	"os"

	"gitee.com/openeuler/PilotGo-plugin-topology-agent/collector"
	"gitee.com/openeuler/PilotGo-plugin-topology-agent/conf"
	"gitee.com/openeuler/PilotGo-plugin-topology-agent/handler"
	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	psutilcollector := &collector.PsutilCollector{}
	if err := psutilcollector.Get_host_info(); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", psutilcollector.Host_0)

	os.Exit(1)

	InitLogger()

	engine := gin.Default()
	handler.InitRouter(engine)
	if err := engine.Run(conf.Config().Http.Agent_addr); err != nil {
		logger.Fatal("failed to run server")
	}

}

func InitLogger() {
	if err := logger.Init(conf.Config().Logopts); err != nil {
		fmt.Printf("logger init failed, please check the config file: %s", err)
		os.Exit(1)
	}
}
