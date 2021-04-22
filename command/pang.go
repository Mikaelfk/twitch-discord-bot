package command

import (
	"github.com/bwmarrin/discordgo"
)

//
// Example command
//

var (
	// define name and description for command
	pangCommand = discordgo.ApplicationCommand{
		Name:        "pang",
		Description: "will pang you",
	}

	// define commandHandler for this command
	pangCommandHandler = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionApplicationCommandResponseData{
				Content: "boom",
			},
		})
	}
)

// function for registering command for the bot to serve
func RegisterPang(commands *[]discordgo.ApplicationCommand, commandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	*commands = append(*commands, pangCommand)
	commandHandlers["pang"] = pangCommandHandler
}
