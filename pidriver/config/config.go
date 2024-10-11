package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	DEFAULT_CONF = "config.json"
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
	"Version": "0.0.1",
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

// Update dynamically updates a key in the configuration map.
func (c *Config) Update(key string, value interface{}) error {
	(*c)[key] = value
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
