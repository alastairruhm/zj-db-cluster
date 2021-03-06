package cmd

import (
	"github.com/alastairruhm/zj-db-cluster/client"
	"github.com/alastairruhm/zj-db-cluster/config"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

// CheckCmd do check the status of db cluster
var CheckCmd = &cobra.Command{
	Use:   "check",
	Short: "check status of database cluster",
	Long: `initialize configuration file.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

// CheckConnectionCmd ...
var CheckConnectionCmd = &cobra.Command{
	Use:   "connection",
	Short: "check connection status of all nodes",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if err := CheckClusterNameArgs(); err != nil {
			errOutput(cmd, err)
		}

		if err := LoadConfig(); err != nil {
			errOutputExit(cmd, err)
		}
	},
	Run: CheckConnection,
}

// CheckConnection ...
func CheckConnection(cmd *cobra.Command, args []string) {
	c := client.NewChecker(config.Config)
	result := c.CheckConnection()
	cmd.Printf(result)
}

// func check(cmd *cobra.Command, args []string) {
// 	switch Item {
// 	case "connection":
// 		c := client.NewCluster(config.Config)
// 		result := c.CheckConnection()
// 		cmd.Printf(result)

// 	case "replica-status":
// 		// for k, v := range Config.DB {
// 		// 	if k != "master" {
// 		// 		go CheckReplicaStatus(cmd, &v)
// 		// 	}
// 		// }
// 	case "replica-consistency":

// 	default:
// 		if err := cmd.Usage(); err != nil {
// 			errOutputExit(cmd, err.Error())
// 		}
// 	}

// }

// CheckReplicaStatus check mysql replication status
// The process will check slave IO status and slave running status
// func CheckReplicaStatus(cmd *cobra.Command, db *config.Database) {
// 	cmd.Printf("check slave %s replication status: ", db.IP)
// 	status, err := CheckSlaveReplicaStatus(db.Dbusername, db.Dbpassword, db.IP, strconv.Itoa(db.Port))
// 	if err != nil {
// 		errOutput(cmd, err)
// 	} else {
// 		if status.SlaveIORunning != "Yes" || status.SlaveSQLRunning != "Yes" {
// 			errOutput(cmd, "error")
// 		} else {
// 			cmd.Printf("ok\n")
// 		}
// 	}
// }

// func CheckNodeConn(cmd *cobra.Command, node interface{}) {
// 	switch node.(type) {
// 	case config.Vip:
// 		v, ok := node.(config.Vip)
// 		cmd.Printf("check VIP %s connection: ", v.IP)
// 		if !ok {
// 			errOutput(cmd, "type assertion is illegal")
// 			os.Exit(-1)
// 		}

// 		err := checkDBConnection(v.Dbusername, v.Dbpassword, v.IP, strconv.Itoa(v.Port))
// 		if err != nil {
// 			errOutput(cmd, err)
// 		} else {
// 			cmd.Printf("OK\n")
// 		}
// 	case config.Atlas:
// 		v, ok := node.(config.Atlas)
// 		cmd.Printf("check Atlas %s connection: ", v.IP)
// 		if !ok {
// 			errOutput(cmd, "type assertion is illegal")
// 			os.Exit(-1)
// 		}
// 		err := checkDBConnection(v.Dbusername, v.Dbpassword, v.IP, strconv.Itoa(v.Port))
// 		if err != nil {
// 			errOutput(cmd, err)
// 		} else {
// 			cmd.Printf("OK\n")
// 		}
// 	case config.Database:
// 		v, ok := node.(config.Database)
// 		cmd.Printf("check Database %s connection: ", v.IP)
// 		if !ok {
// 			errOutput(cmd, "type assertion is illegal")
// 			os.Exit(-1)
// 		}
// 		err := checkDBConnection(v.Dbusername, v.Dbpassword, v.IP, strconv.Itoa(v.Port))
// 		if err != nil {
// 			errOutput(cmd, err)
// 		} else {
// 			cmd.Printf("OK\n")
// 		}
// 	}
// }

func init() {
	CheckCmd.AddCommand(CheckConnectionCmd)
}
