/*
Copyright © 2023 agmmtoo
Copyrights apply to this source code.
Check LICENSE for details.
*/
package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate bash completion for your command",
	Long: `To load your completions run
	source <(pScan completion)
 To load completions automatically on login, add this line to you .bashrc file: $ ~/.bashrc
 source <(pScan completion)
 `,
	RunE: func(cmd *cobra.Command, args []string) error {
		return completionAction(os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}

func completionAction(out io.Writer) error {
	return rootCmd.GenBashCompletion(out)
}
