package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCheckCmd() *cobra.Command {
	var isQuickMode bool
	var cfgFile string

	checkCmd := &cobra.Command{
		Use:   "check",
		Short: "Check the sage.yml file for any errors",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Check command invoked with configs, %s and %v", cfgFile, isQuickMode)
			return nil
		},
	}

	checkCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "Path to custom sage.yml")
	checkCmd.Flags().BoolVarP(&isQuickMode, "quick", "q", false, "Run only fast, low-cost checks")

	return checkCmd
}
