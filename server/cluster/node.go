package cluster

import (
	"strconv"
	"strings"
)

type NodeRole int32

const (
	MASTER  NodeRole = 0
	FOLLOWER	NodeRole = 1
)

func (r NodeRole) String() string {
	switch r {
	case MASTER:
		return "Master"
	case FOLLOWER:
		return "Follower"
	default:
		return ""
	}
}

type Node struct {
	id string
	ip string
	port int
	role NodeRole
}

func (n *Node) String() string {
	return "Node=["+
		"id="+n.id+
		", ip="+n.ip+
		", port="+strconv.Itoa(n.port)+
		", role="+n.role.String()+
		"]"
}

type Cluster struct {
	nodes []*Node
	status string
}

func (c *Cluster) String() string {
	s := "Cluster=["+
		"status=" + c.status +
		", nodes=["
	arr:= make([]string,0)
	for i:=0;i< len(c.nodes);i++ {
		arr = append(arr, c.nodes[i].String())
	}
	s += strings.Join(arr, ",")
	s += "]"
	return s
}

var cluster = &Cluster{
	status:"OK",
	nodes: make([]*Node,0),
}

func AddNode(node *Node) bool {
	cluster.nodes = append(cluster.nodes, node)
	return true
}


