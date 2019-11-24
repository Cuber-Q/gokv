package server

import (
	"encoding/json"
	"gokv/core"
	"log"
	"time"
)

// An StoreOperation implement that is able to sync data with raftNode
type RaftOperation struct {
	// raftNode to sync data
	raftNode *raftNode

	// local storage, which is only kept by current node
	storage *core.Storage
}

func newRaftOperation(raft *raftNode, storage *core.Storage) *RaftOperation {
	return &RaftOperation{
		raftNode: raft,
		storage:  storage,
	}
}

// when Set k-v, this implementation will sync data to all it's
// followers if current node is Leader.
// And then save the k-v in current node's local storage when
// followers return ACK.
func (o *RaftOperation) Set(k, v string) {
	if !o.raftNode.isLeader {
		log.Printf("not Leader, abort SET sop. current node:[%v] k=%v, v=%v", o.raftNode.raft.String(), k, v)
	}
	event := logEntryData{K: k, V: v}
	eventBytes, err := json.Marshal(event)
	if err != nil {
		log.Printf("json.Marshal failed, err:%v", err)
		return
	}

	applyFuture := o.raftNode.raft.Apply(eventBytes, 5*time.Second)
	if err := applyFuture.Error(); err != nil {
		log.Printf("raftNode.Apply failed:%v", err)
		return
	}
	log.Printf("raftNode.Apply OK")
}

// get v from local storage. If the k is not set by current node's Set method
// or by raftNode sync, it should return empty string
func (o *RaftOperation) Get(k string) string {
	return o.storage.GetV2(k)
}

func (o *RaftOperation) Exist(k string) bool {
	return false
}

func (o *RaftOperation) Remove(k string) bool {
	return true
}
