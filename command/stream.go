package command

import (
	"github.com/bwmarrin/discordgo"
	"net/http"
	"twitch-discord-bot/constants"
	"twitch-discord-bot/util"
)

//
// Example command - for authentication with twitch API
//


var commandWord = "stream"

var (


	// define name and description for command
	streamCommand = discordgo.ApplicationCommand{
		Name:        commandWord,
		Description: "will get info about stream",
		Options: []*discordgo.ApplicationCommandOption{{

			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "streamer",
			Description: "Get info of particular stream by name or ID",
			Required:    true,
		},
		},
	}

	// define commandHandler for this command
	streamCommandHandler = func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		// Todo: Check if it is an int, if so - assume it is an ID and try to search with it
		//r := []rune(i.Data.Options[0].StringValue())
		var content string
		var err error
		var channels util.TwitchChannels

		// Search by name
		URL := constants.UrlTwitchChannelName + i.Data.Options[0].StringValue()
		err = util.HandleRequest(URL, http.MethodGet, &channels)

		if err!= nil {
			util.DiscordBotResponder("Something went wrong...", s, i)
			return
		}
		var channel util.Channel
		channel, err = util.SearchByName(i.Data.Options[0].StringValue(),channels)

		// there are no channels with this exact name...
		if err != nil {
			if len(channels.Data)>0{
				// If channels.Data is not empty, just return the first result here
				channel = channels.Data[0].Channel
			} else {
				// If channels.Data is empty, return the error
				util.DiscordBotResponder(err.Error(), s, i)
				return
			}
		}

		content = "Broadcaster: " + channel.DisplayName +
			"\nStream-Title: "+channel.Title +
			"\nLanguage: " + channel.Lang +
			"\nGame: " + channel.GameName
		if channel.IsLive {
			content += "\nStatus: Online"+
				"\nStarted: " + channel.StartedAt +
				"\nStream: "+constants.UrlTwitchStream+channel.LoginName
		} else {
			content += "\nStatus: Offline" +
				"\nThumbnail: " + channel.Thumbnail
		}

		util.DiscordBotResponder(content, s, i)
	}
)

// RegisterStream function for registering command for the bot to serve
func RegisterStream(commands *[]discordgo.ApplicationCommand, commandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	*commands = append(*commands, streamCommand)
	commandHandlers[commandWord] = streamCommandHandler
}

