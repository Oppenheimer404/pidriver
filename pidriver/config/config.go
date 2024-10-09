package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// RequestUpdateConfig defines constants for updating configuration fields.
type RequestUpdateConfig int

const (
	APP_NAME RequestUpdateConfig = iota
	VERSION
	BANNER
)

// Config represents the configuration structure.
type Config struct {
	AppName string `json:"appName"`
	Version string `json:"version"`
	Author string `json:"author"`
	Banner  string `json:"banner"`
}

// Default values for the configuration.
var defaultConfig = Config{
	AppName: "pidriver",
	Version: "0.0.1",
	Author: "Oppenheimer404",
	Banner: `
                                                              
@@@@@@@  @@@ @@@@@@@  @@@@@@@  @@@ @@@  @@@ @@@@@@@@ @@@@@@@  
@@!  @@@ @@! @@!  @@@ @@!  @@@ @@! @@!  @@@ @@!      @@!  @@@ 
@!@@!@!  !!@ @!@  !@! @!@!!@!  !!@ @!@  !@! @!!!:!   @!@!!@!  
!!:      !!: !!:  !!! !!: :!!  !!:  !: .:!  !!:      !!: :!!  
 :       :   :: :  :   :   : : :      ::    : :: ::   :   : : 
                                                              
`,
}

// Load loads the configuration from config.json, or creates it if not found.
func Load() (*Config, error) {
	file, err := os.Open("config.json")
	if os.IsNotExist(err) {
		return createDefaultConfig()
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

// createDefaultConfig initializes the default configuration and writes it to a file.
func createDefaultConfig() (*Config, error) {
	fmt.Println("Config file not found, creating a new one...")

	if err := writeConfigToFile(&defaultConfig); err != nil {
		return nil, err
	}

	return &defaultConfig, nil
}

// Update changes a specified value in the configuration.
func (c *Config) Update(key RequestUpdateConfig, value string) error {
	switch key {
	case APP_NAME:
		c.AppName = value
	case VERSION:
		c.Version = value
	case BANNER:
		c.Banner = value
	default:
		return fmt.Errorf("invalid key: %d", key)
	}

	return writeConfigToFile(c)
}

// Reset restores the configuration to its default values.
func (c *Config) Reset() error {
	fmt.Println("Resetting to default config...")
	return writeConfigToFile(&defaultConfig)
}

// writeConfigToFile writes the provided configuration to the config.json file.
func writeConfigToFile(config *Config) error {
	file, err := os.Create("config.json")
	if err != nil {
		return fmt.Errorf("failed to open config file for writing: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(config); err != nil {
		return fmt.Errorf("failed to save config: %v", err)
	}

	return nil
}
