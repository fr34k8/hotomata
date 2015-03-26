package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/codegangsta/cli"
	"github.com/merd/hotomata"
)

func main() {
	app := cli.NewApp()
	app.Name = "hotomata-inventory"
	app.Usage = "tool to check validity of inventory files and introspect their contents as seen by hotomata"
	app.Version = "0.1.0"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Frederic Gingras",
			Email: "frederic@gingras.cc",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "check",
			Aliases: []string{"c"},
			Usage:   "Verifies a given inventory file is valid",
			Action: func(c *cli.Context) {
				contents, err := ioutil.ReadFile(c.Args().First())
				if err != nil {
					panic(err)
				}

				result, err := hotomata.ValidateInventory(string(contents))
				if err != nil {
					panic(err.Error())
				}

				if result.Valid() {
					fmt.Printf(hotomata.Colorize("The document is valid\n", hotomata.ColorGreen))
				} else {
					fmt.Printf(hotomata.Colorize("The document is not valid. see errors :\n", hotomata.ColorRed))
					for _, desc := range result.Errors() {
						fmt.Printf("- %s\n", desc)
					}
				}
			},
		},
		{
			Name:    "print",
			Aliases: []string{"p"},
			Usage:   "Prints the contents of an inventory file",
			Action: func(c *cli.Context) {
				println("completed task: ", c.Args().First())
			},
		},
	}

	app.Run(os.Args)
}
