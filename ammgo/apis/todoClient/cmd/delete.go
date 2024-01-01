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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:          "del <id>",
	Short:        "Deletes an item from the list",
	SilenceUsage: true,
	Args:         cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiRoot, _ := cmd.Flags().GetString("api-root")
		return delAction(os.Stdout, apiRoot, args[0])
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func delAction(out io.Writer, apiRoot, arg string) error {
	id, err := strconv.Atoi(arg)
	if err != nil {
		return fmt.Errorf("%w: Item id must be a number", ErrNotNumber)
	}
	if err := deleteItem(apiRoot, id); err != nil {
		return err
	}
	return printDel(out, id)
}

func printDel(out io.Writer, id int) error {
	_, err := fmt.Fprintf(out, "Item number %d deleted. \n", id)
	return err
}
