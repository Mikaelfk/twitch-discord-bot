package command

import (
	"github.com/bwmarrin/discordgo"
	"twitch-discord-bot/util"
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
		util.DiscordBotResponder("BOOM", s, i)
	}
)

// function for registering command for the bot to serve
func RegisterPang(commands *[]discordgo.ApplicationCommand, commandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	*commands = append(*commands, pangCommand)
	commandHandlers["pang"] = pangCommandHandler
}
