package util

import (
	"errors"
	"log"
	"net/http"
	"twitch-discord-bot/constants"
	"twitch-discord-bot/db"
)

// TwitchChannels is a collection of multiple Channel(s) and is the response form from GET "https://api.twitch.tv/helix/search/channels?query=pokimane"
type TwitchChannels struct {
	Data []struct{ Channel } `json:"data"`
}

// Channel is a struct over a specific channel in the response from GET "https://api.twitch.tv/helix/search/channels?query=pokimane"
type Channel struct {
	Lang        string `json:"broadcaster_language"`
	DisplayName string `json:"display_name"`
	LoginName   string `json:"broadcaster_login"`
	IsLive      bool   `json:"is_live"`
	Title       string `json:"title"`
	StreamID    string `json:"id"`
	GameName    string `json:"game_name"`
	GameID      string `json:"game_id"`
	Thumbnail   string `json:"thumbnail_url"`
	StartedAt   string `json:"started_at"`
}

type twitchUserSearch struct {
	Data []struct {
		ID string `json:"id"`
	} `json:"data"`
}

// SearchByName takes in a searchName string and a TwitchChannels struct and returns the channel with DisplayName or LoginName like the search.
// If nothing was found, return error and empty Channel
func SearchByName(searchName string, channels TwitchChannels) (Channel, error) {
	for i := 0; i < len(channels.Data); i++ {

		// If the login-name or the displayed name of channel/streamer is equal search...
		if channels.Data[i].Channel.LoginName == searchName || channels.Data[i].Channel.DisplayName == searchName {
			return channels.Data[i].Channel, nil
		}
	}
	return Channel{}, errors.New("no search result found")
}

// GetUserID takes a username string and return the user id as a string.
// If the user does not exist, an error is returned
func GetUserId(username string) (string, error) {
	var twitchUserSearchResponse twitchUserSearch
	err := HandleRequest(constants.UrlTwitchUserName+username, http.MethodGet, &twitchUserSearchResponse)

	if err != nil {
		log.Println("unable to get user id")
		return "", err
	}

	if len(twitchUserSearchResponse.Data) == 0 {
		err := errors.New("user does not exist")
		log.Printf("Error: %v", err.Error())
		return "", err
	}
	userID := twitchUserSearchResponse.Data[0].ID
	return userID, nil
}

// ChannelIdExists checks if a channel id exists for a particular streamer
func ChannelIdExists(streamerID string, channelID string) bool {
	channelIDs, err := db.GetChannelIdsByStreamerID(streamerID)
	if err != nil {
		log.Println("streamer id not found")
		return false
	}

	for _, v := range channelIDs {
		if channelID == v {
			log.Println("channel id found")
			return true
		}
	}
	return false
}
