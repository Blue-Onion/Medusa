package config

import (
	"os"
	"testing"
)

func TestConfigLifecycle(t *testing.T) {
	tempDir := t.TempDir()

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(originalDir)

	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change directory to temp dir: %v", err)
	}

	// 1. config.yaml should not exist initially
	if CheckConfigFile() {
		t.Fatal("CheckConfigFile should return false when config.yaml does not exist")
	}

	// 2. LoadConfig should create default config.yaml
	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	// 3. config.yaml should exist now
	if !CheckConfigFile() {
		t.Fatal("CheckConfigFile should return true after default config.yaml is created")
	}

	// 4. Check cameras
	if len(cfg.Cameras) != 3 {
		t.Fatalf("Expected 3 cameras in default config, got %d", len(cfg.Cameras))
	}

	if cfg.Cameras[0].Name != "cam1" {
		t.Errorf("Expected first camera name 'cam1', got '%s'", cfg.Cameras[0].Name)
	}

	// 5. Check RecordsPath
	if cfg.RecordsPath != "logs" {
		t.Errorf("Expected RecordsPath to be 'logs', got '%s'", cfg.RecordsPath)
	}

	// 6. Check FPS values
	if cfg.Fps["low"] != "2" {
		t.Errorf("Expected fps low to be '2', got '%s'", cfg.Fps["low"])
	}

	if cfg.Fps["medium"] != "5" {
		t.Errorf("Expected fps medium to be '5', got '%s'", cfg.Fps["medium"])
	}

	if cfg.Fps["high"] != "10" {
		t.Errorf("Expected fps high to be '10', got '%s'", cfg.Fps["high"])
	}

	// 7. ReadConfig should also work
	cfg2, err := ReadConfig()
	if err != nil {
		t.Fatalf("ReadConfig failed: %v", err)
	}

	if cfg2.RecordsPath != "logs" {
		t.Errorf("Expected RecordsPath from ReadConfig to be 'logs', got '%s'", cfg2.RecordsPath)
	}
}