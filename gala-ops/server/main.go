package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"gitee.com/openeuler/PilotGo-plugins/sdk/plugin/client"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/gala-ops-plugin/config"
	"openeuler.org/PilotGo/gala-ops-plugin/database"
	"openeuler.org/PilotGo/gala-ops-plugin/httphandler"
)

const Version = "0.0.1"

var PluginInfo = &client.PluginInfo{
	Name:        "gala-ops",
	Version:     Version,
	Description: "gala-ops智能运维工具",
	Author:      "guozhengxin",
	Email:       "guozhengxin@kylinos.cn",
	Url:         "http://192.168.75.100:9999/plugin/gala-ops",
	// ReverseDest: "http://192.168.48.163:3000/",
}

func main() {
	fmt.Println("hello gala-ops")

	if err := database.MysqlInit(config.Config().Mysql); err != nil {
		fmt.Println("failed to initialize database")
		os.Exit(1)
	}

	InitLogger()

	router := gin.Default()

	GlobalClient := client.DefaultClient(PluginInfo)
	// 临时给server赋值
	GlobalClient.Server = "http://192.168.75.100:8888"
	GlobalClient.RegisterHandlers(router)
	InitRouter(router)

	// 临时自定义获取prometheus地址方式
	// promeplugin, err := getpromeplugininfo(GlobalClient.Server)
	// if err != nil {
	// 	logger.Error(err.Error())
	// 	os.Exit(1)
	// }
	// var PromeURL string = promeplugin["Url"].(string)

	// PromePlugin, err := client.GetClient().GetPluginInfo("prometheus")
	// if err != nil {
	// 	logger.Error("failed to get plugin info from pilotgoserver: ", err)
	// 	os.Exit(1)
	// }

	if err := router.Run(config.Config().Http.Addr); err != nil {
		logger.Fatal("failed to run server")
	}
}

func InitLogger() {
	if err := logger.Init(config.Config().Logopts); err != nil {
		fmt.Printf("logger init failed, please check the config file: %s", err)
		os.Exit(1)
	}
}

func InitRouter(router *gin.Engine) {
	api := router.Group("/plugin/gala-ops/api")
	{
		// 脚本执行结果接口
		// api.PUT("/run_script_result", httphandler.RunScriptResult)

		// 安装/升级/卸载gala-gopher监控终端
		api.PUT("/install_gopher", httphandler.InstallGopher)
		api.PUT("/upgrade_gopher", httphandler.UpgradeGopher)
		api.DELETE("/uninstall_gopher", httphandler.UninstallGopher)
	}

	metrics := router.Group("plugin/gala-ops/api/metrics")
	{
		metrics.GET("/targets_list", func(ctx *gin.Context) {
			httphandler.TargetsList(ctx)
		})
	}
}

func getpromeplugininfo(pilotgoserver string) (map[string]interface{}, error) {
	resp, err := http.Get(pilotgoserver + "/api/v1/plugins")
	if err != nil {
		logger.Error("faild to get plugin list: ", err)
	}
	bs, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	data := map[string]interface{}{
		"code": nil,
		"data": nil,
		"msg":  nil,
	}
	err = json.Unmarshal(bs, &data)
	if err != nil {
		logger.Error("unmarshal request plugin info error:%s", err.Error())
	}
	var PromePlugin map[string]interface{}
	for _, p := range data["data"].([]interface{}) {
		if p.(map[string]interface{})["name"] == "Prometheus" {
			PromePlugin = p.(map[string]interface{})
		}
	}
	if len(PromePlugin) == 0 {
		return nil, fmt.Errorf("pilotgo server not add prometheus plugin")
	}
	return PromePlugin, nil
}
