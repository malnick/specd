package resources

import (
	gossResource "github.com/malnick/goss/resource"
	gossSystem "github.com/malnick/goss/system"
)

type Service struct {
	gossResource.Service `yaml:",inline"`
}

func (f Service) Validate(gs *gossSystem.System) []gossResource.TestResult {
	tr := f.Service.Validate(gs)
	return tr
}
