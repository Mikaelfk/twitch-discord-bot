package command

import (
	"net/http"
	"strconv"
	"twitch-discord-bot/constants"
	"twitch-discord-bot/twitchapi"
	"twitch-discord-bot/util"

	"github.com/bwmarrin/discordgo"
)

// Stream stores stream information
type Stream struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	UserLogin   string `json:"user_login"`
	UserName    string `json:"user_name"`
	GameName    string `json:"game_name"`
	Title       string `json:"title"`
	ViewerCount int    `json:"viewer_count"`
	StartedAt   string `json:"started_at"`
	Language    string `json:"language"`
	IsMature    bool   `json:"is_mature"`
}

// AllStreams stores information for several streams
type AllStreams struct {
	Data []struct{ Stream } `json:"data"`
}

const streamCommandWord = "stream"

var maxResults = 3 // Maximum number of shown results in discord

var (

	// define name and description for command
	streamCommand = discordgo.ApplicationCommand{
		Name:        streamCommandWord,
		Description: "Will get the top " + strconv.Itoa(maxResults) + " currently most viewed streams.",
		Options: []*discordgo.ApplicationCommandOption{{

			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "login-name",
			Description: "Get streams by a streamer's name.",
			Required:    false,
		},
			{

				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "game-name",
				Description: "Get top streams that currently plays this game.",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "language",
				Description: "Get currently top streams by their language's ISO 639-1 two letter code.",
				Required:    false,
			},
			{

				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "game-id",
				Description: "Get top streams that currently plays this game.",
				Required:    false,
			},
		},
	}

	// define commandHandler for this command
	streamCommandHandler = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		var content = ""
		URL := constants.URLTwitchStreamInfo + "?first=" + strconv.Itoa(maxResults)

		// Add the optional parameters if any
		for k := 0; k < len(i.Data.Options); k++ {
			if i.Data.Options[k].Name == streamCommand.Options[0].Name {
				URL += "&" + constants.ParaUserLogin + i.Data.Options[k].StringValue()
			} else if i.Data.Options[k].Name == streamCommand.Options[3].Name {
				URL += "&" + constants.ParaGameID + strconv.Itoa(int(i.Data.Options[k].IntValue()))
			} else if i.Data.Options[k].Name == streamCommand.Options[2].Name {
				URL += "&" + constants.ParaLanguage + i.Data.Options[k].StringValue()
			} else if i.Data.Options[k].Name == streamCommand.Options[1].Name {

				games, err := twitchapi.FindGames(i.Data.Options[k].StringValue(), 1, "")
				if err != nil {
					util.DiscordBotResponder(constants.BotUnexpectedErrorMsg, s, i)
					return
				} else if len(games.Data) == 0 {
					util.DiscordBotResponder(constants.BotNoGames, s, i)
					return
				}
				URL += "&" + constants.ParaGameID + games.Data[0].ID // just get the first (and only) element
			} else {
				// This should never really happen...
				util.DiscordBotResponder(constants.BotUnexpectedErrorMsg, s, i)
				return
			}
		}

		var streams AllStreams
		err := util.HandleRequest(URL, http.MethodGet, &streams)
		if err != nil {
			content = constants.BotUnexpectedErrorMsg
		} else if len(streams.Data) == 0 {
			content = constants.BotNoActiveStreamsMsg
		} else {

			var length int
			if len(streams.Data) > maxResults {
				length = maxResults
			} else {
				length = len(streams.Data)
			}

			for i := 0; i < length; i++ {
				content +=
					"\n----------------" +
						"\nStreamer: " + streams.Data[i].UserLogin + // only bother to show the login name
						"\nTitle: " + streams.Data[i].Title +
						"\nGame: " + streams.Data[i].GameName +
						"\nCurrent Viewers: " + strconv.Itoa(streams.Data[i].ViewerCount) +
						"\nStarted at: " + streams.Data[i].StartedAt +
						"\nLanguage: " + streams.Data[i].Language +
						"\nStream: " + constants.URLTwitchStream + streams.Data[i].UserLogin +
						"\n----------------"
			}
		}

		util.DiscordBotResponder(content, s, i)

	}
)

// RegisterStream function for registering command for the bot to serve
func RegisterStream(commands *[]discordgo.ApplicationCommand, commandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	*commands = append(*commands, streamCommand)
	commandHandlers[streamCommandWord] = streamCommandHandler
}
