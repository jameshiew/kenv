package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Prints bash completion to stdout",
	Long: `To load completion into your running shell, run

. <(kenv completion)

To configure your bash shell to load completions for each session add to your bashrc

# ~/.bashrc or ~/.profile
. <(kenv completion)
`,
	Run: func(cmd *cobra.Command, args []string) {
		err := rootCmd.GenBashCompletion(os.Stdout)
		if err != nil {
			cmd.PrintErrf("ERROR: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
