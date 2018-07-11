package main

import (
	"fmt"
	"os"
	"gokv/gokv-cli/cmd"
)

func main() {
	cobraRun()
}

func cobraRun() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}