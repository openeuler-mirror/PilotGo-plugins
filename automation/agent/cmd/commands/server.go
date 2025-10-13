package commands

import (
	"ant-agent/cmd/config/options"
	"ant-agent/pkg/global"
	"ant-agent/pkg/router"
	"ant-agent/pkg/utils"
	"fmt"
	"strings"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"github.com/spf13/cobra"
)

func NewServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Example: `
		# run  the ant-agent-service
		ant-agent
		or
		ant-agent start
		`,
		Use:   cliName,
		Short: "Start the ant-agent",
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

	if err := logger.Init(opt.Config.Logopts); err != nil {
		return err
	}

	ips, err := utils.GetAllHostIPs()
	if err != nil {
		return fmt.Errorf("获取本机ip地址失败: %v", err)
	} else {
		global.HttpAddr = ips[0].IP + opt.Config.HttpServer.Addr[strings.LastIndex(opt.Config.HttpServer.Addr, ":"):]
	}

	if err := router.HttpServerInit().Run(global.HttpAddr); err != nil {
		return err
	}
	return nil
}
