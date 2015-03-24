package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/xeipuuv/gojsonschema"
)

const inventorySchema = `
{
  "type": "object"
}
`

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
				bytes, err := ioutil.ReadFile(c.Args().First())
				if err != nil {
					panic(err)
				}

				schemaLoader := gojsonschema.NewStringLoader(inventorySchema)
				loader := gojsonschema.NewStringLoader(string(bytes))

				result, err := gojsonschema.Validate(schemaLoader, documentLoader)
				if err != nil {
					panic(err.Error())
				}

				if result.Valid() {
					fmt.Printf("The document is valid\n")
				} else {
					fmt.Printf("The document is not valid. see errors :\n")
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
