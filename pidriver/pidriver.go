package main

// Imports
import (
	// default imports
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"time"

	// package imports
	"github.com/oppenheimer404/pidriver/pidriver/bluetooth"
	"github.com/oppenheimer404/pidriver/pidriver/config"
	"github.com/oppenheimer404/pidriver/pidriver/gps"
	"github.com/oppenheimer404/pidriver/pidriver/logging"
	"github.com/oppenheimer404/pidriver/pidriver/wifi"
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

// logFatal logs info and exits on error.
func logFatal(err error) {
	if err != nil {
		logging.Error(err, "Fatal Error")
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

// printBanner displays the banner containing basic info
func printBanner(cfg *config.Config) {
	// Set variables using config values
	banner, _ := (*cfg)[config.BANNER].(string)
	appName, _ := (*cfg)[config.APP_NAME].(string)
	version, _ := (*cfg)[config.VERSION].(string)
	author, _ := (*cfg)[config.AUTHOR].(string)

	// Clear screen before printing
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
		"WiFi":      (*cfg)[config.WIFI_ACTIVE].(bool),
		"Bluetooth": (*cfg)[config.BT_ACTIVE].(bool),
		"GPS":       (*cfg)[config.GPS_ACTIVE].(bool),
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
	printDeviceStatus(cfg)
	// Verify all devices are connected and in working order

	// Create channels for device results [gps, wifi, bluetooth]
	gpsResults := make(chan map[string]interface{})
	wifiResults := make(chan map[string]interface{})
	bluetoothResults := make(chan map[string]interface{})

	// Map to hold latest gpsResults
	var latestLocation map[string]interface{}
	var mutex sync.Mutex

	// Set scan rates based on config values
	gpsRate := time.Duration((*cfg)[config.GPS_RATE].(float64))
	wifiRate := time.Duration((*cfg)[config.WIFI_RATE].(float64))
	bluetoothRate := time.Duration((*cfg)[config.BT_RATE].(float64))

	// Start A, B, C scans as goroutines
	go gps.StartScan(gpsRate*time.Millisecond, gpsResults)
	go wifi.StartScan(wifiRate*time.Millisecond, wifiResults)
	go bluetooth.StartScan(bluetoothRate*time.Millisecond, bluetoothResults)

	// Listen for C results and update the latest result
	go func() {
		for result := range gpsResults {
			mutex.Lock()
			latestLocation = result
			mutex.Unlock()
		}
	}()

	// Main loop: associate A and B results with the latest C result
	for {
		select {
		case wifiData := <-wifiResults:
			mutex.Lock()
			logging.Default(wifiData, latestLocation)
			mutex.Unlock()

		case bluetoothData := <-bluetoothResults:
			mutex.Lock()
			logging.Default(bluetoothData, latestLocation)
			mutex.Unlock()
		}
	}
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
