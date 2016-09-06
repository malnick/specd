package resources

import (
	gossResource "github.com/malnick/goss/resource"
	gossSystem "github.com/malnick/goss/system"
)

type Package struct {
	gossResource.Package `yaml:",inline"`
}

func (f Package) Validate(gs *gossSystem.System) []gossResource.TestResult {
	tr := f.Package.Validate(gs)
	return tr
}
