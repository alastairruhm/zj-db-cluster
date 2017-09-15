package cmd

import (
	"strconv"

	"os"

	"github.com/alastairruhm/zj-db-cluster/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

var (
	// Item hold the flag for item that will be checked
	Item string
)

// CheckCmd do check the status of db cluster
var CheckCmd = &cobra.Command{
	Use:   "check",
	Short: "check status of database cluster",
	Long: `initialize configuration file.
    `,
	Run: check,
}

func check(cmd *cobra.Command, args []string) {
	switch Item {
	case "connection":
		go CheckNodeConn(cmd, Config.Vip)
		for _, v := range Config.Atlas {
			go CheckNodeConn(cmd, v)
		}
		for _, v := range Config.DB {
			go CheckNodeConn(cmd, v)
		}
	case "replica-status":
		for k, v := range Config.DB {
			if k != "master" {
				go CheckReplicaStatus(cmd, &v)
			}
		}
	case "replica-consistency":

	default:
	}

}

func CheckReplicaStatus(cmd *cobra.Command, db *config.Database) {
	cmd.Printf("check slave %s replication status: ", db.IP)
	status, err := CheckSlaveReplicaStatus(db.Dbusername, db.Dbpassword, db.IP, strconv.Itoa(db.Port))
	if err != nil {
		errOutput(cmd, err)
	} else {
		if status.SlaveIORunning != "Yes" || status.SlaveSQLRunning != "Yes" {
			errOutput(cmd, "error")
		} else {
			cmd.Printf("ok\n")
		}
	}
}

func CheckNodeConn(cmd *cobra.Command, node interface{}) {
	switch node.(type) {
	case config.Vip:
		v, ok := node.(config.Vip)
		cmd.Printf("check VIP %s connection: ", v.IP)
		if !ok {
			errOutput(cmd, "type assertion is illegal")
			os.Exit(-1)
		}

		err := checkDBConnection(v.Dbusername, v.Dbpassword, v.IP, strconv.Itoa(v.Port))
		if err != nil {
			errOutput(cmd, err)
		} else {
			cmd.Printf("OK\n")
		}
	case config.Atlas:
		v, ok := node.(config.Atlas)
		cmd.Printf("check Atlas %s connection: ", v.IP)
		if !ok {
			errOutput(cmd, "type assertion is illegal")
			os.Exit(-1)
		}
		err := checkDBConnection(v.Dbusername, v.Dbpassword, v.IP, strconv.Itoa(v.Port))
		if err != nil {
			errOutput(cmd, err)
		} else {
			cmd.Printf("OK\n")
		}
	case config.Database:
		v, ok := node.(config.Database)
		cmd.Printf("check Database %s connection: ", v.IP)
		if !ok {
			errOutput(cmd, "type assertion is illegal")
			os.Exit(-1)
		}
		err := checkDBConnection(v.Dbusername, v.Dbpassword, v.IP, strconv.Itoa(v.Port))
		if err != nil {
			errOutput(cmd, err)
		} else {
			cmd.Printf("OK\n")
		}
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&Item, "item", "i", "", "check item: connection, replica-status, replica-consistency")
}
