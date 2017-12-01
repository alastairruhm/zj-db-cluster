package cmd

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/alastairruhm/zj-db-cluster/config"
	"github.com/spf13/cobra"
)

// VERSION is app version
const VERSION = "0.2.0"

var (
	ClusterName string
	Config      *config.ClusterConfig
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

func errOutputExit(cmd *cobra.Command, msg interface{}) {
	errOutput(cmd, msg)
	os.Exit(-1)
}

// CheckClusterNameArgs will return error if global args ClusterName is ""
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

// ReadConfigFile ...
func ReadConfigFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// LoadConfig ...
func LoadConfig() (err error) {
	var cfgFilePath string
	cfgFilePath, err = config.GetCfgFilePath(ClusterName)
	if err != nil {
		return
	}
	var data string
	data, err = ReadConfigFile(cfgFilePath)
	if err != nil {
		return
	}
	err = config.ParseConfig(data)
	if err != nil {
		return
	}
	return
}
