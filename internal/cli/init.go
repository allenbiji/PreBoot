package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewInitCmd() *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Run init to initalize the project and generate sage-auto.yaml",
		Long:  "Run this command to scan your entire repository and from the inferences in your repo, a sage-auto.yaml file will be generated and which can also be extended via a sage.yaml file",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("init command has been invoked")
			return nil
		},
	}

	return initCmd
}
