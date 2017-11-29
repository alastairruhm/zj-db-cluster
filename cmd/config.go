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
			errOutput(cmd, errors.New("name flags is required to specify the cluster"))
			os.Exit(-1)
		}
	},
}

// ConfigInitCmd will generate a configuration file
var ConfigInitCmd = &cobra.Command{
	Use:   "init",
	Short: "generate configuration file",
	Run:   initConfig,
}

// ConfigTestCmd test target file
var ConfigTestCmd = &cobra.Command{
	Use:   "test",
	Short: "test configuration file",
	Run:   testConfig,
}

// ConfigListCmd will list the configuration files in pre-defined path
var ConfigListCmd = &cobra.Command{
	Use:   "list",
	Short: "list configuration file",
	Run:   listConfig,
}

func initConfig(cmd *cobra.Command, args []string) {
	appPath, err := config.GetAppPath()
	if err != nil {
		errOutput(cmd, err)
		os.Exit(-1)
	}

	b, err := exist(appPath)

	if err != nil {
		errOutput(cmd, err)
		os.Exit(-1)
	}

	if !b { // If path doesn't yet exist, create it
		err = os.Mkdir(appPath, 0700)
		if err != nil {
			errOutput(cmd, errors.New("create working directory failed"))
			os.Exit(-1)
		}
	}

	configFile := ClusterName + ".toml"
	configFilePath := filepath.Join(appPath, configFile)

	b, err = exist(configFilePath)
	if err != nil {
		errOutput(cmd, err)
		os.Exit(-1)
	}
	if b {
		errOutput(cmd, errors.New("config file "+configFile+" already exist"))
		os.Exit(-1)
	}
	// 注意权限最好是 0600
	err = writeConfigToFile(appPath, configFile, template)
	if err != nil {
		errOutput(cmd, err)
		os.Exit(-1)
	}
}

func testConfig(cmd *cobra.Command, args []string) {
	cmd.Printf("test OK\n")
}

func listConfig(cmd *cobra.Command, args []string) {
	cmd.Printf("test OK\n")
}

func init() {
	ConfigCmd.AddCommand(ConfigInitCmd)
	ConfigCmd.AddCommand(ConfigTestCmd)
}
