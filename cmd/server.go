package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
	"gokv/server"
	"fmt"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start gokv server",
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}

var (
	port = 9901
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

	cmd.Flags().IntVarP(&port, "port", "P", 9901, "specify server port")

	return cmd
}

func runServer() {
	// start server
	server.Server(port)

	fmt.Println("server has running at port:", port)
	fmt.Println("press CTRL + C to exit")

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
