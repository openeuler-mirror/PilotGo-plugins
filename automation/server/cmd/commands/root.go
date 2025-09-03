package commands

import (
	"os"

	"github.com/spf13/cobra"
)

const cliName = "start"

var rootCmd = &cobra.Command{
	Use:   "automation",
	Short: "automation CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
	DisableAutoGenTag: true,
	SilenceUsage:      true,
}

func Execute() {
	if len(os.Args) == 1 {
		rootCmd.SetArgs([]string{cliName})
	}
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(NewServerCommand())
}
