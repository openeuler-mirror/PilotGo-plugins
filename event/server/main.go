/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugins licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Tue Jun 4 15:19:07 2024 +0800
 */
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
	plugin_manage "openeuler.org/PilotGo/PilotGo-plugin-event/client"
	"openeuler.org/PilotGo/PilotGo-plugin-event/config"
	"openeuler.org/PilotGo/PilotGo-plugin-event/db"
	"openeuler.org/PilotGo/PilotGo-plugin-event/router"
	"openeuler.org/PilotGo/PilotGo-plugin-event/service"
)

func main() {
	if err := config.Init(); err != nil {
		fmt.Println("failed to load configure, exit..", err)
		os.Exit(-1)
	}

	if err := logger.Init(config.Config().Logopts); err != nil {
		fmt.Printf("logger init failed, please check the config file: %s", err)
		os.Exit(-1)
	}

	/*
		1. 连接influxdb2
		2. 创建插件客户端
		3. 添加页面拓展点
		4. eventbus监听
	*/
	db.InfluxdbInit(config.Config().Influxd)
	plugin_manage.EventClient = client.DefaultClient(plugin_manage.Init(config.Config().PluginEvent))
	service.AddExtentions()
	service.EventBusInit()

	if err := router.HttpServerInit(config.Config().HttpServer); err != nil {
		logger.Error("http server init failed, error:%v", err)
		os.Exit(-1)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for {
		s := <-c
		switch s {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			logger.Info("signal interrupted: %s", s.String())
			goto EXIT
		default:
			logger.Info("unknown signal: %s", s.String())
		}
	}

EXIT:
	logger.Info("exit system, bye~")
}
