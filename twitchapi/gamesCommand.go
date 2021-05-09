package twitchapi

import (
	"strconv"
	"strings"
	"twitch-discord-bot/constants"
	"twitch-discord-bot/util"
)

type game struct {
	ArtURL string `json:"box_art_url"`
	ID     string `json:"id"`
	Name   string `json:"name"`
}

type gamesData struct {
	Data []game `json:"data"`
}

// FindGames finds the 'first' games that is somewhat similar to gameName
func FindGames(gameName string, first int) (gamesData, error) {
	// Replaces possible spaces with "-" (dashes) before calling the handleRequest method
	gameName = strings.Join(strings.Split(gameName, " "), "-")
	URL := constants.URLTwitchGames + gameName + "&first=" + strconv.Itoa(first)

	var data gamesData
	err := util.HandleRequest(URL, "GET", &data)

	return data, err
}
