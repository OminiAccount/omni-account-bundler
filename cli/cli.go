package cli

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/omni-account/client/config"
	"github.com/omni-account/client/processor"
	"github.com/urfave/cli/v2"

	orLog "github.com/omni-account/client/utils/log"
	"os"
	"os/signal"
	"syscall"
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
		Value:   "../../spv_develop.toml",
		Aliases: []string{"c"},
		Usage:   "path to config file",
		EnvVars: []string{"SPV_CONFIG"},
	}
	ProveModeFlag = &cli.UintFlag{
		Name:    "prove-mode",
		Value:   0,
		Aliases: []string{"p"},
		Usage:   "orbiter-bridge-prove or spv-front-prove?",
	}
)

// runProcessor is the entrypoint into the SPV Processor service.
func runProcessor(ctx *cli.Context) error {
	configPath := ctx.String(ConfigFlag.Name)
	conf, err := config.LoadConfig(configPath)

	log.Info("Starting processor")

	if err != nil {
		log.Crit("Failed to load config", "message", err)
	}

	processor, err := processor.NewProcessor(conf)
	if err != nil {
		log.Error("Unable to create processor", "error", err)
		return err
	}

	if err = processor.Start(); err != nil {
		return err
	}

	log.Info("Processor started")

	signalChan := make(chan os.Signal)
	// SIGHUP: terminal closed
	// SIGINT: Ctrl+C
	// SIGTERM: program exit
	// SIGQUIT: Ctrl+/
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	waitElegantExit(signalChan, processor)

	return nil
}

func waitElegantExit(signalChan chan os.Signal, processor *processor.Processor) {
	for i := range signalChan {
		switch i {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			processor.Stop()
			os.Exit(0)
		}
	}
}

// Run make an instance method on Cli called Run that runs cli
// and returns an error
func (c *Cli) Run(args []string) error {
	return c.app.Run(args)
}

func NewCli(GitVersion string, GitCommit string, GitDate string) *Cli {
	orLog.SetupDefaults()

	flags := []cli.Flag{
		ConfigFlag,
		ProveModeFlag,
	}

	app := &cli.App{
		Name:        "Processor",
		Version:     "test",
		Description: "A spv processor",
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
