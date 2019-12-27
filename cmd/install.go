package cmd

import (
	"runtime"

	"github.com/jameshiew/kenv/cmd/internal/utils"
	"github.com/jameshiew/kenv/internal"

	"github.com/spf13/cobra"
)

var (
	_OS          string
	architecture string
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a specific kubectl version",
	Long:  `Install a specific kubectl version`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		utils.RunAndExit(cmd, func() error {
			version := args[0]
			return internal.Install(version, _OS, architecture)
		})
	},
}

func init() {
	installCmd.Flags().StringVarP(&_OS, "os", "o", runtime.GOOS, "OS for which to install kubectl")
	installCmd.Flags().StringVarP(&architecture, "architecture", "a", runtime.GOARCH, "Architecture for which to install kubectl")
	rootCmd.AddCommand(installCmd)
}
