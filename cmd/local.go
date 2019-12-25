package cmd

import (
	"github.com/jameshiew/kenv/cmd/internal/utils"
	"github.com/jameshiew/kenv/internal"

	"github.com/spf13/cobra"
)

var localCmd = &cobra.Command{
	Use:     "local <version>",
	Short:   "See or set the kubectl version specific to this directory",
	Long:    `See or set the kubectl version specific to this directory`,
	Example: "kenv local 1.17.0",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			utils.PrintAndExit(cmd, internal.LocalVersion)
		}
		version := args[0]
		utils.RunAndExit(cmd, func() error {
			return internal.SetLocalVersion(version)
		})
	},
}

func init() {
	rootCmd.AddCommand(localCmd)
}
