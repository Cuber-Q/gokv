package server

import (
	"gokv/core"
	"log"
	"strconv"
	"strings"
	"sync"
)

type NodeRole int32

const (
	LEADER   NodeRole = 0
	FOLLOWER NodeRole = 1
)

var mutex sync.Mutex

func (r NodeRole) String() string {
	switch r {
	case LEADER:
		return "Leader"
	case FOLLOWER:
		return "Follower"
	default:
		return ""
	}
}

// main struct
type Node struct {
	// unique node id
	id string

	// service for gokv-cli
	ip   string
	port int

	// service for rest-client
	restIp   string
	restPort int
	restMux  *RestMux

	// raft node instance
	role           NodeRole
	raftTCPAddress string
	raftNode       *raftNode

	// local storage and it's operation
	store *core.Storage
	sop   core.StoreOperation
}

func newNode(conf *NodeConf) *Node {
	node := &Node{
		id:       conf.ip + ":" + strconv.Itoa(conf.port),
		ip:       conf.ip,
		port:     conf.port,
		restIp:   conf.restIp,
		restPort: conf.restPort,
		role:     conf.role,
	}

	// create local storage
	storage := core.NewStorage()

	// create raft node which could make it work in cluster mode
	// when a raft node is created, it will listen on ip:port and connect
	// other node automatically and run the cluster
	node.raftNode = newRaftNode(conf, storage)

	// create the core sop of cluster
	node.sop = newRaftOperation(node.raftNode, storage)

	// create the rest server which supplies rest sop for user
	node.restMux = newRestMux(conf.restIp, conf.restPort, node.sop, node)

	// start rest server async
	go func() {
		node.restMux.start()
	}()

	return node
}

// implementation of ClusterOperation interface.
// join a new node to the current raft cluster.
// Note that the join sop will success only when the
// new node has to start up itself successfully.
func (n *Node) Join(newNode string) string {
	if !n.raftNode.isLeader {
		return "not leader"
	}

	addr := strings.Split(newNode, ":")
	if len(addr) < 2 || addr[0] == "" || addr[1] == "" {
		log.Fatalf("invalid newNode addr:%v", newNode)
		return "param error"
	}

	result := n.raftNode.AddNode(newNode)
	if "ok" == result {
		log.Printf("join node:%v success!", newNode)
	} else {
		log.Printf("Error joining peer to raftNode, newNode:%s", newNode)
	}

	return result
}

func (n *Node) String() string {
	return "id=" + n.id +
		", Node=[" +
		"ip=" + n.ip +
		", port=" + strconv.Itoa(n.port) +
		", role=" + n.role.String() +
		", raftNode=" + n.raftNode.raft.String() +
		"]"
}

type NodeConf struct {
	ip             string
	port           int
	role           NodeRole
	restIp         string
	restPort       int
	raftTCPAddress string
	dataDir        string
}
