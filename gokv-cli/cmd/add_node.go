package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"gokv/gokv-cli/handler"
)

func NewAddNodeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:"addNode",
		Short:"execute on a leader node to add a new node as its follower. specify address of new node",
		Run: func(cmd *cobra.Command, args []string) {
			runAddNode(args)
		},
	}
	return cmd
}
func runAddNode(args []string) {
	if len(args) < 1 {
		fmt.Println("addNode command's arg shoule just ONE, for example: addNode <ip>:<raft_port> ")
		return
	}

	//value := core.Get(args[0])
	value := handler.AddNode(args[0])
	fmt.Println(value)
}
