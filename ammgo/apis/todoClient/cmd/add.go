/*
Copyright © 2023 agmmtoo
Copyrights apply to this source code.
Check LICENSE for details.
*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <task>",
	Short: "Add a new task to the list.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiRoot, _ := cmd.Flags().GetString("api-root")
		return addAction(os.Stdout, apiRoot, args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addAction(out io.Writer, apiRoot string, args []string) error {
	task := strings.Join(args, " ")
	if err := addItem(apiRoot, task); err != nil {
		return err
	}
	return printAdd(out, task)
}

func printAdd(out io.Writer, task string) error {
	_, err := fmt.Fprintf(out, "Added task %q to the list.\n", task)
	return err
}
