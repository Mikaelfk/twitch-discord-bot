package util

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	DiscordToken string
	TwitchToken  string
	ServerID     string
}

func LoadConfig(config *Configuration) error {
	file, err := os.Open("config.json")
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return err
	}

	return nil
}
