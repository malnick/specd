package cli

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"

	"github.com/malnick/specd/api"
	"github.com/malnick/specd/config"
)

var args = os.Args

func printBanner() {
	fmt.Println(`                    ___ `)
	fmt.Println(` ___ ___  ___  ___ | . \`)
	fmt.Println(`<_-<| . \/ ._>/ | '| | |`)
	fmt.Println(`/__/|  _/\___.\_|_.|___/`)
	fmt.Println(`    |_|                 `)

}

func Start() error {
	printBanner()
	app := cli.NewApp()
	app.Name = "specD"
	app.Usage = "A lightweight server inspection utility and HTTP API implemented in Go"
	appConfig := config.Configuration()

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "verbose",
			Usage:       "Verbose logging mode.",
			Destination: &appConfig.FlagVerbose,
		},
		cli.BoolFlag{
			Name:        "json-log",
			Usage:       "JSON logging mode.",
			Destination: &appConfig.FlagJSONLog,
		},
		cli.StringFlag{
			Name:        "state-path, s",
			Usage:       "Path to state.yaml",
			Value:       appConfig.StatePath,
			Destination: &appConfig.StatePath,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:     "api",
			Usage:    "Run specd API.",
			Category: "run",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:        "port, p",
					Value:       appConfig.FlagAPIPort,
					Usage:       "API port when running in API mode.",
					Destination: &appConfig.FlagAPIPort,
				},
			},
			Action: func(c *cli.Context) {
				if appConfig.FlagVerbose {
					log.SetLevel(log.DebugLevel)
				}
				api.Start(appConfig)
				os.Exit(0)
			},
		},
		{
			Name:     "report",
			Usage:    "Run specd without enforcing state; generate report only",
			Category: "run",
			Action: func(c *cli.Context) {
				if appConfig.FlagVerbose {
					log.SetLevel(log.DebugLevel)
				}
				if err := RunReport(appConfig); err != nil {
					log.Error(err)
					os.Exit(1)
				}
				os.Exit(0)
			},
		},
	}
	app.Run(args)
	return nil
}
