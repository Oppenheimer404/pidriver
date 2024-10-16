package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

const (
	DEFAULT_CONF = "config.json"

	APP_NAME    = "AppName"
	AUTHOR      = "Author"
	BANNER      = "Banner"
	VERSION     = "Version"
	WIFI_ACTIVE = "WifiEnabled"
	WIFI_RATE   = "WifiRate"
	BT_ACTIVE   = "BluetoothEnabled"
	BT_RATE     = "BluetoothRate"
	GPS_ACTIVE  = "GPSEnabled"
	GPS_RATE    = "GPSRate"
)

// Config represents a dynamic configuration map
type Config map[string]interface{}

// Default configuration values
var defaultConfig = Config{
	"AppName": "pidriver",
	"Author":  "Oppenheimer404",
	"Banner": `
	
	@@@@@@@  @@@ @@@@@@@  @@@@@@@  @@@ @@@  @@@ @@@@@@@@ @@@@@@@  
	@@!  @@@ @@! @@!  @@@ @@!  @@@ @@! @@!  @@@ @@!      @@!  @@@ 
	@!@@!@!  !!@ @!@  !@! @!@!!@!  !!@ @!@  !@! @!!!:!   @!@!!@!  
	!!:      !!: !!:  !!! !!: :!!  !!:  !: .:!  !!:      !!: :!!  
	:       :   :: :  :   : : :      ::    : :: ::   :   : : 
	
	`,
	"Version":          "0.0.1",
	"WifiEnabled":      true,
	"WifiRate":         1000,
	"BluetoothEnabled": true,
	"BluetoothRate":    1000,
	"GPSEnabled":       true,
	"GPSRate":          4000,
	"GPSTimeout":       60,
	"VerboseMode":      true,
}

// createDefaultConfig initializes the default configuration and writes it to a file.
func createDefaultConfig(filePath string) (*Config, error) {
	fmt.Println("Creating new configuration file with default values...")

	if err := writeConfigToFile(&defaultConfig, filePath); err != nil {
		return nil, err
	}

	return &defaultConfig, nil
}

// writeConfigToFile writes the configuration to a specified file.
func writeConfigToFile(config *Config, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to open config file %s for writing: %v", filePath, err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ") // Pretty print
	if err := encoder.Encode(config); err != nil {
		return fmt.Errorf("failed to save config to file %s: %v", filePath, err)
	}

	return nil
}

// List lists the current configuration
func (c *Config) Help() {
	fmt.Fprintf(os.Stderr, `Configuration Options:
	
General Configuration:
	- AppName <new_name>         Set the name of the application.
	- Version <new_version>      Update the application version.
	- Banner <new_banner>        Update the banner displayed at startup.

WiFi Configuration:
	- WifiEnabled <true/false>   Enable or disable WiFi scanning.
	- WifiRate <rate>            Set the scanning rate for WiFi networks (in milliseconds).

Bluetooth Configuration:
	- BluetoothEnabled <true/false> Enable or disable Bluetooth scanning.
	- BluetoothRate <rate>       Set the scanning rate for Bluetooth devices (in milliseconds).

GPS Configuration:
	- GPSEnabled <true/false>    Enable or disable GPS tracking.
	- GPSRate <rate>             Set the location update rate for GPS (in milliseconds).
	- GPSTimeout <timeout>       Set the maximum wait time for a GPS fix (in seconds).

Misc. Configuration:
	- VerboseMode <true/false>		Enable or disable logging output to console

To modify a setting, use the following syntax:
	--config <Option> <new_value>
Examples:
	--config RateWifi 500
	--config WifiEnabled true
	--config GPSTimeout 30

For more details, visit the documentation or run 'pidriver --help'.
`)
}

// parseInt converts interface to int if possible.
func parseInt(value interface{}) (int, error) {
	switch v := value.(type) {
	case string:
		return strconv.Atoi(v)
	case float64:
		return int(v), nil
	case int:
		return v, nil
	default:
		return 0, fmt.Errorf("cannot convert %v to int", value)
	}
}

// parseFloat converts interface to float64 if possible.
func parseFloat(value interface{}) (float64, error) {
	switch v := value.(type) {
	case string:
		return strconv.ParseFloat(v, 64)
	case float64:
		return v, nil
	case int:
		return float64(v), nil
	default:
		return 0, fmt.Errorf("cannot convert %v to float64", value)
	}
}

// parseBool converts interface to bool if possible.
func parseBool(value interface{}) (bool, error) {
	switch v := value.(type) {
	case string:
		return strconv.ParseBool(v)
	case bool:
		return v, nil
	default:
		return false, fmt.Errorf("cannot convert %v to bool", value)
	}
}

// Update dynamically updates a key in the configuration map.
func (c *Config) Update(key string, value interface{}) error {
	// Check if the key exists in the config
	if existingValue, exists := (*c)[key]; exists {
		switch existingValue.(type) {
		case int:
			// Convert value to int
			newValue, err := parseInt(value)
			if err != nil {
				return fmt.Errorf("invalid value for %s: %v", key, err)
			}
			(*c)[key] = newValue
		case float64:
			// Convert value to float64
			newValue, err := parseFloat(value)
			if err != nil {
				return fmt.Errorf("invalid value for %s: %v", key, err)
			}
			(*c)[key] = newValue
		case bool:
			// Convert value to bool
			newValue, err := parseBool(value)
			if err != nil {
				return fmt.Errorf("invalid value for %s: %v", key, err)
			}
			(*c)[key] = newValue
		case string:
			// Convert value to string
			(*c)[key] = fmt.Sprintf("%v", value)
		default:
			return fmt.Errorf("unsupported type for key: %s", key)
		}
	} else {
		return fmt.Errorf("unknown configuration value: try `-c help`")
	}

	// Write the updated config back to the file
	return writeConfigToFile(c, DEFAULT_CONF)
}

// Reset restores the configuration to its default values.
func (c *Config) Reset() error {
	*c = defaultConfig
	return writeConfigToFile(c, DEFAULT_CONF)
}

// List lists the current configuration
func (c *Config) List() {
	fmt.Println("Current Configuration:")
	for key, value := range *c {
		fmt.Printf("%s: %v\n", key, value)
	}
}

// Load loads the configuration from config.json, or creates it if not found.
func Load() (*Config, error) {
	filePath := DEFAULT_CONF
	file, err := os.Open(filePath)
	if os.IsNotExist(err) {
		return createDefaultConfig(filePath)
	} else if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %v", err)
	}

	return &config, nil
}
