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


var (
	// define name and description for command
	streamCommand = discordgo.ApplicationCommand{
		Name:        "stream",
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
		URL := constants.URL_TWITCH_CHANNEL_NAME + i.Data.Options[0].StringValue()
		err = util.HandleRequest(URL, http.MethodGet, &channels)

		if err!= nil {
			DiscordBotResponder("Something went wrong...", s, i)
			return
		}

		channel := channels.Data[0].Channel
		content = "Broadcaster: " + channel.DisplayName +
			"\nStream-Title: "+channel.Title +
			"\nLanguage: " + channel.Lang +
			"\nGame: " + channel.GameName
		if channel.IsLive {
			content += "\nStatus: Online"+
				"\nStarted: " + channel.StartedAt
		} else {
			content += "\nStatus: Offline"
		}
		content += "\nThumbnail: " + channel.Thumbnail

		DiscordBotResponder(content, s, i)
	}
)


func DiscordBotResponder(content string, s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Content:content,
		},
	})
}



// RegisterStream function for registering command for the bot to serve
func RegisterStream(commands *[]discordgo.ApplicationCommand, commandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	*commands = append(*commands, streamCommand)
	commandHandlers["stream"] = streamCommandHandler
}

