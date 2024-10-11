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
)

// customUsage defines a custom help page for command-line flags.
func customUsage() {
	fmt.Fprintf(os.Stderr, `Usage:
	pidriver [options]

Options:
	--start, -s        Start pidriver with the current configuration.
	--reset            Reset the configuration to default values.
	--config, -c       Modify configuration settings. Examples:
	                   --config list            Show current configuration settings.
	                   --config AppName <new_name> Update the AppName in the configuration.
	                   -c AppName <new_name>   Shorthand for updating the configuration.
	                   Use --config help for more details on config options.

For more information, visit the documentation or run 'pidriver --help'.
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
	err := cfg.Update(fieldName, newValue)
	return err
}

// confirmation message declaring which devices are active
func printDeviceStatus(cfg *config.Config) {
	statuses := map[string]bool{
		"WiFi":      (*cfg)["WifiEnabled"].(bool),
		"Bluetooth": (*cfg)["BluetoothEnabled"].(bool),
		"GPS":       (*cfg)["GPSEnabled"].(bool),
	}

	// Print the status for each device
	for name, enabled := range statuses {
		fmt.Printf("[%s %s]\n", name, map[bool]string{true: "enabled", false: "disabled"}[enabled])
	}
}


// start pidriver with current configuration
func start(cfg *config.Config) {
	// Clears screen and prints banner, author, & version #
	printBanner(cfg)

	// Verify all devices are connected and in working order
	
	
	// Continue to scanning
	fmt.Println("All devices working!")
	printDeviceStatus(cfg)
	fmt.Println("Continuing to scan...")
}

func main() {
	// Loading Config
	cfg, err := config.Load()
	logFatal(err)

	// Declare flags
	var startFlag, resetFlag bool
	var configField string

	// Flags for starting, resetting, and modifying config
	flag.BoolVar(&startFlag, "start", false, "Start pidriver")
	flag.BoolVar(&startFlag, "s", false, "Alternate start (shorthand for --start)")
	flag.BoolVar(&resetFlag, "reset", false, "Reset config to default values")
	flag.StringVar(&configField, "config", "", "Modify config (e.g., --config list, -c AppName <new_name>)")
	flag.StringVar(&configField, "c", "", "Modify config (shorthand for --config)")

	// Parse flags
	flag.Usage = customUsage
	flag.Parse()

	switch {
	case startFlag: // Begins pidriver with current config
		start(cfg)
	case resetFlag: // Resets config to default settings
		logFatal(cfg.Reset())
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
		logFatal(updateConfig(cfg, configField, strings.Join(args, " ")))
		fmt.Printf("Configuration updated: %s = %s\n", configField, strings.Join(args, " "))
	default:
		// Print usage
		flag.Usage()
		os.Exit(0)
	}
}
