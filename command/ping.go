package command

import (
	"github.com/bwmarrin/discordgo"
)

//
// Example command
//

var (
	// define name and description for command
	pingCommand = discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "will pong you",
	}

	// define commandHandler for this command
	pingCommandHandler = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionApplicationCommandResponseData{
				Content: "pong",
			},
		})
	}
)

// function for registering command for the bot to serve
func RegisterPing(commands *[]discordgo.ApplicationCommand, commandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	*commands = append(*commands, pingCommand)
	commandHandlers["ping"] = pingCommandHandler
}
