package cmd

import (
	"fmt"
	"os"

	"github.com/jameshiew/kenv/cmd/internal/utils"
	"github.com/jameshiew/kenv/internal"

	"github.com/spf13/cobra"
)

const (
	name     = "kenv"
	homepage = "https://github.com/jameshiew/kenv"
)

const versionTemplate = `{{with .Name}}{{printf "%s " .}}{{end}}{{printf "%s" .Version}}
`

var Version string

var rootCmd = &cobra.Command{
	Use:     name,
	Version: Version,
	Short:   "kenv manages your kubectl versions",
	Long: `kenv is a CLI tool for managing different versions of kubectl.
	Complete documentation is available at ` + homepage,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			fmt.Println(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var getRootCmd = &cobra.Command{
	Use:   "root",
	Short: "Get the root directory for kenv",
	Long:  "Get the root directory for kenv",
	Run: func(cmd *cobra.Command, args []string) {
		utils.PrintAndExit(cmd, internal.RootDir)
	},
}

func init() {
	rootCmd.SetVersionTemplate(versionTemplate)
	rootCmd.AddCommand(getRootCmd)
}
