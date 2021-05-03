package twitchAPI

import (
	"errors"
	"log"
	"twitch-discord-bot/constants"
	"twitch-discord-bot/util"
)

type teamStruct struct {
	Data []struct {
		Users []struct {
			UserID    string `json:"user_id"`
			UserName  string `json:"user_name"`
			UserLogin string `json:"user_login"`
		} `json:"users"`
		BackgroundImageURL interface{} `json:"background_image_url"`
		Banner             interface{} `json:"banner"`
		Info               string      `json:"info"`
		ThumbnailURL       string      `json:"thumbnail_url"`
		TeamName           string      `json:"team_name"`
		TeamDisplayName    string      `json:"team_display_name"`
		ID                 string      `json:"id"`
	} `json:"data"`
}

// Gets the team name
func GetTeamName(name string) (string, error) {

	var teamName string
	var teamStruct teamStruct

	err := util.HandleRequest(constants.UrlTwitchGetTeams+name, "GET", &teamStruct)

	if err != nil {
		log.Fatal(err)
		return "", errors.New("no team found")
	}

	teamName = teamStruct.Data[0].TeamName

	return teamName, nil

}

// Gets all members of a twitch team
func GetAllTeamMembers(name string) ([]string, error) {
	var teamMembers []string
	var teamStruct teamStruct

	err := util.HandleRequest(constants.UrlTwitchGetTeams+name, "GET", &teamStruct)

	if err != nil {
		log.Fatal(err)
		return nil, errors.New("no team found")
	}

	for _, member := range teamStruct.Data[0].Users {
		teamMembers = append(teamMembers, member.UserName)
	}

	return teamMembers, nil
}

// Gets team members that are live
func GetLiveTeamMembers(name string) ([]string, error) {

	var channelStruct util.Channel
	var members []string
	var liveMembers []string

	members, _ = GetAllTeamMembers(name)
}
