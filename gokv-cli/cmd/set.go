package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"strings"
	"gokv/gokv-cli/handler"
)

func NewSetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:"set",
		Short:"set <key> <value>",
		Run: func(cmd *cobra.Command, args []string) {
			runSet(args)
		},
	}


	return cmd
}

func runSet(args []string) {
	if len(args) < 2 {
		fmt.Println("set command's args should not less than 2 args, for example: set your_key your_value")
		return
	}

	key := args[0]

	valueSlice := args[1:]
	value := strings.Join(valueSlice,"")

	result := handler.Set(key, value)
	fmt.Println(result)
}
