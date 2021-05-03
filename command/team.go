package command

import (
	"twitch-discord-bot/twitchAPI"
	"twitch-discord-bot/util"

	"github.com/bwmarrin/discordgo"
)

//
// Example command
//

var (
	// define name and description for command
	teamCommand = discordgo.ApplicationCommand{
		Name:        "team",
		Description: "will retrieve info about a twitch team",
	}

	// define commandHandler for this command
	teamCommandHandler = func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		teamName, _ := twitchAPI.GetTeamName("luminosity")
		members, _ := twitchAPI.GetTeamMembers("luminosity")

		responseString := "Team name: " + teamName + "\n"

		responseString += "Team members: \n"
		for _, member := range members {
			responseString += member + ", "
		}

		util.DiscordBotResponder(responseString, s, i)
	}
)

// function for registering command for the bot to serve
func RegisterTeam(commands *[]discordgo.ApplicationCommand, commandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	*commands = append(*commands, teamCommand)
	commandHandlers["team"] = teamCommandHandler
}
