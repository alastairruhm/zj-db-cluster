package cmd

import (
	"database/sql"
	"os"
	"path/filepath"

	"fmt"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ReplicaStatus ...
type ReplicaStatus struct {
	SlaveIORunning  string
	SlaveSQLRunning string
}

// Check if a file or directory exists.
func exist(_path string) (bool, error) {
	_, err := os.Stat(_path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func writeConfigToFile(_path string, file string, text string) error {
	f, err := os.Create(filepath.Join(_path, file))
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(text)
	return err
}

func checkDBConnection(dbuser string, dbpasswd string, host string, port string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?timeout=5s", dbuser, dbpasswd, host, port)
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		return err
	}
	err = db.Ping()
	return err
}

// CheckSlaveReplicaStatus ...
func CheckSlaveReplicaStatus(dbuser string, dbpasswd string, host string, port string) (ReplicaStatus, error) {
	var status ReplicaStatus
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?timeout=5s", dbuser, dbpasswd, host, port)
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		return status, err
	}

	rows, err := db.Query("show slave status")
	if err != nil {
		return status, err
	}
	cols, _ := rows.Columns()
	buff := make([]interface{}, len(cols)) // temp slice
	data := make([]string, len(cols))      // data storing slice
	for i := range buff {
		buff[i] = &data[i]
	}
	for rows.Next() {
		rows.Scan(buff...)
	}
	for k, col := range data {
		if cols[k] == "Slave_IO_Running" {
			status.SlaveIORunning = col
		}
		if cols[k] == "Slave_SQL_Running" {
			status.SlaveSQLRunning = col
		}
	}
	return status, nil
}
