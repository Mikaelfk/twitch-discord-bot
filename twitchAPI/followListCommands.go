package twitchAPI

import (
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

func GetFollowList(username string) []string {
	userID, err := util.GetUserId(username)
	if err != nil {
		log.Printf("Error: %v", err)
		return []string{}
	}
	var twitchFollowListResponse twitchFollowList
	util.HandleRequest(constants.UrlTwitchFollowlist+userID, http.MethodGet, &twitchFollowListResponse)

	if len(twitchFollowListResponse.Data) <= 0 {
		log.Println("User does not follow anyone")
		return []string{}
	}
	streamers := []string{}
	for _, s := range twitchFollowListResponse.Data {
		streamers = append(streamers, s.ToName)
	}
	cursor := ""
	for twitchFollowListResponse.Pagination.Cursor != "" {
		cursor = twitchFollowListResponse.Pagination.Cursor
		twitchFollowListResponse = twitchFollowList{}

		util.HandleRequest(constants.UrlTwitchFollowlist+userID+"&after="+cursor, http.MethodGet, &twitchFollowListResponse)
		for _, s := range twitchFollowListResponse.Data {
			streamers = append(streamers, s.ToName)
		}
	}
	return streamers
}
