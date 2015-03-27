package main

import (
	"io/ioutil"

	"github.com/codegangsta/cli"
	"github.com/merd/hotomata"
)

func run(c *cli.Context) {
	var contents []byte
	var err error

	// Parse masterplan args
	var masterPlanFile = c.Args().First()
	if masterPlanFile == "" {
		writeError("Error: A masterplan is required. e.g. `hotomata masterplan.yml`", nil)
	}

	// Parse inventory args
	var inventoryFile = c.GlobalString("inventory")
	if inventoryFile == "" {
		writeError("Error: An inventory file is required. e.g. `hotomata --inventory inventory.json`", nil)
	}

	// Parse actual inventory
	contents, err = ioutil.ReadFile(inventoryFile)
	if err != nil {
		writeError("Error: Unable to read inventory file at "+inventoryFile, err)
	}
	_, err = hotomata.ParseInventory(contents)
	if err != nil {
		writeError("Error: Unable to parse inventory file, verify your JSON syntax", err)
	}

	// Parse actual masterplan
	contents, err = ioutil.ReadFile(masterPlanFile)
	if err != nil {
		writeError("Error: Unable to read masterplan file at "+masterPlanFile, err)
	}
	_, err = hotomata.ParseMasterPlan(contents)
	if err != nil {
		writeError("Error: Unable to parse masterplan file, verify your YAML syntax", err)
	}

}
