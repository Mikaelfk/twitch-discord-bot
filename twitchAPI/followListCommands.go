package twitchAPI

import (
	"net/http"
	"twitch-discord-bot/constants"
	"twitch-discord-bot/util"
)

type twitchUserSearch struct {
	Data []struct {
		ID string `json:"id"`
	} `json:"data"`
}

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
	var twitchUserSearchResponse twitchUserSearch
	util.HandleRequest(constants.UrlTwitchUserName+username, http.MethodGet, &twitchUserSearchResponse)
	userID := twitchUserSearchResponse.Data[0].ID

	var twitchFollowListResponse twitchFollowList
	util.HandleRequest(constants.UrlTwitchFollowlist+userID, http.MethodGet, &twitchFollowListResponse)

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
