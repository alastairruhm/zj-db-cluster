package cmd

import (
	color "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

// VERSION shows app version
const VERSION = "0.0.1"

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("zijin database cluster cmd tool version %s\n", color.Cyan(VERSION))
	},
}
