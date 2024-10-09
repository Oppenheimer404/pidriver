package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/oppenheimer404/pidriver/pidriver/gps"
	"github.com/oppenheimer404/pidriver/pidriver/status"
	"github.com/oppenheimer404/pidriver/pidriver/config"
	// "github.com/oppenheimer404/pidriver/pidriver/wifi"
)

// Custom help page
func customUsage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "	pidriver \n")
	fmt.Fprintf(os.Stderr, "  --start	Start pidriver with current configuration\n")
	fmt.Fprintf(os.Stderr, "  -s		Shortcut for starting the process\n")
	fmt.Fprintf(os.Stderr, "  --reset	Reset config to default values\n")
}

// Error handling
func checkFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Pseudo "main" function
func start() {
	// Verify everything is working
	deviceStatus, err := status.Verify(status.ALL)
	checkFatal(err)
	if !deviceStatus {
		fmt.Println("Something Broke")
	} else {
		fmt.Println("All devices working!")
	}

	// Fetch GPS location as an example
	gpsLocation, err := gps.Request(gps.CURRENT)
	checkFatal(err)
	fmt.Println(gpsLocation)
	// Additional functionality like wifi.Request() can be added
}

func main() {
	// Load the configuration
	cfg, err := config.Load()
	checkFatal(err)
	fmt.Print(cfg.Banner, "\n")
	fmt.Printf("[%s] v%s by %s\n", cfg.AppName, cfg.Version, cfg.Author)

	// Define the flags
	var startFlagA, startFlagB, resetFlag bool
	flag.BoolVar(&startFlagA, "start", false, "Start the process")
	flag.BoolVar(&startFlagB, "s", false, "Shortcut for starting the process")
	flag.BoolVar(&resetFlag, "reset", false, "Reset config to default values")

	flag.Usage = customUsage // Assign custom usage function

	// Parse the flags
	flag.Parse()

	// Handle the flags
	switch {
	case startFlagA || startFlagB: // Run pidriver
		start()
	case resetFlag: // Reset configuration file
		err = cfg.Reset()
		checkFatal(err)
		fmt.Println("Config has been reset successfully!")
	default:
		// If no valid flags are provided, print custom usage and exit
		flag.Usage()
		os.Exit(1)   // Exit with a non-zero status code
	}
}

