package hotomata

import (
	"gopkg.in/yaml.v2"
)

type MasterPlan struct {
	MachineFilters []*MachineFilter
	Vars           PlanVars
	Plans          []string
}

type MachineFilter struct {
	Param   string
	Pattern string
}

func (mf *MachineFilter) MatchesMachine(machine *Machine) bool {
	return true
}

func ParseMasterPlan(yamlSource []byte) ([]*MasterPlan, error) {
	var plans = []*MasterPlan{}

	// Unmarshal raw yaml
	var rawPlans []struct {
		Machines map[string]string
		Vars     map[string]interface{}
		Plans    []string
	}
	err := yaml.Unmarshal(yamlSource, &rawPlans)
	if err != nil {
		return plans, err
	}

	// Fill structs that are nicer to work with
	for _, rawPlan := range rawPlans {
		plan := &MasterPlan{MachineFilters: []*MachineFilter{}}

		for k, v := range rawPlan.Machines {
			plan.MachineFilters = append(plan.MachineFilters, &MachineFilter{
				Param:   k,
				Pattern: v,
			})
		}

		plan.Vars = rawPlan.Vars

		plan.Plans = rawPlan.Plans

		plans = append(plans, plan)
	}

	return plans, nil
}
