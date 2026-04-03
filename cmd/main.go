package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	camera "github.com/Blue-Onion/MahilAi/handler/Camera"
	record "github.com/Blue-Onion/MahilAi/handler/Record"
	"github.com/Blue-Onion/MahilAi/handler/config"
	"github.com/Blue-Onion/MahilAi/handler/csv"
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
		fpsMode := flag.NewFlagSet("show-record", flag.ExitOnError)
		
		mode := fpsMode.String("m", "", "mode")
	
		
		fpsMode.Parse(os.Args[1:])
		camera.StartCameraWork(cfg,*mode)

	case "show-config":
		config.ShowConfig()
	case "show-record":

		showRecordCmd := flag.NewFlagSet("show-record", flag.ExitOnError)

		date := showRecordCmd.String("date", "", "Date in format YYYY-MM-DD")
		cam := showRecordCmd.String("cam", "", "Camera name")

		showRecordCmd.Parse(os.Args[2:])
		record.ShowRecord(*date, *cam)
	case "download":
		downloadCsv := flag.NewFlagSet("show-record", flag.ExitOnError)

		date := downloadCsv.String("date", "", "Date in format YYYY-MM-DD")
		cam := downloadCsv.String("cam", "", "Camera name")
		downloadCsv.Parse(os.Args[2:])
		csv.ExportToCsv(*date, *cam)

	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Commands: start, show-config, show-record")
		os.Exit(1)
	}
}
