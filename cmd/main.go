package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	camera "github.com/Blue-Onion/MahilAi/handler/Camera"
	record "github.com/Blue-Onion/MahilAi/handler/Record"
	"github.com/Blue-Onion/MahilAi/handler/config"
)


func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error in connecting to config file")
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: mahilai <command> [options]")
		fmt.Println("Commands: start, show-config, show-record")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {

	case "start":
		log.Print("Config loaded. Starting camera work...")
		camera.StartCameraWork(cfg)

	case "show-config":
		config.ShowConfig()

	case "show-record":

		showRecordCmd := flag.NewFlagSet("show-record", flag.ExitOnError)

		date := showRecordCmd.String("date", "", "Date in format YYYY-MM-DD")
		cam := showRecordCmd.String("cam", "", "Camera name")


		showRecordCmd.Parse(os.Args[2:])

		record.ShowRecord(*date, *cam)

	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Commands: start, show-config, show-record")
		os.Exit(1)
	}
}