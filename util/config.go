// Package util provides utility functions
package util

import (
	"encoding/json"
	"os"
)

// Config, stores config structs
var Config configuration

type configuration struct {
	DiscordBotToken string
	DiscordServerID string
	TwitchClientID  string
	TwitchAuthToken string
}

// LoadConfig, loads the config values
func LoadConfig() error {
	file, err := os.Open("config.json")
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		return err
	}

	return nil
}
