package Client

import "github.com/alastairruhm/zj-db-cluster/config"

type Checker interface {
	CheckDBConnection()
	CheckReplicaStatus()
	CheckReplicaConsistency()
}

type Cluster struct {
	DBUser      string
	DBPassword  string
	VIP         config.Vip
	AtlasProxys map[string]config.Atlas
}

// func NewCluster(c config.ClusterConfig) (*Cluster, error) {
// 	return
// }

// var _ Cluster = &Client{}
