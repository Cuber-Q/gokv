package gokv_shard

import (
	"gokv/gokv-shard/lb"
	"gokv/gokv-shard/strategy"
)

// Dispatcher is the main module which called by clients
// and dispatches every request to specific gokv-servers with configured
// ShardingStrategy, aggregates results of every gokv-server and responses
// to clients. It's the main entrance of whole gokv system for clients
// when running in sharding.
type Dispatcher struct {
	shardStrategy strategy.ShardStrategy
	shards        []*lb.Shard
	routeTable    *lb.RouteTable
}

// todo return result
func (dp *Dispatcher) dispatch(key string) {
	// 1. get sharding value of key
	keySharding := dp.shardStrategy.KeySharding(key)

	// 2. get shard
	shard := dp.sharding(keySharding)

	// 3. execute request on shard
	shard.execute()
}

func (self *Dispatcher) sharding(u uint32) *lb.Shard {
	return self.shards[0]
}
