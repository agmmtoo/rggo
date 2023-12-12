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

	"agmmtoo.me/ammgo/cobra/pScan/scan"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List hosts in hosts list",
	RunE: func(cmd *cobra.Command, args []string) error {
		hostsFile, err := cmd.Flags().GetString("hosts-file")
		if err != nil {
			return err
		}
		return listAction(os.Stdout, hostsFile, args)
	},
}

func init() {
	hostsCmd.AddCommand(listCmd)
}

func listAction(out io.Writer, hostFile string, args []string) error {
	hl := &scan.HostsList{}
	if err := hl.Load(hostFile); err != nil {
		return err
	}
	for _, h := range hl.Hosts {
		if _, err := fmt.Fprintf(out, h); err != nil {
			return err
		}
	}
	return nil
}
