package client

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/alastairruhm/zj-db-cluster/config"
)

// Checker ...
type Checker interface {
	CheckConnection() string
	CheckReplicaStatus()
	CheckReplicaConsistency()
}

// Cluster ...
type Cluster struct {
	Username   string
	Password   string
	VIPNode    config.Vip
	AtlasNodes map[string]config.Atlas
	DBNodes    map[string]config.Database
}

// NewCluster create a new instance with given config
func NewCluster(c config.ClusterConfig) *Cluster {
	cluster := &Cluster{
		Username:   c.Username,
		Password:   c.Password,
		VIPNode:    c.Vip,
		AtlasNodes: c.Atlas,
		DBNodes:    c.DB,
	}

	return cluster
}

// CheckConnection check all datasource connection
func (c *Cluster) CheckConnection() string {
	chVip := make(chan string)
	chAtlas := make(chan string)
	chDB := make(chan string)
	go func(vip config.Vip) {
		result := fmt.Sprintf("check VIP %s connection: ", vip.IP)
		result = result + CheckDBConnection(c.Username, c.Password, vip.IP, strconv.Itoa(vip.Port)) + "\n"
		chVip <- result
	}(c.VIPNode)
	close(chVip)

	for _, atlas := range c.AtlasNodes {
		go func(a config.Atlas) {
			result := fmt.Sprintf("check Atlas %s connection: ", a.IP)
			result = result + CheckDBConnection(c.Username, c.Password, a.IP, strconv.Itoa(a.Port)) + "\n"
			chAtlas <- result
		}(atlas)
	}
	close(chAtlas)

	for _, database := range c.DBNodes {
		go func(d config.Database) {
			result := fmt.Sprintf("check Database %s connection: ", d.IP)
			result = result + CheckDBConnection(c.Username, c.Password, d.IP, strconv.Itoa(d.Port)) + "\n"
			chDB <- result
		}(database)
	}
	close(chDB)

	resVip := ""
	resAtlas := ""
	resDB := ""

	for {
		select {
		case r, ok := <-chVip:
			if !ok {
				chVip = nil
			}
			resVip = r

		case r, ok := <-chAtlas:
			if !ok {
				chAtlas = nil
			}
			resAtlas += r

		case r, ok := <-chDB:
			if !ok {
				chDB = nil
			}
			resDB = resDB + r
		}
	}

	return resVip + resAtlas + resDB
}

func (c *Cluster) CheckReplicaStatus() {

}

func (c *Cluster) CheckReplicaConsistency() {

}

var _ Checker = &Cluster{}

// CheckDBConnection ...
func CheckDBConnection(user string, pass string, host string, port string) string {
	db, err := NewDBConn(user, pass, host, port)
	defer db.Close()
	if err != nil {
		return err.Error()
	}
	err = db.Ping()
	if err != nil {
		return err.Error()
	}
	return "OK"
}

// NewDBConn return a new database connection
func NewDBConn(user string, pass string, host string, port string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?timeout=3s", user, pass, host, port)
	db, err := sql.Open("mysql", dsn)
	return db, err
}
