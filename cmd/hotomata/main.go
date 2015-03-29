package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/merd/hotomata"
)

func main() {
	app := cli.NewApp()
	app.Name = "hotomata"
	app.Usage = "tool to execute masterplans and do remdote execution"
	app.Version = "0.1.0"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Frederic Gingras",
			Email: "frederic@gingras.cc",
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "inventory, i",
			Value:  "inventory.json",
			Usage:  "inventory file location",
			EnvVar: "HOTOMATA_INVENTORY_FILE",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "Runs a given masterplan against an inventory of machines",
			Action:  run,
		},
		{
			Name:  "debug",
			Usage: "Few sub commands for debuging your plans",
			Subcommands: []cli.Command{
				{
					Name:   "plan",
					Usage:  "Visualise a specific plan",
					Action: debugPlan,
				},
			},
		},
	}

	app.Run(os.Args)
}

func writeError(message string, err error) {
	var completeMessage = message
	if err != nil {
		completeMessage = fmt.Sprintf("%s (%s)", completeMessage, err.Error())
	}
	fmt.Print(hotomata.Colorize(completeMessage+"\n", hotomata.ColorRed))
	os.Exit(1)
}
