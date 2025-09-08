package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/cmd/config/options"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/global"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/router"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/service"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/pkg/utils"
)

func NewServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Example: `
		# run  the automation-service
		automation 
		or
		automation start
		`,
		Use:   cliName,
		Short: "Start the automation",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run()
		},
	}
	cmd.ResetFlags()
	return cmd
}
func Run() error {
	opt, err := options.NewOptions().TryLoadFromDisk()
	if err != nil {
		return err
	}

	ips, err := utils.GetAllHostIPs()
	if err != nil {
		return fmt.Errorf("获取本机ip地址失败: %v", err)
	} else {
		global.App.HttpAddr = ips[0].IP + opt.Config.HttpServer.Addr[strings.LastIndex(opt.Config.HttpServer.Addr, ":"):]
	}

	manager := service.NewServiceManager(
		&service.LoggerService{Conf: opt.Config.Logopts},
		&service.MySQLService{Conf: opt.Config.Mysql},
		&service.RedisService{Conf: opt.Config.Redis},
		&service.EtcdService{Conf: opt.Config.Etcd, ServerConf: global.App.HttpAddr},
	)
	if err := manager.InitAll(); err != nil {
		return err
	}
	defer manager.CloseAll()

	if err := router.HttpServerInit().Run(global.App.HttpAddr); err != nil {
		return err
	}
	return nil
}
