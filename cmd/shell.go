package cmd

import (
	"github.com/jameshiew/kenv/cmd/internal/utils"
	"github.com/jameshiew/kenv/internal"

	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "See or set the shell kubectl version",
	Long:  "See or set the shell kubectl version",
	Args:  cobra.NoArgs, // setting the shell version must be done externally
	Run: func(cmd *cobra.Command, args []string) {
		utils.PrintAndExit(cmd, internal.ShellVersion)
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)
}
