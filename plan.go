package hotomata

import (
	"gopkg.in/yaml.v2"
)

type PlanVars map[string]interface{}

type PlanCall struct {
	Name  string
	Run   string
	Plan  string
	Local bool
	Vars  PlanVars
}

type Plan struct {
	Name        string
	DefaultVars PlanVars
	PlanCalls   []PlanCall
}

func ParsePlan(planName string, yamlSource []byte) (*Plan, error) {
	var plan = &Plan{Name: planName}

	// Unmarshal raw yaml
	var rawPlan struct {
		Vars  map[string]interface{}
		Plans []map[string]interface{}
	}
	err := yaml.Unmarshal(yamlSource, &rawPlan)
	if err != nil {
		return plan, err
	}

	// Fill structs that are nicer to work with
	plan.Vars = rawPlan.Vars
	for _, rawPlanCall := range rawPlan.Plans {
		planCall := &PlanCall{}

		// Parse PlanCall $name
		if rawName, ok := rawPlanCall["$name"]; ok {
			if name, ok := rawName.(string); ok {
				planCall.Name = name
			} else {
				return plan, newError("Error parsing plan: %s: $name is not a string", planName)
			}
		} else {
			return plan, newError("Error parsing plan: %s: $name is required", planName)
		}

		// Parse PlanCall $run
		if rawRun, ok := rawPlanCall["$run"]; ok {
			if run, ok := rawRun.(string); ok {
				planCall.Run = run
			} else {
				return plan, newError("Error parsing plan: %s: $run is not a string (%s)", planName, planCall.Name)
			}
		}

		// Parse PlanCall $run
		if rawPlanCallPlan, ok := rawPlanCall["$plan"]; ok {
			if planCallPlan, ok := rawPlanCallPlan.(string); ok {
				planCall.Plan = callPlanName
			} else {
				return plan, newError("Error parsing plan: %s: $plan is not a string (%s)", planName, planCall.Name)
			}
		}

		planCall.Local = false
		if local, ok := rawPlanCall["$local"].(bool); ok {
			planCall.Local = local
		} else {
			return plan, newError("Error parsing plan: %s: $local is not a bool (%s)", planName, planCall.Name)
		}

		// Verify we have an action to do
		if planCall.Run == "" && planCall.Plan == "" {
			return plan, newError("Error parsing plan: %s: $run or $plan is required (%s)", planName, planCall.Name)
		}

		plan.PlanCalls = append(plan.PlanCalls, planCall)
	}

	return plan, nil
}
