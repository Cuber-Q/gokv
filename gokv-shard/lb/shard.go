package lb

import (
	"gokv/gokv-shard/strategy"
	"gokv/server"
)

// A shard is a [replicate-group] of a shard of whole data range.
// Contains a raft cluster, which shares a set of same data;
// And a [route-table], so that the dispatcher can dispatch the
// read/write request to the target cluster.
type Shard struct {
	shardStrategy strategy.ShardStrategy
	routeTable    *RouteTable

	// replica group
	rg *server.Cluster
}

// specify which data can mapping to current shard
type RouteTable struct {
	routes []*Route
}

func (self *RouteTable) findRoute(keySharding uint64) *Route {
	for _, r := range self.routes {
		if r.match(keySharding) {
			return r
		}
	}
	return nil
}

type Route struct {
	start        uint64
	includeStart bool
	end          uint64
	includeEnd   bool
}

func (self *Route) match(keySharding uint64) bool {
	if self.includeStart && keySharding == self.start {
		return true
	}
	if self.includeEnd && keySharding == self.end {
		return true
	}
	return keySharding > self.start && keySharding < self.end
}

// To construct a shard, there are some steps:
// 1. build up a replica-group (or the raft cluster) with config;
// 2. specify the route table
// 3. start the replica-group
func NewShard() *Shard {

	// specify the shard strategy

	// create clusters
	clusterConf := &server.ClusterConf{}
	server.NewCluster(clusterConf)

	return &Shard{}
}

func sharding() {

}
