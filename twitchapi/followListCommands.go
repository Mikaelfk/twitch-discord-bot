// Package twitchapi provides helper functions for the twitch API
package twitchapi

import (
	"errors"
	"log"
	"net/http"
	"twitch-discord-bot/constants"
	"twitch-discord-bot/util"
)

type twitchFollowList struct {
	Total int `json:"total"`
	Data  []struct {
		FromName string `json:"from_name"`
		ToName   string `json:"to_name"`
	} `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

// GetFollowList, gets the follower list of a user
func GetFollowList(username string) ([]string, error) {
	userID, err := util.GetUserId(username)
	if err != nil {
		log.Printf("Error: %v", err)
		return []string{}, err
	}
	var twitchFollowListResponse twitchFollowList
	err = util.HandleRequest(constants.UrlTwitchFollowlist+userID, http.MethodGet, &twitchFollowListResponse)

	if err != nil {
		log.Println("unable to get followers list")
		return nil, err
	}

	if len(twitchFollowListResponse.Data) == 0 {
		log.Println("User does not follow anyone")
		return []string{}, errors.New("user does not follow anyone")
	}
	streamers := []string{}
	for _, s := range twitchFollowListResponse.Data {
		streamers = append(streamers, s.ToName)
	}
	cursor := ""
	for twitchFollowListResponse.Pagination.Cursor != "" {
		cursor = twitchFollowListResponse.Pagination.Cursor
		twitchFollowListResponse = twitchFollowList{}

		err = util.HandleRequest(constants.UrlTwitchFollowlist+userID+"&after="+cursor, http.MethodGet, &twitchFollowListResponse)

		if err != nil {
			log.Println("unable to get followers list by cursor")
			return nil, err
		}

		for _, s := range twitchFollowListResponse.Data {
			streamers = append(streamers, s.ToName)
		}
	}
	return streamers, nil
}
