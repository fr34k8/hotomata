package hotomata

import ()

type MasterPlan struct {
	MachineFilters []MachineFilter
}

type MachineFilter struct {
	Param       string
	valueString string
	valueArray  []string
}

func (mf *MachineFilter) MatchesMachine(machine *Machine) bool {
	return true
}

func ParseMasterPlan(yamlSource []byte) ([]*MasterPlan, error) {
	return []*MasterPlan{}, nil
}
