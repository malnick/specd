package resources

import (
	"io/ioutil"
	"os"
	"testing"

	"gopkg.in/yaml.v2"

	gossResource "github.com/malnick/goss/resource"
)

var (
	testPkg = Package{
		Package: gossResource.Package{
			Title:     "testPkg",
			Installed: true,
		},
	}

	testFile = File{
		File: gossResource.File{
			Title:  "/testFile",
			Exists: true,
		},
	}

	testService = Service{
		Service: gossResource.Service{
			Title:   "testService",
			Running: true,
		},
	}

	testResources = []Resource{
		&testPkg,
		&testFile,
		&testService,
	}

	tr = Resources{
		Resources: testResources,
	}
)

func TestValidateResources(t *testing.T) {
	testResults := ValidateResources(tr)
	for _, result := range testResults {
		switch result.Title {
		case "testPkg":
			switch result.Property {
			case "installed":
				if result.Found[0] != "false" {
					t.Error("Expected testPkg installed property to be false, got", result.Found[0])
				}
			}

		case "testFile":
			switch result.Property {
			case "exists":
				if result.Found[0] != "false" {
					t.Error("Expected /testFile exists property to be false, got", result.Found[0])
				}
			}

		case "testService":
			switch result.Property {
			case "enabled":
				if result.Found[0] != "false" {
					t.Error("Expected testService enabled property to be false, got", result.Found[0])

				}
			case "running":
				if result.Found != nil {
					t.Error("Expected testService running property to be nil, got", result.Found[0])

				}
			}
		}
	}
}

var mockStateYaml = map[string]interface{}{
	"packages": map[string]interface{}{
		"foo": map[string]interface{}{
			"installed": true,
		},
	},
	"files": map[string]interface{}{
		"/bar.txt": map[string]interface{}{
			"exists": true,
		},
	},
	"services": map[string]interface{}{
		"baz": map[string]interface{}{
			"running": true,
		},
	},
}

func TestCreateResources(t *testing.T) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "prefix")
	if err != nil {
		t.Error("Could not create temp file:", err.Error())
	}
	defer os.Remove(tmpFile.Name())
	writeme, yErr := yaml.Marshal(mockStateYaml)
	if yErr != nil {
		panic(yErr)
	}
	tmpFile.Write(writeme)

	testState, err := LoadStateYaml(tmpFile.Name())
	if err != nil {
		t.Error("Expected no loading errors, got", err.Error())
	}

	testResources := CreateResources(testState)

	if len(testResources.Resources) != 3 {
		t.Error("Expected 3 resources, got", len(testResources.Resources))
	}
}

var badYaml = "foo"

func TestLoadStateYaml(t *testing.T) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "prefix")
	if err != nil {
		t.Error("Could not create temp file:", err.Error())
	}
	defer os.Remove(tmpFile.Name())
	writeme, yErr := yaml.Marshal(mockStateYaml)
	if yErr != nil {
		panic(yErr)
	}
	tmpFile.Write(writeme)

	testState, err := LoadStateYaml(tmpFile.Name())
	if err != nil {
		t.Error("Expected no loading errors, got", err.Error())
	}

	numSvc := len(testState.Services)
	if numSvc != 1 {
		t.Error("Expected 1 service, got", numSvc)
	}
	if nameSvc, ok := testState.Services["baz"]; !ok {
		t.Error("Expected key name baz, got", ok)
	} else {
		if !nameSvc.Running {
			t.Error("Expected baz to be running, got", nameSvc.Running)
		}
	}
}

func TestBadYamlLoadStateYaml(t *testing.T) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "prefix")
	if err != nil {
		t.Error("Could not create temp file:", err.Error())
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Write([]byte(badYaml))

	_, loadErr := LoadStateYaml(tmpFile.Name())
	if loadErr == nil {
		t.Error("Expected loading errors, got", loadErr)
	}
}

func TestBadPathLoadStateYaml(t *testing.T) {
	_, pathErr := LoadStateYaml("/foo")
	if pathErr == nil {
		t.Error("Expected path error, got", pathErr)
	}
}
