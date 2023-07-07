package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Configuration represents the project configuration.
type Configuration struct {
	EmailNotifications bool   `json:"email_notifications"`
	LogLevel           string `json:"log_level"`
	// Add more configuration properties as needed
}

// LoadConfiguration loads the configuration from a JSON file.
func LoadConfiguration(filename string) (*Configuration, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Error reading configuration file:", err)
		return nil, err
	}

	var config Configuration
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Println("Error unmarshaling configuration:", err)
		return nil, err
	}

	return &config, nil
}

// SaveConfiguration saves the configuration to a JSON file.
func SaveConfiguration(config *Configuration, filename string) error {
	data, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		log.Println("Error marshaling configuration:", err)
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		log.Println("Error writing configuration file:", err)
		return err
	}

	return nil
}
