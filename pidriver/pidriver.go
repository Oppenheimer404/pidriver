package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	// package imports
	"github.com/oppenheimer404/pidriver/pidriver/config"
	"github.com/oppenheimer404/pidriver/pidriver/diagnose"
)

func main() {
	// Loading Config
	cfg, err := config.Load() // config.Load() returns (*Config, error)
	diagnose.Error(err)       // diagnose any errors

	// Declare all flags
	var startFlag, resetFlag bool // boolean flags
	var configField string        // string flags

	// Flags for starting, resetting, and modifying config
	flag.BoolVar(&startFlag, "start", false, "Start pidriver")
	flag.BoolVar(&startFlag, "s", false, "Start pidriver (shorthand for --start)")
	flag.BoolVar(&resetFlag, "reset", false, "Reset config to default values")
	flag.StringVar(&configField, "config", "", "Modify config (e.g., --config list, -c AppName <new_name>)")
	flag.StringVar(&configField, "c", "", "Modify config (shorthand for --config)")

    // Set default usage info to custom message (-h or --help)
	flag.Usage = printUsageInfo

	// Parse flags
	flag.Parse()

	switch {
	case startFlag: // Begins pidriver with current config
		start(cfg)
	case resetFlag: // Resets config to default settings
		err := cfg.Reset()
		diagnose.Error(err)
		fmt.Println("Config has been reset successfully!")
	case configField != "": // Edit config via cli
		// Checks for `-c list` and lists config if run
		if configField == "list" {
			cfg.List()
			break
		}
		// Checks for `-c help` and lists config help info
		if configField == "help" {
			cfg.Help()
			break
		}
		// Ensures both field name and new value are provided
		args := flag.Args()
		if len(args) < 1 {
			fmt.Println("Please provide both field name and new value.")
			flag.Usage()
			os.Exit(1)
		}
		// Update config with new value
		err := updateConfig(cfg, configField, strings.Join(args, " "))
		diagnose.Error(err)
		fmt.Printf("Configuration updated: %s = %s\n", configField, strings.Join(args, " "))
	default:
		// Print usage
		flag.Usage()
		os.Exit(0)
	}
}
