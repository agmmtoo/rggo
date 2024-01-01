/*
Copyright © 2024 agmmtoo
Copyrights apply to this source code.
Check LICENSE for details.
*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:          "complete <id>",
	Short:        "Mark an item as complete",
	SilenceUsage: true,
	Args:         cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiRoot, _ := cmd.Flags().GetString("api-root")
		return completeAction(os.Stdout, apiRoot, args[0])
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}

func completeAction(out io.Writer, apiRoot, arg string) error {
	id, err := strconv.Atoi(arg)
	if err != nil {
		return fmt.Errorf("%w: Item id must be a number", ErrNotNumber)
	}
	if err := completeItem(apiRoot, id); err != nil {
		return err
	}
	return printComplete(out, id)
}

func printComplete(out io.Writer, id int) error {
	_, err := fmt.Fprintf(out, "Item number %d marked as completed.\n", id)
	return err
}
