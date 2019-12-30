package main

import (
	"fmt"
	"gokv/server/cmd"
	"log"
	"os"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	cobraRun()
}

func cobraRun() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
