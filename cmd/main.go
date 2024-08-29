package main

import (
	"github.com/OAAC/cli"
	"github.com/ethereum/go-ethereum/log"
	"os"
)

var (
	GitVersion = ""
	GitCommit  = ""
	GitDate    = ""
)

func main() {
	app := cli.NewCli(GitVersion, GitCommit, GitDate)

	if err := app.Run(os.Args); err != nil {
		log.Crit("Application failed", "message", err)
	}
}
