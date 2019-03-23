package cluster

import (
	"testing"
	"log"
)

func TestAddNode(t *testing.T) {
	node := &Node{id:"1", ip:"127.0.0.1", port:8000, role:MASTER}
	AddNode(node)
	log.Println(cluster)
	log.Println(MASTER)
}