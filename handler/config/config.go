package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Camera struct {
	Name   string `yaml:"name"`
	Source string `yaml:"source"`
}

type Config struct {
	Cameras     []Camera          `yaml:"cameras"`
	RecordsPath string            `yaml:"recordsPath"`
	Fps         map[string]string `yaml:"fps"`
}
type Event struct {
	Camera     string
	Time       float64
	Event      string
	Confidence float64
}

func CreateDefaultConfig() error {
	defaultConfig := `cameras:
  - name: cam1
    source: "0"
  - name: cam2
    source: "video.mp4"
  - name: cam3
    source: "rtsp://admin:pass@192.168.1.20:554/stream"

recordsPath: "logs"

fps:
  low: "2"
  medium: "5"
  high: "10"
`
	return os.WriteFile("config.yaml", []byte(defaultConfig), 0644)
}
func CheckConfigFile() bool {
	_, err := os.Stat("config.yaml")
	if err != nil {
		return false
	}
	return true
}
func ReadConfig() (*Config, error) {
	cfg := &Config{}
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
func LoadConfig() (*Config, error) {

	if !CheckConfigFile() {
		err := CreateDefaultConfig()
		if err != nil {
			return nil, err
		}
	}
	cameras, err := ReadConfig()
	if err != nil {

		return nil, err
	}

	return cameras, nil
}
func ShowConfig() {
	cfg, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("========== Medusa Config ==========")

	fmt.Println("\nRecords Path:", cfg.RecordsPath)

	fmt.Println("\nCameras:")
	for i, camera := range cfg.Cameras {
		fmt.Printf("  %d) %-10s -> %s\n", i+1, camera.Name, camera.Source)
	}

	fmt.Println("\nFPS Presets:")
	for key, val := range cfg.Fps {
		fmt.Printf("  %-7s : %s\n", key, val)
	}

	fmt.Println("\n===================================")
}
