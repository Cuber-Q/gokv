package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"gokv/gokv-cli/handler"
)

func NewGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:"get",
		Short:"get value of <key>",
		Run: func(cmd *cobra.Command, args []string) {
			runGet(args)
		},
	}
	return cmd
}
func runGet(args []string) {
	if len(args) < 1 {
		fmt.Println("get command's arg shoule just ONE, for example: get <key> ")
		return
	}

	//value := core.Get(args[0])
	value := handler.Get(args[0])
	fmt.Println(value)
}
