// Package util provides utility functions
package util

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Config stores config structs
var Config configuration

type configuration struct {
	DiscordBotToken                  string
	DiscordServerID                  string
	TwitchClientID                   string
	TwitchAuthToken                  string
	TwitchWebhooksSecret             string
	EnableSubscriptionsFunctionality bool
	CallbackURL                      string
	Port                             int
}

// LoadConfig loads the config values
func LoadConfig(path string) error {
	file, err := os.Open(filepath.Join(filepath.Clean(path), "config.json"))
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
