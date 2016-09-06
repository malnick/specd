package cli

import (
	"testing"

	"github.com/malnick/specd/config"
)

func TestRunReport(t *testing.T) {
	appConfig := config.Configuration()
	if err := RunReport(appConfig); err == nil {
		t.Error("Expected errors about missing state.yaml, got", err.Error())
	}
}
