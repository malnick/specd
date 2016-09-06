package cli

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/malnick/specd/config"
	"github.com/malnick/specd/resources"
)

func RunReport(c config.Config) error {
	s, err := resources.LoadStateYaml(c.StatePath)
	if err != nil {
		log.Error(err)
		return err
	}
	r := resources.CreateResources(s)

	testResults := resources.ValidateResources(r)
	jsonData, _ := json.MarshalIndent(testResults, "", "\t")
	fmt.Println(string(jsonData))
	return nil
}
