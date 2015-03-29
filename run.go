package hotomata

import (
	"io/ioutil"
	"path"
	"strings"
)

const planFileExt = ".yaml"
const logFillRune = '-'

type Run struct {
	plans     map[string]*Plan
	inventory []InventoryMachine
}

func NewRun() *Run {
	return &Run{
		plans:     map[string]*Plan{},
		inventory: []InventoryMachine{},
	}
}

func (r *Run) DiscoverPlans(directory string) error {
	var loadFolder func(string) error
	loadFolder = func(folder string) error {
		folders, err := ioutil.ReadDir(folder)
		if err != nil {
			return err
		}
		for _, f := range folders {
			if f.IsDir() {
				err = loadFolder(path.Join(folder, f.Name()))
				if err != nil {
					return err
				}
				continue
			} else if !strings.HasSuffix(f.Name(), planFileExt) {
				continue
			}

			// Ok, at this point we got a .yaml file to load
			contents, err := ioutil.ReadFile(path.Join(folder, f.Name()))
			if err != nil {
				return err
			}

			planName := strings.TrimSuffix(f.Name(), planFileExt)
			plan, err := ParsePlan(planName, contents)
			if err != nil {
				return err
			}

			r.plans[planName] = plan
		}

		return nil
	}

	return loadFolder(directory)
}

func (r *Run) Plan(name string) (*Plan, bool) {
	plan, ok := r.plans[name]
	return plan, ok
}

func (r *Run) Plans() map[string]*Plan {
	return r.plans
}

func (r *Run) LoadInventory(machines []InventoryMachine) {
	r.inventory = append(r.inventory, machines...)
}

func (r *Run) RunMasterPlans(logger Logger, masterplans []*MasterPlan) (*RunReport, error) {
	report := &RunReport{}

	for _, masterplan := range masterplans {
		err := r.RunMasterPlan(logger, report, masterplan)
		if err != nil {
			return report, err
		}
	}

	return report, nil
}

func (r *Run) RunMasterPlan(logger Logger, report *RunReport, masterplan *MasterPlan) error {
	machines := masterplan.FilterMachines(r.inventory)

	logRunStart(logger, machines)

	return nil
}

func logRunStart(logger Logger, machines []InventoryMachine) {
	var runMachineNames string
	for _, m := range machines {
		runMachineNames = runMachineNames + m.Name + " "
	}
	logLine(logger, logFillRune, "RUN: [ %s]", runMachineNames)
}
