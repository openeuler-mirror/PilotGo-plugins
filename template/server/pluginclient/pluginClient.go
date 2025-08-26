package pluginclient

import (
	"context"
	"os"

	"gitee.com/openeuler/PilotGo-plugin-template/conf"
	"gitee.com/openeuler/PilotGo/sdk/common"
	"gitee.com/openeuler/PilotGo/sdk/go-micro/registry"
	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
)

var Global_Client *client.Client

var Global_Context context.Context

func InitPluginClient() {
	sr, err := registry.NewServiceRegistrar(&registry.Options{
		Endpoints:   conf.Global_Config.Etcd.Endpoints,
		ServiceAddr: conf.Global_Config.Template.Addr,
		ServiceName: conf.Global_Config.Etcd.ServiveName,
		Version:     conf.Global_Config.Etcd.Version,
		MenuName:    conf.Global_Config.Etcd.MenuName,
		Icon:        conf.Global_Config.Etcd.Icon,
		DialTimeout: conf.Global_Config.Etcd.DialTimeout,
		Extentions:  GetExtentions(),
		Permissions: GetPermissions(),
	})
	if err != nil {
		logger.Error("failed to initialize registry: %s", err)
		os.Exit(-1)
	}

	client, err := client.NewClient(conf.Global_Config.Etcd.ServiveName, sr.Registry)
	if err != nil {
		logger.Error("failed to create plugin client: %s", err)
		os.Exit(-1)
	}
	Global_Client = client
	Global_Context = context.Background()
}

func GetExtentions() []common.Extention {
	var ex []common.Extention
	pe1 := &common.PageExtention{
		Type:       common.ExtentionPage,
		Name:       "主页面扩展",
		URL:        "/page",
		Permission: "plugin.template.page/menu",
	}
	me2 := &common.MachineExtention{
		Type:       common.ExtentionMachine,
		Name:       "机器扩展",
		URL:        "/host",
		Permission: "plugin.template/function",
	}
	be3 := &common.BatchExtention{
		Type:       common.ExtentionBatch,
		Name:       "批次扩展",
		URL:        "/batch",
		Permission: "plugin.template/function",
	}
	ex = append(ex, pe1, me2, be3)
	return ex
}

func GetPermissions() []common.Permission {
	var pe []common.Permission
	return pe
}
