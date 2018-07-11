package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "gokv",
	Short: "gokv is a simple fast k-v store system coded in go",
}

func init() {

	RootCmd.AddCommand(
		NewServerCmd(),
	)

	cobra.EnablePrefixMatching = true
}