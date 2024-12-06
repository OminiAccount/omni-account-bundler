package main

import (
	"github.com/OAB/cli"
	"github.com/OAB/utils/log"
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
		log.Info("Application failed", "message", err)
	}
}
