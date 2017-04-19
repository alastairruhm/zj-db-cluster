package cmd

import (
	"errors"
	"path/filepath"

	"os"

	"github.com/alastairruhm/zj-db-cluster/config"
	"github.com/spf13/cobra"
)

var template = `[vip]
ip = 
port = 
dbusername = 
dbpassword = 

[atlas]
  
  [atlas.a]
  ip = 
  port = 
  dbusername = 
  dbpassword = 

  [atlas.b]
  ip = 
  port = 
  dbusername = 
  dbpassword = 

[database]
  
  [database.master]
  ip = 
  port = 
  dbusername = 
  dbpassword =

  [database.slave_a]
  ip = 
  port = 
  dbusername = 
  dbpassword =

  [database.slave_b]
  ip = 
  port = 
  dbusername = 
  dbpassword = 
`

// ConfigCmd sub-command of zcloud about server
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "operation about configuration",
	Long: `initialize configuration file.
    `,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if ClusterName == "" {
			errExitOutput(cmd, errors.New("name flags is required to specify the cluster"))
		}
	},
}

var ConfigInitCmd = &cobra.Command{
	Use:   "init",
	Short: "initiate configuration file",
	Run:   initConfig,
}

var ConfigTestCmd = &cobra.Command{
	Use:   "test",
	Short: "test configuration file",
	Run:   testConfig,
}

func initConfig(cmd *cobra.Command, args []string) {
	appPath, err := config.GetAppPath()
	if err != nil {
		errExitOutput(cmd, err)
	}

	b, err := exist(appPath)

	if err != nil {
		errExitOutput(cmd, err)
	}

	if !b { // If path doesn't yet exist, create it
		err = os.Mkdir(appPath, 0700)
		if err != nil {
			errExitOutput(cmd, errors.New("create working directory failed"))
		}
	}

	configFile := ClusterName + ".toml"
	configFilePath := filepath.Join(appPath, configFile)

	b, err = exist(configFilePath)
	if err != nil {
		errExitOutput(cmd, err)
	}
	if b {
		errExitOutput(cmd, errors.New("config file "+configFile+" already exist"))
	}
	// 注意权限最好是 0600
	err = writeConfigToFile(appPath, configFile, template)
	if err != nil {
		errExitOutput(cmd, err)
	}
}

func testConfig(cmd *cobra.Command, args []string) {
	cmd.Printf("test OK\n")
}

func init() {
	ConfigCmd.AddCommand(ConfigInitCmd)
	ConfigCmd.AddCommand(ConfigTestCmd)
}
