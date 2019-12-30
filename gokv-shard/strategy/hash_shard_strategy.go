package strategy

import (
	"hash/fnv"
)

type HashShardStrategy struct {
}

func (h *HashShardStrategy) KeySharding(key string) (uint32, error) {
	hs := fnv.New32()
	_, e := hs.Write([]byte(key))
	if e != nil {
		return 0, e
	}
	return hs.Sum32(), nil
}
