//go:build !appengine
// +build !appengine

package main

import (
	"context"
	"github.com/gonuts/commander"
	_ "net/http/pprof"

	"fmt"
	"os"
	"yap/app"
	"yap/webapi"
)

var cmd = &commander.Command{
	UsageLine: os.Args[0] + " app|api",
	Short:     "invoke yap as a standalone app or as an api server",
}

func init() {
	cmd.Subcommands = append(app.AllCommands().Subcommands, webapi.AllCommands().Subcommands...)
	//cmd.Subcommands = app.AllCommands().Subcommands
}

func exit(err error) {
	fmt.Printf("**error**: %v\n", err)
	os.Exit(1)
}

func main() {
	// Create a context
	ctx := context.Background()

	// Dispatch command with the context and command-line arguments
	if err := cmd.Dispatch(ctx, os.Args[1:]); err != nil {
		exit(err)
	}

	return
}
