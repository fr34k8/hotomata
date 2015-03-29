package hotomata

import (
	"io/ioutil"
	"path"
	"strings"
)

const planFileExt = ".yaml"

type Run struct {
	plans map[string]*Plan
}

func NewRun() *Run {
	return &Run{plans: map[string]*Plan{}}
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
