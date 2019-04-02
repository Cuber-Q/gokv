package server

import (
	"github.com/hashicorp/raft"
	"log"
	"os"
	"time"
	"path/filepath"
	"github.com/hashicorp/raft-boltdb"
	"net"
)

type raftNode struct {
	raft           *raft.Raft
	fsm            *FSM
	leaderNotifyCh chan bool
}

func NewRaftNode(cfg *OriginConfig, ctx *Context) (*raftNode, error){

	raftConfig := raft.DefaultConfig()
	raftConfig.LocalID = raft.ServerID(cfg.RaftTCPAddress)
	raftConfig.Logger = log.New(os.Stderr, "raft: ", log.Ldate|log.Ltime)
	raftConfig.SnapshotInterval = 20 * time.Second
	raftConfig.SnapshotThreshold = 2
	leaderNotifyCh := make(chan bool, 1)
	raftConfig.NotifyCh = leaderNotifyCh

	transport, err := newRaftTransport(cfg)
	if err != nil {
		return nil, err
	}

	if err := os.MkdirAll(cfg.DataDir, 0700); err != nil {
		return nil, err
	}

	fsm := &FSM{
		Ctx: ctx,
		Log: log.New(os.Stderr, "FSM: ", log.Ldate|log.Ltime),
	}
	snapshotStore, err := raft.NewFileSnapshotStore(cfg.DataDir, 1, os.Stderr)
	if err != nil {
		return nil, err
	}

	logStore, err := raftboltdb.NewBoltStore(filepath.Join(cfg.DataDir, "raft-log.bolt"))
	if err != nil {
		return nil, err
	}

	stableStore, err := raftboltdb.NewBoltStore(filepath.Join(cfg.DataDir, "raft-stable.bolt"))
	if err != nil {
		return nil, err
	}

	_raftNode, err := raft.NewRaft(raftConfig, fsm, logStore, stableStore, snapshotStore, transport)
	if err != nil {
		return nil, err
	}

	// is the first raft node
	if cfg.Leader {
		configuration := raft.Configuration{
			Servers: []raft.Server{
				{
					ID:      raftConfig.LocalID,
					Address: transport.LocalAddr(),
				},
			},
		}
		_raftNode.BootstrapCluster(configuration)
	}

	return &raftNode{raft: _raftNode, fsm: fsm, leaderNotifyCh: leaderNotifyCh}, nil

}


func newRaftTransport(cfg *OriginConfig) (*raft.NetworkTransport, error) {
	address, err := net.ResolveTCPAddr("tcp", cfg.RaftTCPAddress)
	if err != nil {
		return nil, err
	}
	transport, err := raft.NewTCPTransport(address.String(), address, 3, 10*time.Second, os.Stderr)
	if err != nil {
		return nil, err
	}
	return transport, nil
}

func (self *raftNode) AddNode(node string) string {
	addFuture := self.raft.AddVoter(raft.ServerID(node),raft.ServerAddress(node), 0, 0)

	if err := addFuture.Error(); err != nil {
		log.Printf("Error addNode, node:%s, err:%v", node, err)
		return "error"
	}

	return "ok"
}