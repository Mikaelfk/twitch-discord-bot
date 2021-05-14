package command

import (
	"log"
	"twitch-discord-bot/constants"
	"twitch-discord-bot/twitchapi"
	"twitch-discord-bot/util"

	"github.com/bwmarrin/discordgo"
)

var (
	// define name and description for command
	teamCommand = discordgo.ApplicationCommand{
		Name:        "team",
		Description: "Will retrieve info about a twitch team",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "team-name",
				Description: "The name of the team",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        "is-live",
				Description: "Boolean for live",
				Required:    true,
			},
		},
	}

	// define commandHandler for this command
	teamCommandHandler = func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		teamName := i.Data.Options[0].StringValue()

		exist, err := twitchapi.TeamExist(teamName)
		if err != nil {
			log.Println("unexpected error")
			util.DiscordBotResponder(constants.BotUnexpectedErrorMsg, s, i)
			return
		}

		if !exist {
			util.DiscordBotResponder(constants.BotNoResultsMsg, s, i)
			return
		}

		responseString := "Team name: " + teamName + "\n"

		var members []string

		isLive := i.Data.Options[1].BoolValue()

		if isLive {
			responseString += "Live team members: \n"
			members, _ = twitchapi.GetLiveTeamMembers(teamName)
		} else {
			responseString += "Team members: \n"
			members, _ = twitchapi.GetAllTeamMembers(teamName)
		}

		for _, member := range members {
			responseString += member + ", "
		}
		util.DiscordBotResponder(responseString, s, i)
	}
)

// RegisterTeam function for registering command for the bot to serve
func RegisterTeam(commands *[]discordgo.ApplicationCommand, commandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	*commands = append(*commands, teamCommand)
	commandHandlers["team"] = teamCommandHandler
}
