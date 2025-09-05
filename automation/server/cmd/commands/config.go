package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/cmd/config"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/pkg/utils"
)

func NewTemplateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Example: `
		# generate the default configuration file
		automation template
		`,
		Use:   "template",
		Short: "Generate the default configuration file",
		RunE: func(cmd *cobra.Command, args []string) error {
			return templateRun()
		},
	}
	cmd.ResetFlags()
	return cmd
}

func templateRun() error {
	operator := utils.NewYamlOpeartor(config.DefaultConfigTemplate,
		utils.WithCommentsTagFlag(utils.HeadComment),
		utils.WithDefaultTagName(utils.DefaultTagName))
	yamlContent, err := operator.Marshal()
	if err != nil {
		fmt.Printf("Marshal error: %v\n", err)
		return err
	}
	err = os.WriteFile("./automation.yaml", yamlContent, os.ModePerm)
	if err != nil {
		fmt.Printf("create templete error: %v\n", err)
		return err
	}

	fmt.Println("create automation.yaml file success, please search it in current directory.")
	return nil
}
