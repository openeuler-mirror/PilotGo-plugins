package commands

import (
	"ant-agent/cmd/config"
	"ant-agent/pkg/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
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
	err = os.WriteFile("./ant-agent.yaml", yamlContent, os.ModePerm)
	if err != nil {
		fmt.Printf("create templete error: %v\n", err)
		return err
	}

	fmt.Println("create ant-agent.yaml file success, please search it in current directory.")
	return nil
}
