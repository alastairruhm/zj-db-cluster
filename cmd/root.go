package cmd

import (
	"errors"
	"io/ioutil"

	"github.com/alastairruhm/zj-db-cluster/config"
	"github.com/spf13/cobra"
)

// VERSION is app version
const VERSION = "0.2.0"

var (
	ClusterName string
	Config      *config.Config
)

// RootCmd ...
var RootCmd = &cobra.Command{
	Use:   "zj-db-cluster",
	Short: "zijin database cluster command line tool",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// default command output usage string
		cmd.Usage()
	},
}

func errOutput(cmd *cobra.Command, msg interface{}) {
	cmd.Printf("Error: %s\n", msg)
}

// CheckClusterNameArgs ...
func CheckClusterNameArgs() error {
	if ClusterName == "" {
		return errors.New("name flags '-n' is required to specify the cluster")
	}
	return nil
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&ClusterName, "name", "n", "", "cluster name")
	RootCmd.AddCommand(ConfigCmd)
	RootCmd.AddCommand(CheckCmd)
}

func readConfigFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
