package command

import (
	"github.com/bwmarrin/discordgo"
	"net/http"
	"strconv"
	"twitch-discord-bot/util"
)

//
// Example command - for authentication with twitch API
//


// Example struct
type twitchChannels struct {
	Data []struct {
		DisplayName string `json:"display_name"`
		IsLive      bool   `json:"is_live"`
		Title       string `json:"title"`
	} `json:"data"`
}

var (
	// define name and description for command
	streamCommand = discordgo.ApplicationCommand{
		Name:        "stream",
		Description: "will get info about stream",
	}

	// define commandHandler for this command
	streamCommandHandler = func(s *discordgo.Session, i *discordgo.InteractionCreate) {


		// TODO: 
		var channel twitchChannels
		_ = util.HandleRequest("https://api.twitch.tv/helix/search/channels?query=pokimane", http.MethodGet, &channel)

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionApplicationCommandResponseData{
				Content: "display_name: "+channel.Data[0].DisplayName+"\ntitle: "+channel.Data[0].Title+"\nis_live: "+strconv.FormatBool(channel.Data[0].IsLive),
			},
		})
	}
)

// RegisterStream function for registering command for the bot to serve
func RegisterStream(commands *[]discordgo.ApplicationCommand, commandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	*commands = append(*commands, streamCommand)
	commandHandlers["stream"] = streamCommandHandler
}

