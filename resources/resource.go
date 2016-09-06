package resources

import (
	"io/ioutil"

	log "github.com/Sirupsen/logrus"

	"github.com/codegangsta/cli"
	gossResource "github.com/malnick/goss/resource"
	gossSystem "github.com/malnick/goss/system"
	"gopkg.in/yaml.v2"
)

type Resource interface {
	Validate(*gossSystem.System) []gossResource.TestResult
}

type Resources struct {
	Resources []Resource
}

type State struct {
	Packages map[string]Package `yaml:"packages" json:"packages"`
	Files    map[string]File    `yaml:"files" json:"files"`
	Services map[string]Service `yaml:"services" json:"services"`
	Memory   map[string]Memory  `yaml:"memory" json:"memory"`
}

func LoadStateYaml(statePath string) (*State, error) {
	var s *State

	fileBytes, err := ioutil.ReadFile(statePath)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if err := yaml.Unmarshal(fileBytes, &s); err != nil {
		log.Error(err)
		return nil, err
	}

	return s, nil
}

func ValidateResources(r Resources) []gossResource.TestResult {
	results := []gossResource.TestResult{}
	cliWrapper := cli.NewApp()
	cliWrapper.Action = func(c *cli.Context) {
		sys := gossSystem.New(c)
		for _, r := range r.Resources {
			tr := r.Validate(sys)
			results = append(results, tr...)
		}
	}
	cliWrapper.Run([]string{""})

	return results
}

func CreateResources(s *State) Resources {
	var r = Resources{}
	if len(s.Files) > 0 {
		for resourceName, _ := range s.Files {
			resourceData := s.Files[resourceName]
			log.WithFields(log.Fields{
				"Name": resourceName,
			}).Debug("File")
			resourceData.SetID(resourceName)
			r.Resources = append(r.Resources, &resourceData)
		}
	}

	if len(s.Services) > 0 {
		for resourceName, _ := range s.Services {
			resourceData := s.Services[resourceName]
			log.WithFields(log.Fields{
				"Name": resourceName,
			}).Debug("Service")
			resourceData.SetID(resourceName)
			r.Resources = append(r.Resources, &resourceData)
		}
	}

	if len(s.Packages) > 0 {
		for resourceName, _ := range s.Packages {
			resourceData := s.Packages[resourceName]
			log.WithFields(log.Fields{
				"Name": resourceName,
			}).Debug("Package")
			resourceData.SetID(resourceName)
			r.Resources = append(r.Resources, &resourceData)
		}
	}

	if len(s.Memory) > 0 {
		for resourceName, _ := range s.Memory {
			resourceData := s.Memory[resourceName]
			log.WithFields(log.Fields{
				"Name": resourceName,
			}).Debug("Memory")
			resourceData.SetID(resourceName)
			r.Resources = append(r.Resources, &resourceData)
		}
	}
	return r
}
