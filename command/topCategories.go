package command

import (
	"strconv"
	"twitch-discord-bot/constants"
	"twitch-discord-bot/twitchAPI"
	"twitch-discord-bot/util"

	"github.com/bwmarrin/discordgo"
)

var (
	// define name and description for command
	topCategoriesCommand = discordgo.ApplicationCommand{
		Name:        "top-categories",
		Description: "will fetch all the current top categories",
	}

	// define commandHandler for this command
	topCategoriesCommandHandler = func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		// Stores the current top 13 categories on twitch
		contents, err := twitchAPI.GetTopCategories()

		// Displays an error if something goes wrong
		if err != nil {
			util.DiscordBotResponder(constants.BotUnexpectedErrorMsg, s, i)
		}

		// The response
		response := "The current top 13 games are: "

		// For each content in the contents list, add it to the response string
		for index, content := range contents {
			response += "\n" + strconv.Itoa(index+1) + ". " + content
		}
		util.DiscordBotResponder(response, s, i)
	}
)

// function for registering command for the bot to serve
func RegisterTopCategories(commands *[]discordgo.ApplicationCommand, commandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	*commands = append(*commands, topCategoriesCommand)
	commandHandlers["top-categories"] = topCategoriesCommandHandler
}
