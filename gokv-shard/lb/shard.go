package lb

import "gokv/server"

type ShardStrategy int32

const (
	RANGE              ShardStrategy = 1
	CONSISTENT_HASHING ShardStrategy = 1
)

type Shard struct {
}

func NewShard() *Shard {

	// specify the shard strategy

	// create clusters
	clusterConf := &server.ClusterConf{}
	server.NewCluster(clusterConf)

	return &Shard{}
}

func sharding() {

}
