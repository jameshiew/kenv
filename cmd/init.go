package cmd

import (
	"fmt"

	"github.com/jameshiew/kenv/cmd/internal/utils"
	"github.com/jameshiew/kenv/internal"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Set up the environment to use kenv",
	Long:    `This command initializes the kenv directory and runs "eval $(kenv init -)" in this kenv session, so that kenv can be used`,
	Example: "kenv init\n. <(kenv init -)",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return nil
		}
		if len(args) > 1 || args[0] != "-" {
			return fmt.Errorf("only '-' may be specified as an argument (to print out shell script)")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			utils.PrintAndExit(cmd, internal.ShellInitialization)
		}
		utils.RunAndExit(cmd, func() error {
			err := internal.Initialize()
			if err != nil {
				return err
			}
			fmt.Println(internal.InitializationHelpString())
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
