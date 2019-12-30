package cmd

import (
	"log"
	"gokv/server"
	"os"
	"os/signal"
	"syscall"
	"github.com/spf13/cobra"
	"strconv"
)

var (
	port       = 9901
	host       = "127.0.0.1"
	cluster    = []string{}
	serverType = "singleton"
	baseDataDir = "~/gokv/data/"
	raftPort = 10000
	leader = false
	//stop <-chan bool
)

func NewServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "start gokv server",
		Run: func(cmd *cobra.Command, args []string) {
			runServer()
		},
	}

	cmd.Flags().IntVarP(&port, "port", "p", 9901, "specify server rest port")
	cmd.Flags().IntVar(&raftPort, "raftPort",  10001, "specify server raft port")
	cmd.Flags().StringVar(&host, "host", "127.0.0.1",
		"specify server host")
	cmd.Flags().StringVarP(&serverType, "type", "t", "singleton",
		"specify server type. you can config server as 'singleton' or 'cluster'")
	cmd.Flags().StringVar(&baseDataDir, "baseDataDir", "","specify persistent base data dir")
	cmd.Flags().StringArrayVarP(&cluster, "cluster", "c", []string{},
		"specify cluster info, a <ip>:<_port> string with ',' to split. for example: --cluster=127.0.0.1:9902,127.0.0.1:9903")
	cmd.Flags().BoolVar(&leader, "leader", false, "specify current node is leader of the cluster. Only effect in cluster mode")
	return cmd
}

func runServer() {
	// start server
	log.Println("server will run on: ", host, port)
	log.Println("server has has configured as: ", serverType)
	if len(cluster) > 0 {
		log.Println("server cluster info: ", cluster)
	}
	log.Println("press CTRL + C to exit")


	cfg := &server.OriginConfig{
		Port:       port,
		Host:       host,
		Cluster:    cluster,
		ServerType: serverType,
		DataDir: baseDataDir + host,
		RaftPort: raftPort,
		RaftTCPAddress:host + ":" + strconv.Itoa(raftPort),
		Leader: leader,
	}

	server.NewServer(cfg)

	// listen os signals to exit
	var sigs = make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	//if stop != nil {
	//	select {
	//	case <-sigs:
	//	case <-stop:
	//	}
	//} else {
	//	<-sigs
	//}
	<-sigs
}
