package cmd

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/alastairruhm/zj-db-cluster/config"
	color "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var (
	ClusterName string
	Config      *config.Config
)

var RootCmd = &cobra.Command{
	Use:   "zj-db-cluster",
	Short: "zijin database cluster command line tool",
	Long:  ``,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if ClusterName == "" {
			errOutput(cmd, errors.New("name flags '-i' is required to specify the cluster"))

			cmd.Help()
			os.Exit(-1)

		}
		cfgFilePath, err := config.GetCfgFilePath(ClusterName)
		if err != nil {
			errOutput(cmd, err)
			os.Exit(-1)
		}
		data, err := readConfigFile(cfgFilePath)
		if err != nil {
			errOutput(cmd, err)
			os.Exit(-1)
		}
		Config, err = config.ParseConfigData(data)
		if err != nil {
			errOutput(cmd, err)
			os.Exit(-1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func errOutput(cmd *cobra.Command, msg interface{}) {
	cmd.Printf(color.Sprintf(color.Red("Error: %s\n"), color.Red(msg)))
}

func infoOutput(cmd *cobra.Command, msg interface{}) {
	cmd.Printf(color.Sprintf(color.Blue(msg)))
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
