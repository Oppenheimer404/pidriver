package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/oppenheimer404/pidriver/pidriver/config"
	"github.com/oppenheimer404/pidriver/pidriver/gps"
	"github.com/oppenheimer404/pidriver/pidriver/status"
)

// customUsage defines a custom help page for command-line flags.
func customUsage() {
	fmt.Fprintf(os.Stderr, `Usage:
	pidriver 
	--start    Start pidriver with current configuration
	-s         Shortcut for starting the process
	--reset    Reset config to default values
	--config   Modify config (e.g., --config list)
	-c         Modify config (shorthand for --config)
`)
}

// logFatal logs and exits on error.
func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// clearScreen clears the terminal screen based on the OS.
func clearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	logFatal(err)
}


// printBanner displays the application's banner, version, and author.
func printBanner(cfg *config.Config) {
	banner, _ := (*cfg)["Banner"].(string)
	appName, _ := (*cfg)["AppName"].(string)
	version, _ := (*cfg)["Version"].(string)
	author, _ := (*cfg)["Author"].(string)

	clearScreen()
	fmt.Printf("%s\n[%s] v%s by %s\n", banner, appName, version, author)
}

// updateConfig handles updating the configuration based on field name and value.
func updateConfig(cfg *config.Config, fieldName, newValue string) error {
	switch fieldName {
	case "list":
		cfg.List()
	case "AppName", "Version", "Banner":
		return cfg.Update(fieldName, newValue)
	default:
		return fmt.Errorf("unknown configuration field: %s", fieldName)
	}
	return nil
}

// start verifies device statuses and fetches GPS data.
func start(cfg *config.Config) {
	printBanner(cfg)

	deviceStatus, err := status.Verify(status.ALL)
	logFatal(err)

	if deviceStatus {
		fmt.Println("All devices working!")
	} else {
		fmt.Println("Something broke")
	}

	gpsLocation, err := gps.Request(gps.CURRENT)
	logFatal(err)
	fmt.Println(gpsLocation)
}

func main() {
	cfg, err := config.Load()
	logFatal(err)

	var startFlag, resetFlag bool
	var configField string

	// Flags for starting, resetting, and modifying config
	flag.BoolVar(&startFlag, "start", false, "Start the process")
	flag.BoolVar(&startFlag, "s", false, "Shortcut for starting the process")
	flag.BoolVar(&resetFlag, "reset", false, "Reset config to default values")
	flag.StringVar(&configField, "config", "", "Modify config (e.g., -c AppName newname)")
	flag.StringVar(&configField, "c", "", "Modify config (shorthand for --config)")

	flag.Usage = customUsage
	flag.Parse()

	switch {
	case startFlag:
		start(cfg)
	case resetFlag:
		logFatal(cfg.Reset())
		fmt.Println("Config has been reset successfully!")
	case configField != "":
		if configField == "list" {
			logFatal(updateConfig(cfg, configField, ""))
			break
		}
		args := flag.Args()
		if len(args) < 1 {
			fmt.Println("Please provide both field name and new value.")
			flag.Usage()
			os.Exit(1)
		}
		logFatal(updateConfig(cfg, configField, strings.Join(args, " ")))
		fmt.Printf("Configuration updated: %s = %s\n", configField, strings.Join(args, " "))
	default:
		flag.Usage()
		os.Exit(0)
	}
}
