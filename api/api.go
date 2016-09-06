package api

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/malnick/specd/config"
)

func Start(appConfig config.Config) error {
	port := fmt.Sprintf(":%d", appConfig.FlagAPIPort)
	r := mux.NewRouter()
	r.HandleFunc("/specd/api/v1/", RootHandler)
	r.HandleFunc("/specd/api/v1/run/report/", RunReportHandler).
		Methods("GET", "POST")

	log.Infof("Starging specd API on 0.0.0.0%s", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Error(err)
		return err
	}
	return nil
}
