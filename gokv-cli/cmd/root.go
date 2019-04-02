package cmd


import (
	"github.com/spf13/cobra"
	"log"
	"github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
	"fmt"
	"strings"
	"os"
	"os/exec"
	"gokv/gokv-cli/config"
)

var RootCmd = &cobra.Command{
	Use:   "gokv-cli",
	Short: "gokv-cli is a simple fast k-v store system command app coded in go",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	RootCmd.AddCommand(
		NewSetCmd(),
		NewGetCmd(),
		NewAddNodeCmd(),
	)
	RootCmd.PersistentFlags().StringVar(&config.Ctx.Endpoint,"endpoint", "127.0.0.1:9901","specify endponits")
	cobra.EnablePrefixMatching = true
}

func run()  {
	log.Println("current endpoints are:\n", config.Ctx.Endpoint)
	p := prompt.New(
		Executor,
		Completer,
		prompt.OptionTitle("gokv cli-prompt: interactive gokv-cli"),
		prompt.OptionPrefix("gokv-cli>"),
		prompt.OptionInputTextColor(prompt.Yellow),
		prompt.OptionCompletionWordSeparator(completer.FilePathCompletionSeparator),
	)
	p.Run()
}

func Completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "set", Description: "set <key> <value>"},
		{Text: "get", Description: "get <key>"},
		{Text: "exist", Description: "check out if <key> exist"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func Executor(s string)  {
	s = strings.TrimSpace(s)
	if s == "" {
		return
	} else if s == "quit" || s == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	}

	cmd := exec.Command("/bin/sh", "-c", "gokv-cli "+s)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}
}