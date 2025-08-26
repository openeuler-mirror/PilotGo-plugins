/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin licensed under the Mulan Permissive Software License, Version 2.
 * See LICENSE file for more details.
 * Author: wubijie <wubijie@kylinos.cn>
 * Date: Wed Nov 15 15:28:38 2023 +0800
 */
package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"gitee.com/openeuler/PilotGo/sdk/go-micro/registry"
	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
	"openeuler.org/PilotGo/configmanage-plugin/config"
	"openeuler.org/PilotGo/configmanage-plugin/db"
	"openeuler.org/PilotGo/configmanage-plugin/global"
	"openeuler.org/PilotGo/configmanage-plugin/router"
	"openeuler.org/PilotGo/configmanage-plugin/service"
)

var config_file string

func main() {
	fmt.Println("hello plugin-config")

	// 初始化配置文件内容
	flag.StringVar(&config_file, "conf", "./config.yaml", "plugin-config configuration file")
	flag.Parse()
	err := config.Init(config_file)
	if err != nil {
		logger.Info("failed to load configure: %s", err)
		os.Exit(-1)
	}

	// 初始化日志信息
	if err := logger.Init(config.Config().Logopts); err != nil {
		logger.Info("logger init failed, please check the config file: %s", err)
		os.Exit(-1)
	}

	// mysql db初始化
	if err := db.MysqldbInit(config.Config().Mysql); err != nil {
		logger.Info("mysql db init failed, please check again: %s", err)
		os.Exit(-1)
	}

	// 初始化数据库表
	err = service.Init()
	if err != nil {
		logger.Info("init db table error: %s\n", err)
		os.Exit(-1)
	}

	sr, err := registry.NewServiceRegistrar(&registry.Options{
		Endpoints:   config.Config().Etcd.Endpoints,
		ServiceAddr: config.Config().HttpServer.Addr,
		ServiceName: config.Config().Etcd.ServiveName,
		Version:     config.Config().Etcd.Version,
		MenuName:    config.Config().Etcd.MenuName,
		Icon:        config.Config().Etcd.Icon,
		DialTimeout: config.Config().Etcd.DialTimeout,
		Extentions:  service.GetExtentions(),
		Permissions: service.GetPermission(),
	})
	if err != nil {
		logger.Error("failed to initialize registry: %s", err)
		os.Exit(-1)
	}

	client, err := client.NewClient(config.Config().Etcd.ServiveName, sr.Registry)
	if err != nil {
		logger.Error("failed to create plugin client: %s", err)
		os.Exit(-1)
	}

	global.GlobalClient = client
	service.GetTags()
	// 初始化路由信息
	server := router.InitRouter()
	go router.RegisterAPIs(server)
	if err := server.Run(config.Config().HttpServer.Addr); err != nil {
		logger.Info("failed to run server: %s", err)
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
	fmt.Println("Thanks to choose plugin-config!")
}
