package cmd

import (
	"github.com/jameshiew/kenv/cmd/internal/utils"
	"github.com/jameshiew/kenv/internal"

	"github.com/spf13/cobra"
)

var prefixCmd = &cobra.Command{
	Use:   "prefix",
	Short: "Get the kubectl prefix",
	Long:  `Get the path to the current kubectl binary or for a specific version`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			utils.PrintAndExit(cmd, internal.CurrentPrefixDir)
		}
		version := args[0]
		utils.PrintAndExit(cmd, func() (s string, err error) {
			return internal.PrefixDir(version)
		})
	},
}

func init() {
	rootCmd.AddCommand(prefixCmd)
}
