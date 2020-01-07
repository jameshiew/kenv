package cmd

import (
	"os"

	"github.com/jameshiew/kenv/cmd/internal/utils"
	"github.com/jameshiew/kenv/internal"

	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "See or set the shell kubectl version",
	Long:  "See or set the shell kubectl version",
	Args:  cobra.MaximumNArgs(1), // setting the shell version must be done externally
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			cmd.PrintErrln("Cannot set shell version - has `eval $(kenv init -)` been run for this shell session?")
			os.Exit(1)
		}
		utils.PrintAndExit(cmd, internal.ShellVersion)
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)
}
