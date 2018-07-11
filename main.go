package main

import (
	"gokv/cmd"
	"os"
	"fmt"
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