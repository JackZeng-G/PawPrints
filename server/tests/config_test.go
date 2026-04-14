package tests

import (
	"testing"
	"pawprints-server/internal/config"
)

func TestConfig_Load(t *testing.T) {
	cfg := config.Load()
	if cfg.Server.Port == 0 {
		t.Error("Server port should not be zero")
	}
	if cfg.Database.Path == "" {
		t.Error("Database path should not be empty")
	}
}

func TestConfig_Defaults(t *testing.T) {
	cfg := config.Load()
	if cfg.Server.Port != 8080 {
		t.Errorf("Default port should be 8080, got %d", cfg.Server.Port)
	}
	if cfg.Server.UploadDir != "./uploads" {
		t.Errorf("Default upload dir should be ./uploads, got %s", cfg.Server.UploadDir)
	}
}
