package resources

import (
	gossResource "github.com/malnick/goss/resource"
	gossSystem "github.com/malnick/goss/system"
)

type File struct {
	gossResource.File `yaml:",inline"`
}

func (f File) Validate(gs *gossSystem.System) []gossResource.TestResult {
	tr := f.File.Validate(gs)
	return tr
}
