package cmd

import (
	"errors"
	"path/filepath"

	"os"

	"github.com/alastairruhm/zj-db-cluster/config"
	"github.com/alastairruhm/zj-db-cluster/utils"
	"github.com/spf13/cobra"
)

var template = `dbusername = 
dbpassword = 
[vip]
ip = 
port = 


[atlas]
  
  [atlas.a]
  ip = 
  port = 

  [atlas.b]
  ip = 
  port = 

[database]
  
  [database.master]
  ip = 
  port = 

  [database.slave_a]
  ip = 
  port = 

  [database.slave_b]
  ip = 
  port = 
`

// ConfigCmd sub-command of zcloud about server
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "operation about configuration file",
	Long: `configuration file operation.
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

// ConfigInitCmd will generate a configuration file
var ConfigInitCmd = &cobra.Command{
	Use:   "init",
	Short: "generate configuration file",
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := CheckClusterNameArgs(); err != nil {
			errOutput(cmd, err)
		}
	},
	Run: initConfig,
}

// ConfigTestCmd test target file
var ConfigTestCmd = &cobra.Command{
	Use:   "test",
	Short: "test configuration file",
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := CheckClusterNameArgs(); err != nil {
			errOutput(cmd, err)
		}
	},
	Run: testConfig,
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

	b, err := utils.Exist(appPath)

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

	b, err = utils.Exist(configFilePath)
	if err != nil {
		errOutput(cmd, err)
		os.Exit(-1)
	}
	if b {
		errOutput(cmd, errors.New("config file "+configFile+" already exist"))
		os.Exit(-1)
	}
	// 注意权限最好是 0600
	err = utils.WriteFile(appPath, configFile, template)
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
	ConfigCmd.AddCommand(ConfigListCmd)
}
