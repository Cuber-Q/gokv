package server

import (
	"fmt"
	"github.com/hashicorp/raft"
	"github.com/hashicorp/raft-boltdb"
	"gokv/core"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type raftNode struct {
	raft           *raft.Raft
	fsm            *FSM
	leaderNotifyCh chan bool
	isLeader       bool
}

func newRaftNode(conf *NodeConf, storage *core.Storage) *raftNode {
	raftConfig := raft.DefaultConfig()
	raftConfig.LocalID = raft.ServerID(conf.raftTCPAddress)
	raftConfig.Logger = log.New(os.Stderr, "raftNode: ", log.Ldate|log.Ltime)
	raftConfig.SnapshotInterval = 20 * time.Second
	raftConfig.SnapshotThreshold = 2
	leaderNotifyCh := make(chan bool, 1)
	raftConfig.NotifyCh = leaderNotifyCh
	//raftConfig.StartAsLeader = conf.role == LEADER

	transport, err := _newRaftTransport(conf)
	if err != nil {
		log.Panicln(err)
		return nil
	}

	if err := os.MkdirAll(conf.dataDir, 0700); err != nil {
		log.Panicln(err)
		return nil
	}

	fsm := &FSM{
		Id:      conf.ip + ":" + strconv.Itoa(conf.port),
		Ctx:     &Context{},
		Log:     log.New(os.Stderr, "FSM: ", log.Ldate|log.Ltime),
		Storage: storage,
	}
	snapshotStore, err := raft.NewFileSnapshotStore(conf.dataDir, 1, os.Stderr)
	if err != nil {
		log.Panicln(err)
		return nil
	}

	logStore, err := raftboltdb.NewBoltStore(filepath.Join(conf.dataDir, "raftNode-log.bolt"))
	if err != nil {
		log.Panicln(err)
		return nil
	}

	stableStore, err := raftboltdb.NewBoltStore(filepath.Join(conf.dataDir, "raftNode-stable.bolt"))
	if err != nil {
		log.Panicln(err)
		return nil
	}

	_raft, err := raft.NewRaft(raftConfig, fsm, logStore, stableStore, snapshotStore, transport)
	if err != nil {
		log.Panicln(err)
		return nil
	}

	// is the first raftNode node
	if conf.role == LEADER {
		configuration := raft.Configuration{
			Servers: []raft.Server{
				{
					ID:      raftConfig.LocalID,
					Address: transport.LocalAddr(),
				},
			},
		}
		fmt.Println("raftConfig.LocalID: " + configuration.Servers[0].ID)
		fmt.Println("raftConfig.Address: " + configuration.Servers[0].Address)
		_raft.BootstrapCluster(configuration)
	}

	_raftNode := &raftNode{
		raft:           _raft,
		fsm:            fsm,
		leaderNotifyCh: leaderNotifyCh,
		isLeader:       conf.role == LEADER,
	}
	return _raftNode
}

// deprecated
func NewRaftNode(cfg *OriginConfig, ctx *Context) (*raftNode, error) {
	raftConfig := raft.DefaultConfig()
	raftConfig.LocalID = raft.ServerID(cfg.RaftTCPAddress)
	raftConfig.Logger = log.New(os.Stderr, "raftNode: ", log.Ldate|log.Ltime)
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

	//snapshotStore, err := raftNode.NewFileSnapshotStore(cfg.DataDir, 1, os.Stderr)
	//if err != nil {
	//	return nil, err
	//}
	snapshotStore := raft.NewInmemSnapshotStore()

	logStore, err := raftboltdb.NewBoltStore(filepath.Join(cfg.DataDir, "raftNode-log.bolt"))
	if err != nil {
		return nil, err
	}

	stableStore, err := raftboltdb.NewBoltStore(filepath.Join(cfg.DataDir, "raftNode-stable.bolt"))
	if err != nil {
		return nil, err
	}

	_raft, err := raft.NewRaft(raftConfig, fsm, logStore, stableStore, snapshotStore, transport)
	if err != nil {
		return nil, err
	}

	// is the first raftNode node
	if cfg.Leader {
		configuration := raft.Configuration{
			Servers: []raft.Server{
				{
					ID:      raftConfig.LocalID,
					Address: transport.LocalAddr(),
				},
			},
		}
		_raft.BootstrapCluster(configuration)
	}

	raftNode := &raftNode{raft: _raft, fsm: fsm, leaderNotifyCh: leaderNotifyCh}

	listenLeaderChange(raftNode)
	return raftNode, nil

}

func listenLeaderChange(rn *raftNode) {
	go func() {
		for {
			select {
			case isLeader := <-rn.leaderNotifyCh:
				{
					log.Printf("leader status changed, current is leader:[%v]", isLeader)
					rn.isLeader = isLeader
				}
			}
		}
	}()
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

func _newRaftTransport(conf *NodeConf) (*raft.NetworkTransport, error) {
	address, err := net.ResolveTCPAddr("tcp", conf.raftTCPAddress)
	if err != nil {
		return nil, err
	}
	transport, err := raft.NewTCPTransport(address.String(), address, 3, 10*time.Second, os.Stderr)
	if err != nil {
		return nil, err
	}
	return transport, nil
}

func (r *raftNode) AddNode(node string) string {
	addFuture := r.raft.AddVoter(raft.ServerID(node), raft.ServerAddress(node), 0, 0)

	if err := addFuture.Error(); err != nil {
		log.Printf("Error AddNode, node:%s, err:%v", node, err)
		return "error"
	}

	return "ok"
}
