package server

import (
	//"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"testing"
	"time"
)

var leader *Node
var f1 *Node
var f2 *Node

var l_conf = &NodeConf{
	ip:             "127.0.0.1",
	port:           9000,
	restIp:         "127.0.0.1",
	restPort:       8000,
	role:           LEADER,
	raftTCPAddress: "127.0.0.1:10000",
	dataDir:        "/Users/Cuber_Q/dev/gokv/db/127.0.0.1:9000",
}

var f1_conf = &NodeConf{
	ip:             "127.0.0.1",
	port:           9001,
	restIp:         "127.0.0.1",
	restPort:       8001,
	role:           FOLLOWER,
	raftTCPAddress: "127.0.0.1:10001",
	dataDir:        "/Users/Cuber_Q/dev/gokv/db/127.0.0.1:9001",
}

var f2_conf = &NodeConf{
	ip:             "127.0.0.1",
	port:           9002,
	restIp:         "127.0.0.1",
	restPort:       8002,
	role:           FOLLOWER,
	raftTCPAddress: "127.0.0.1:10002",
	dataDir:        "/Users/Cuber_Q/dev/gokv/db/127.0.0.1:9002",
}

var k = "aKey"
var v = "aValue"

func TestAddNode(t *testing.T) {
	node := &Node{id: "1", ip: "127.0.0.1", port: 8000, role: LEADER}
	GetCluster().AddNode(node)
	log.Println(cluster)
	log.Println(LEADER)
}

func Test_Leader(t *testing.T) {
	go func() {
		// test all leader node and follower nodes start
		leader = testNodesStart(l_conf)

		// test leader AddVoter, to construct a replicate cluster
		time.Sleep(10 * time.Second)
		testAddRaftVoter(f1_conf)
		//testAddRaftVoter(f2_conf)

		// test leader Set and leader Get
		testLeaderOperation(k, v)

		// test leader multi set sop
		testLeaderMultiSet()

		// test leader shutdown
		//testLeaderShutdown()

	}()

	time.Sleep(1000 * 1000 * time.Millisecond)
}

func Test_F1(t *testing.T) {
	go func() {
		f1 = testNodesStart(f1_conf)

		time.Sleep(15 * time.Second)
		// test followers Get
		testFollowerGet(f1, k)

		time.Sleep(5 * time.Second)
		// test follower Set
		testFollowerSet(f1)
	}()
	time.Sleep(1000 * 1000 * time.Millisecond)
}

func TestF2(t *testing.T) {
	go func() {
		f2 = testNodesStart(f2_conf)

		time.Sleep(15 * time.Second)
		// test followers Get
		testFollowerGet(f2, k)

		time.Sleep(5 * time.Second)
		// test follower Set
		testFollowerSet(f2)
	}()
	time.Sleep(1000 * 1000 * time.Millisecond)
}

func testNodesStart(conf *NodeConf) *Node {
	node := newNode(conf)
	log.Println(node.String())
	time.Sleep(5 * time.Second)
	return node
}

func testAddRaftVoter(conf *NodeConf) {
	leader.Join(conf.raftTCPAddress)
}

func testLeaderOperation(k, v string) {
	log.Printf("ready to testLeaderOperation k=%v,v=%v", k, v)
	leader.sop.Set(k, v)
	log.Printf("node:%v [Get] k=%v, v=%v", leader.id, k, leader.sop.Get(k))
}

func testLeaderMultiSet() {
	for i := 0; i < 100; i++ {
		leader.sop.Set("k:"+strconv.Itoa(i), "v:"+strconv.Itoa(i))
	}
}

func testFollowerGet(f *Node, k string) {
	time.Sleep(5 * time.Second)
	log.Printf("ready to testFollowerGet k=%v", k)
	log.Printf("node:%v [Get] k=%v, v=%v", f.id, k, f.sop.Get(k))
}

func testFollowerSet(f *Node) {
	time.Sleep(5 * time.Second)
	log.Printf("ready to testFollowerSet")
	f.sop.Set("bKey", "bValue")
}

func testLeaderShutdown() {
	time.Sleep(20 * time.Second)
	leader.raftNode.raft.Shutdown()
}
