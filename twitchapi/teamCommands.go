package twitchapi

import (
	"errors"
	"log"
	"strings"
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
	} `json:"data"`
}

// StreamInfo stores the username of streamers
type StreamInfo struct {
	Data []struct {
		UserName string `json:"user_name"`
	} `json:"data"`
}

// TeamExist returns true if a team exists
func TeamExist(name string) (bool, error) {
	var teamInfo teamStruct
	name = strings.Join(strings.Split(name, " "), "-")
	err := util.HandleRequest(constants.URLTwitchGetTeams+name, "GET", &teamInfo)
	if err != nil {
		log.Fatal(err)
		return false, errors.New("no team found")
	}
	if teamInfo.Data == nil {
		return false, nil
	}

	return true, nil
}

// GetAllTeamMembers gets all members of a twitch team
func GetAllTeamMembers(name string) ([]string, error) {
	var teamMembers []string
	var teamInfo teamStruct

	if teamInfo.Data == nil {
		return nil, errors.New("no team found")
	}
	name = strings.Join(strings.Split(name, " "), "-")
	err := util.HandleRequest(constants.URLTwitchGetTeams+name, "GET", &teamInfo)

	if err != nil {
		log.Fatal(err)
		return nil, errors.New("no team found")
	}

	for _, member := range teamInfo.Data[0].Users {
		teamMembers = append(teamMembers, member.UserName)
	}

	return teamMembers, nil
}

// GetLiveTeamMembers gets team members that are live
func GetLiveTeamMembers(name string) ([]string, error) {
	var liveTeamMembers []string
	var teamInfo teamStruct
	name = strings.Join(strings.Split(name, " "), "-")
	err := util.HandleRequest(constants.URLTwitchGetTeams+name, "GET", &teamInfo)

	if teamInfo.Data == nil {
		return nil, errors.New("no team found")
	}

	if err != nil {
		log.Fatal(err)
		return nil, errors.New("no team found")
	}
	bigRequest := constants.URLTwitchStreamInfo + "?"

	for _, member := range teamInfo.Data[0].Users {
		bigRequest += "user_id=" + member.UserID + "&"
	}

	bigRequest = strings.TrimSuffix(bigRequest, "&")

	var streamInfo StreamInfo
	err = util.HandleRequest(bigRequest, "GET", &streamInfo)
	if err != nil {
		log.Println("unable to retrieve streamer info")
		return nil, err
	}

	for _, liveMember := range streamInfo.Data {
		liveTeamMembers = append(liveTeamMembers, liveMember.UserName)
	}

	return liveTeamMembers, nil
}
