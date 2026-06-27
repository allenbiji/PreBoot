package cli

import (
	"github.com/allenbiji/clone-sage/internal/config"
	"github.com/allenbiji/clone-sage/internal/engine"
	"github.com/spf13/cobra"
)

func NewCheckCmd() *cobra.Command {
	var isQuickMode bool
	var cfgFile string

	checkCmd := &cobra.Command{
		Use:   "check",
		Short: "Check the sage.yml file for any errors",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.LoadFrom(cfgFile)
			if err != nil {
				return err
			}

			return engine.Run(cfg, isQuickMode)
		},
	}

	checkCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "Path to custom sage.yml")
	checkCmd.Flags().BoolVarP(&isQuickMode, "quick", "q", false, "Run only fast, low-cost checks")

	return checkCmd
}
