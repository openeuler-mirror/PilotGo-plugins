package service

import (
	"fmt"

	"gitee.com/openeuler/PilotGo/sdk/common"
	"gitee.com/openeuler/PilotGo/sdk/go-micro/registry"
	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/cmd/config/options"
)

type EtcdService struct {
	ServerConf *options.HttpServer
	Conf       *options.Etcd
}

func (e *EtcdService) Name() string {
	return "Etcd"
}

func (e *EtcdService) Init(ctx *AppContext) error {
	sr, err := registry.NewServiceRegistrar(&registry.Options{
		Endpoints:   e.Conf.Endpoints,
		ServiceAddr: e.ServerConf.Addr,
		ServiceName: e.Conf.ServiceName,
		Version:     e.Conf.Version,
		MenuName:    e.Conf.MenuName,
		Icon:        e.Conf.Icon,
		DialTimeout: e.Conf.DialTimeout,
		Extentions:  getExtentions(),
		Permissions: getPermissions(),
	})
	if err != nil {
		return err
	}

	client, err := client.NewClient(e.Conf.ServiceName, sr.Registry)
	if err != nil {
		return fmt.Errorf("failed to create plugin client: %s", err)
	}
	ctx.Client = client
	return nil
}

func (e *EtcdService) Close() error {
	return nil
}

func getExtentions() []common.Extention {
	var ex []common.Extention
	pe1 := &common.PageExtention{
		Type:       common.ExtentionPage,
		Name:       "高危语句规则",
		URL:        "/dangerous_rule",
		Permission: "plugin.automation.page/menu",
	}
	pe2 := &common.PageExtention{
		Type:       common.ExtentionPage,
		Name:       "脚本库",
		URL:        "/script_library",
		Permission: "plugin.automation.page/menu",
	}
	pe3 := &common.PageExtention{
		Type:       common.ExtentionPage,
		Name:       "自定义脚本",
		URL:        "/custom_scripts",
		Permission: "plugin.automation.page/menu",
	}
	ex = append(ex, pe1, pe2, pe3)
	return ex
}
func getPermissions() []common.Permission {
	var pe []common.Permission
	return pe
}
