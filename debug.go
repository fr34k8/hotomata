package main

import (
	"io/ioutil"

	"github.com/codegangsta/cli"
	"github.com/davecgh/go-spew/spew"
	"github.com/merd/hotomata"
)

func debugPlan(c *cli.Context) {
	var contents []byte
	var err error

	// Parse plan args
	var planName = c.Args().First()
	if planName == "" {
		writeError("Error: A plan is required. e.g. `hotomata debug plan redis`", nil)
	}

	// Discover plans
	contents, err = ioutil.ReadFile(masterPlanFile)
	if err != nil {
		writeError("Error: Unable to read masterplan file at "+masterPlanFile, err)
	}
	// Fetch concerned plan
	masterplans, err := hotomata.ParseMasterPlan(contents)
	if err != nil {
		writeError("Error: Unable to parse masterplan file, verify your YAML syntax", err)
	}
	spew.Dump(masterplans)
}
