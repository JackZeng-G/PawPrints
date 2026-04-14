package tests

import (
	"testing"
	"pawprints-server/internal/bootstrap"
)

func TestBootstrap_App(t *testing.T) {
	app := bootstrap.NewApp()
	defer app.Close()
	if app.DB == nil {
		t.Error("DB should not be nil")
	}
	if app.Router == nil {
		t.Error("Router should not be nil")
	}
}

func TestBootstrap_Close(t *testing.T) {
	app := bootstrap.NewApp()
	err := app.Close()
	if err != nil {
		t.Errorf("Close should not error: %v", err)
	}
}