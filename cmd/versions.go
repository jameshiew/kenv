package cmd

import (
	"github.com/jameshiew/kenv/cmd/internal/utils"
	"github.com/jameshiew/kenv/internal"
	"strings"

	"github.com/spf13/cobra"
)

var versionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "See all installed versions",
	Long:  `See all installed versions`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		utils.PrintAndExit(cmd, func() (string, error) {
			versions, err := internal.InstalledVersions()
			if err != nil {
				return "", err
			}
			var sb strings.Builder
			for _, v := range versions {
				sb.WriteString(v + "\n")
			}
			return strings.TrimSpace(sb.String()), nil
		})
	},
}

func init() {
	rootCmd.AddCommand(versionsCmd)
}
