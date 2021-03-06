package server

import (
//"github.com/goinggo/mapstructure"
)

type GoKVServer struct {
	ip   string
	port int
	raft *raftNode
}

type OriginConfig struct {
	Port           int
	Host           string
	DataDir        string
	Cluster        []string
	ServerType     string
	JoinCluster    string
	RaftPort       int
	RaftTCPAddress string
	Leader         bool
}

var server = &GoKVServer{}

func (server *GoKVServer) start() {
	//log.Println("server has already on ", server.ip, ":", server.port)
	////http.ListenAndServe(server.ip +":"+ strconv.Itoa(server.port), newRestMux())
	//
	//var l net.Listener
	//var err error
	//l, err = net.Listen("tcp", server.ip+":"+strconv.Itoa(server.port))
	//if err != nil {
	//	log.Fatal(fmt.Sprintf("listen http failed"))
	//}
	//
	//go func() {
	//	//http.Serve(l, newRestMux("127.0.0.1", 7000, &RaftOperation{}))
	//}()
}

func NewServer(cfg *OriginConfig) {
	server = &GoKVServer{
		ip:   cfg.Host,
		port: cfg.Port,
	}

	if cfg.ServerType == "singleton" {
		// run as singleton
		server.start()
	} else if cfg.ServerType == "cluster" {
		// run as cluster
		server.raft, _ = NewRaftNode(cfg, &Context{})
		server.start()
	}
}

func AddNode(node string) string {
	return server.raft.AddNode(node)
}
