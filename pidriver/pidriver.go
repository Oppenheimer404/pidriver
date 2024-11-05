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

func start(cfg *config.Config) {

}

func main() {
	// Loading Config
	cfg, err := config.Load() // config.Load() returns (*Config, error)
	diagnose.Error(err)       // diagnose errors

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
	flag.Usage = config.CustomUsageInfo

	// Parse user flags (e.g. -s, --start)
	flag.Parse()
	switch { // Switch for differentiating flags
	case startFlag: // -s or --start
		start(cfg) // start pidriver with current config values
	case resetFlag: // --reset
		err := cfg.Reset()  // reset config value to default values
		diagnose.Error(err) // diagnose errors
	case configField != "": // -c or --config
		if configField == "list" { // special input `list``
			cfg.List() // list all current config values
			break
		}
		if configField == "help" { // special input `help` (not the same as -h)
			cfg.Help() // list configuration help info
			break
		}
		// Evaluate config arguments [field_name new_value] expected
		args := flag.Args()
		if len(args) < 1 { // determine if there is at least two values
			flag.Usage() // print usage information
			err := fmt.Errorf("0[Configuration] config requires both field_name and new_value")
			diagnose.Error(err) // diagnose errors
		}
		err := cfg.Update(configField, strings.Join(args, " ")) // update config value
		diagnose.Error(err)                                     // diagnose errors
		fmt.Printf("Configuration updated: %s = %s\n", configField, strings.Join(args, " "))
	default:
		flag.Usage()
		os.Exit(0)
	}
}
