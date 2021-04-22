package util

import (
	"encoding/json"
	"os"
)

// store config ciles
var Config configuration

type configuration struct {
	DiscordBotToken string
	DiscordServerID string
	TwitchClientID  string
	TwitchAuthToken string
}

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
