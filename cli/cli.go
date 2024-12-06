package cli

import (
	"github.com/OAB/config"
	"github.com/OAB/processor"
	"github.com/OAB/utils/log"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Cli struct {
	GitVersion string
	GitCommit  string
	GitDate    string
	app        *cli.App
	Flags      []cli.Flag
}

var (
	ConfigFlag = &cli.StringFlag{
		Name:    "config",
		Value:   "./develop.toml",
		Aliases: []string{"c"},
		Usage:   "path to config file",
		EnvVars: []string{"CONFIG"},
	}
)

// runProcessor is the entrypoint into the OAB(Omni Account Abstraction Client) Processor service.
func runProcessor(ctx *cli.Context) error {
	conf, err := config.Load(ctx)
	if err != nil {
		panic("Failed to load config, error:"+err.Error())
	}
	log.Init(conf.Log)
	log.Infof("%+v", conf)
	log.Info("Starting OAB processor")

	proc, err := processor.NewProcessor(conf)
	if err != nil {
		log.Error("Unable to create OAB processor", "error", err)
		return err
	}

	if err = proc.Start(); err != nil {
		return err
	}

	log.Info("OAB Processor started")

	signalChan := make(chan os.Signal, 1)
	// SIGHUP: terminal closed
	// SIGINT: Ctrl+C
	// SIGTERM: program exit
	// SIGQUIT: Ctrl+/
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	waitElegantExit(signalChan, proc)

	return nil
}

func waitElegantExit(signalChan chan os.Signal, proc *processor.Processor) {
	s := <-signalChan
	log.Infof("receive exit signal: %v", s)
	proc.Stop()
	time.Sleep(time.Second*5)
	os.Exit(0)
}

// Run make an instance method on Cli called Run that runs cli
// and returns an error
func (c *Cli) Run(args []string) error {
	return c.app.Run(args)
}

func NewCli(GitVersion string, GitCommit string, GitDate string) *Cli {
	flags := []cli.Flag{
		ConfigFlag,
	}

	app := &cli.App{
		Name:        "Processor",
		Version:     "test",
		Description: "An OAB Processor",
		Commands: []*cli.Command{
			{
				Name:        "processor",
				Flags:       flags,
				Description: "Runs the processor service",
				Action:      runProcessor,
			},
		},
	}

	return &Cli{
		app:   app,
		Flags: flags,
	}
}
