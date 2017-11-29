package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"-v"},
	Short:   "Print the version semantic number",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("zijin database cluster tool version %s\n", VERSION)
	},
}
