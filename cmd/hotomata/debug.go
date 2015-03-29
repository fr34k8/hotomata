package main

import (
	"fmt"
	"os"
	"path"

	"github.com/codegangsta/cli"
	"github.com/kiasaki/hotomata"
)

func setupDebug(c *cli.Context) *hotomata.Run {
	var err error
	var cwd string

	if cwd, err = os.Getwd(); err != nil {
		panic(err)
	}

	// Discover plans
	run := hotomata.NewRun()
	plansFolder := path.Join(cwd, "plans")
	if err = run.DiscoverPlans(plansFolder); err != nil {
		writeError("Error: Unable to load plans", err)
	}

	return run
}

func debugPlan(c *cli.Context) {
	run := setupDebug(c)

	// Parse plan args
	var planName = c.Args().First()
	if planName == "" {
		writeError("Error: A plan is required. e.g. `hotomata debug plan redis`", nil)
	}

	// Fetch concerned plan
	plan, ok := run.Plan(planName)
	if !ok {
		writeError("Error: Unable to find plan '"+planName+"'", nil)
	}

	writePlan("", run, plan, true)
}

func debugPlans(c *cli.Context) {
	run := setupDebug(c)

	writef(hotomata.ColorNone, "All plans\n")

	for _, p := range run.Plans() {
		writePlan("", run, p, false)
		fmt.Println("")
	}
}

func writePlan(in string, run *hotomata.Run, p *hotomata.Plan, detailed bool) {
	// Bump indentation each level
	in = in + "  "

	if detailed {
		writef(hotomata.ColorMagenta, "%sName: %s", in, p.Name)
	} else {
		writef(hotomata.ColorMagenta, "%s%s", in, p.Name)

	}

	if detailed {
		writef(hotomata.ColorCyan, "%sVars:", in)
		for k, v := range p.Vars {
			writef(hotomata.ColorGreen, "%s  %s: %v", in, k, v)
		}
	}

	if detailed {
		writef(hotomata.ColorCyan, "%sPlans:", in)
	}
	for _, planCall := range p.PlanCalls {
		if planCall.Run != "" {
			writef(hotomata.ColorGreen, "%s  $run: %s", in, planCall.Run)
		} else {
			plan, found := run.Plan(planCall.Plan)
			if found {
				writePlan(in, run, plan, detailed)
			} else {
				if detailed {
					writef(hotomata.ColorRed, "%s  Name: %s (missing)", in, planCall.Plan)
				} else {
					writef(hotomata.ColorRed, "%s  %s (missing)", in, planCall.Plan)
				}
			}
		}
	}
}
