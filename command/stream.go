package command

import (
	"github.com/bwmarrin/discordgo"
	"net/http"
	"strconv"
	"twitch-discord-bot/constants"
	"twitch-discord-bot/util"
)


type Stream struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	UserLogin string `json:"user_login"`
	UserName  string `json:"user_name"`
	GameName  string `json:"game_name"`
	Title     string `json:"title"`
	ViewerCount int `json:"viewer_count"`
	StartedAt string `json:"started_at"`
	Language  string `json:"language"`
	IsMature  bool `json:"is_mature"`
}
type AllStreams struct {
	Data []struct {Stream} `json:"data"`
}

var streamCommandWord = "stream"
var maxResults = 3 // Maximum number of shown results in discord

var (


	// define name and description for command
	streamCommand = discordgo.ApplicationCommand{
		Name:        streamCommandWord,
		Description: "Will get the top "+strconv.Itoa(maxResults)+" currently most viewed streams.",
		Options: []*discordgo.ApplicationCommandOption{{

			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "login-name",
			Description: "Get streams by a streamer's name.",
			Required:    false,
		},
		{

			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "game-id",
			Description: "Get top streams that currently plays this game.",
			Required:    false,
		},
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "language",
			Description: "Get currently top streams by their language's ISO 639-1 two letter code.",
			Required:    false,
		},
		},
	}

	// define commandHandler for this command
	streamCommandHandler = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		var content = ""
		URL := constants.UrlTwitchStreamInfo+"?first="+strconv.Itoa(maxResults)

		// Add the optional parameters if any
		for k:=0; k<len(i.Data.Options); k++{
			if i.Data.Options[k].Name==streamCommand.Options[0].Name{
				URL += "&"+ constants.ParaUserLogin + i.Data.Options[k].StringValue()
			} else if i.Data.Options[k].Name==streamCommand.Options[1].Name{
				URL += "&"+ constants.ParaGameId + i.Data.Options[k].StringValue()
			} else if i.Data.Options[k].Name==streamCommand.Options[2].Name{
				URL += "&"+ constants.ParaLanguage + i.Data.Options[k].StringValue()
			} else {
				// This should never really happen...
				util.DiscordBotResponder("I encountered an unexpected error :/",s,i)
				return
			}
		}

		var streams AllStreams
		err := util.HandleRequest(URL, http.MethodGet, &streams)
		if err!= nil {
			content = "Something went wrong..."
		} else if len(streams.Data)<=0 {
			content = "I couldn't find any active streams... :'("
		} else {

			var length int
			if len(streams.Data)>maxResults{
				length = maxResults
			} else {
				length = len(streams.Data)
			}

			for i:=0; i<length; i++ {
				content +=
					"\n----------------"+
					"\nStreamer: " + streams.Data[i].UserLogin + // only bother to show the login name
					"\nTitle: " + streams.Data[i].Title +
					"\nGame: " + streams.Data[i].GameName +
					"\nCurrent Viewers: " + strconv.Itoa(streams.Data[i].ViewerCount) +
					"\nStarted at: " + streams.Data[i].StartedAt +
					"\nLanguage: " + streams.Data[i].Language +
					"\nStream: " + constants.UrlTwitchStream + streams.Data[i].UserLogin +
					"\n----------------"
			}
		}

		util.DiscordBotResponder(content, s, i)

	}
)

//RegisterStream function for registering command for the bot to serve
func RegisterStream(commands *[]discordgo.ApplicationCommand, commandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	*commands = append(*commands, streamCommand)
	commandHandlers[streamCommandWord] = streamCommandHandler
}


