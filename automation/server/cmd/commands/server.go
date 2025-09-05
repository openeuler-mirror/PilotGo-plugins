package commands

import (
	"github.com/spf13/cobra"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/cmd/config/options"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/service"
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

	manager := service.NewServiceManager(
		&service.LoggerService{Conf: opt.Config.Logopts},
		&service.MySQLService{Conf: opt.Config.Mysql},
		&service.RedisService{Conf: opt.Config.Redis},
		&service.EtcdService{Conf: opt.Config.Etcd, ServerConf: opt.Config.HttpServer},
	)
	if err := manager.InitAll(); err != nil {
		return err
	}
	defer manager.CloseAll()

	return nil
}
