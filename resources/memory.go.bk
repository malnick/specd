package resources

import (
	gossResource "github.com/malnick/goss/resource"
	gossSystem "github.com/malnick/goss/system"
)

type Memory struct {
	gossResource.Memory `yaml:",inline"`
}

func (m Memory) Validate(gs *gossSystem.System) []gossResource.TestResult {
	tr := m.Memory.Validate(gs)
	return tr
}
