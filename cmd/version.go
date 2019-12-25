package cmd

import (
	"github.com/jameshiew/kenv/cmd/internal/utils"
	"github.com/jameshiew/kenv/internal"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get the current kubectl version in use",
	Long:  `Get the current kubectl version in use`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		utils.PrintAndExit(cmd, internal.CurrentVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
