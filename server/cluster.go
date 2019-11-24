package server

import (
	"log"
	"strings"
)

type Cluster struct {
	currNode *Node
	nodes    []*Node
	status   string
}

// cluster config, include leader node and flowers nodes's ips and ports
type ClusterConf struct {
	curr   *NodeConf
	fellow []*NodeConf
}

// create new cluster instance
// it should be called only once
func NewCluster(clusterConf *ClusterConf) *Cluster {
	cluster := &Cluster{
		nodes:  createNodes(clusterConf),
		status: "creating",
	}
	return cluster
}

func createNodes(conf *ClusterConf) []*Node {
	nodes := make([]*Node, 0)
	nodes = append(nodes, newNode(conf.curr))
	for i, nodeConf := range conf.fellow {
		log.Println("creating the num [%d] flower: %s", i, nodeConf)
		nodes = append(nodes, newNode(nodeConf))
	}
	return nodes
}

func (c *Cluster) String() string {
	s := "Cluster=[" +
		"status=" + c.status +
		", nodes=["
	arr := make([]string, 0)
	for i := 0; i < len(c.nodes); i++ {
		arr = append(arr, c.nodes[i].String())
	}
	s += strings.Join(arr, ",")
	s += "]"
	return s
}

var cluster = &Cluster{
	status: "OK",
	nodes:  make([]*Node, 0),
}

func GetCluster() *Cluster {
	return cluster
}

func (c *Cluster) AddNode(node *Node) {
	mutex.Lock()
	defer mutex.Unlock()

	cluster.nodes = append(cluster.nodes, node)
}
