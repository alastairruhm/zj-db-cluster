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
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if ClusterName == "" {
			errExitOutput(cmd, errors.New("name flags is required to specify the cluster"))
		}
		cfgFilePath, err := config.GetCfgFilePath(ClusterName)
		if err != nil {
			errExitOutput(cmd, err)
		}
		data, err := readConfigFile(cfgFilePath)
		if err != nil {
			errExitOutput(cmd, err)
		}
		Config, err = config.ParseConfigData(data)
		if err != nil {
			errExitOutput(cmd, err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func errExitOutput(cmd *cobra.Command, msg interface{}) {
	cmd.Printf(color.Sprintf(color.Red("Error: %s"), color.Red(msg)))
	os.Exit(-1)
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
