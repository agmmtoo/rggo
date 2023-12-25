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
	"strconv"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view <id>",
	Short: "View details about a single item",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiRoot, _ := cmd.Flags().GetString("api-root")
		return viewAction(os.Stdout, apiRoot, args[0])
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}

func viewAction(out io.Writer, apiRoot, arg string) error {
	id, err := strconv.Atoi(arg)
	if err != nil {
		return fmt.Errorf("%w: Item id must be a number", ErrNotNumber)
	}
	i, err := getOne(apiRoot, id)
	if err != nil {
		return err
	}
	return printOne(out, i)
}

func printOne(out io.Writer, i item) error {
	w := tabwriter.NewWriter(out, 14, 2, 0, ' ', 0)
	fmt.Fprintf(w, "Task:\t%s\n", i.Task)
	fmt.Fprintf(w, "Created at:\t%s\n", i.CreatedAt.Format(timeFormat))
	if i.Done {
		fmt.Fprintf(w, "Completed:\t%s\n", "Yes")
		fmt.Fprintf(w, "Completed At:\t%s\n", i.CompletedAt.Format(timeFormat))
		return w.Flush()
	}
	fmt.Fprintf(w, "Completed:\t%s\n", "No")
	return w.Flush()
}
