package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"

	"github.com/malnick/specd/config"
	rLib "github.com/malnick/specd/resources"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	log.Debugf("%s /specd/api/v1/", r.Method)
	appConfig := config.Configuration()
	versionRevision := map[string]string{
		"version":  appConfig.Version,
		"revision": appConfig.Revision,
	}
	json, err := json.Marshal(versionRevision)
	if err != nil {
		log.Error(err)
		return
	}
	w.Write(json)
}

func RunReportHandler(w http.ResponseWriter, r *http.Request) {
	log.Debugf("%s /specd/api/v1/run/report/", r.Method)
	c := config.Configuration()

	s, err := getStateFromFileOrRequestData(r, c)
	if err != nil {
		handleInternalServerError(err, w)
		return
	}

	resources := rLib.CreateResources(s)
	testResults := rLib.ValidateResources(resources)
	jsonData, _ := json.MarshalIndent(testResults, "", "\t")
	log.Debug(string(jsonData))
	w.Write(jsonData)

}

func getStateFromFileOrRequestData(r *http.Request, c config.Config) (*rLib.State, error) {
	if r.Method == "GET" {
		return rLib.LoadStateYaml(c.StatePath)
	}
	if r.Method == "POST" {
		s := rLib.State{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&s)
		return &s, err
	}
	return &rLib.State{}, errors.New(fmt.Sprintf("Method not supported: %s", r.Method))
}

func handleInternalServerError(err error, w http.ResponseWriter) {
	log.Error(err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("{}"))
	return
}
