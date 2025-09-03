package commands

import (
	"fmt"

	"github.com/spf13/cobra"
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
	fmt.Println("jinlaile")
	return nil
}
