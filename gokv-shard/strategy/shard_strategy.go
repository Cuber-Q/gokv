package strategy

type ShardStrategy interface {
	// get sharding value of key
	KeySharding(key string) uint32
}
