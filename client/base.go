package client

import (
	"database/sql"
	"fmt"
	"strconv"
	"sync"

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

// NewChecker create a new instance with given config
var NewChecker = func(c config.ClusterConfig) Checker {
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

	var checkGroup sync.WaitGroup
	go func(vip config.Vip) {
		checkGroup.Add(1)
		result := fmt.Sprintf("check VIP %s connection: ", vip.IP)
		result = result + CheckDBConnection(c.Username, c.Password, vip.IP, strconv.Itoa(vip.Port)) + "\n"
		// fmt.Println("check vip ", vip.IP)
		chVip <- result
		close(chVip)
		checkGroup.Done()
	}(c.VIPNode)

	go func() {
		checkGroup.Add(1)
		var atlasGroup sync.WaitGroup
		for _, atlas := range c.AtlasNodes {
			atlasGroup.Add(1)
			go func(a config.Atlas) {
				result := fmt.Sprintf("check Atlas %s connection: ", a.IP)
				result = result + CheckDBConnection(c.Username, c.Password, a.IP, strconv.Itoa(a.Port)) + "\n"
				// fmt.Println("check atlas ", a.IP)
				chAtlas <- result
				atlasGroup.Done()
			}(atlas)
		}
		atlasGroup.Wait()
		close(chAtlas)
		checkGroup.Done()
	}()

	go func() {
		checkGroup.Add(1)
		var dbGroup sync.WaitGroup
		for _, database := range c.DBNodes {
			dbGroup.Add(1)
			go func(d config.Database) {
				result := fmt.Sprintf("check Database %s connection: ", d.IP)
				result = result + CheckDBConnection(c.Username, c.Password, d.IP, strconv.Itoa(d.Port)) + "\n"
				// fmt.Println("check db ", d.IP)
				chDB <- result
				dbGroup.Done()
			}(database)
		}
		dbGroup.Wait()
		close(chDB)
		checkGroup.Done()
	}()

	resVip := ""
	resAtlas := ""
	resDB := ""

	for {
		select {
		case r, ok := <-chVip:
			if !ok {
				chVip = nil
			} else {
				resVip = r
				// fmt.Println("chvip")
			}
		case r, ok := <-chAtlas:
			if !ok {
				chAtlas = nil
			} else {
				resAtlas += r
				// fmt.Println("atlas")
			}

		case r, ok := <-chDB:
			if !ok {
				chDB = nil
			} else {
				resDB = resDB + r
				// fmt.Println("chdb")
			}
		}
		// fmt.Println(chVip, chAtlas, chDB)
		if chVip == nil && chAtlas == nil && chDB == nil {
			break
		}
	}
	checkGroup.Wait()

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
