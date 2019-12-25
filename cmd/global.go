package cmd

import (
	"github.com/jameshiew/kenv/cmd/internal/utils"
	"github.com/jameshiew/kenv/internal"

	"github.com/spf13/cobra"
)

var globalCmd = &cobra.Command{
	Use:   "global <version>",
	Short: "See or set the default kubectl version to be used globally",
	Long: `See or set the global kubectl version. 

By default, the global version is "system", which is the first version of kubectl on your path not managed by kenv.`,
	Example: "kenv global 1.17.0",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			utils.PrintAndExit(cmd, internal.GlobalVersion)
		}
		version := args[0]
		utils.RunAndExit(cmd, func() error {
			return internal.SetGlobalVersion(version)
		})
	},
}

func init() {
	rootCmd.AddCommand(globalCmd)
}
