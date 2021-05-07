package command

import (
	"fmt"
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
				Required:    false,
			},
		},
	}

	// define commandHandler for this command
	teamCommandHandler = func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		var members []string

		teamName := fmt.Sprintf("%v", i.Data.Options[0].Value)
		isLive := i.Data.Options[1].BoolValue()

		print(teamName)
		print(isLive)

		exist, err := twitchapi.TeamExist(teamName)
		if err != nil {
			log.Println("unexpected error")
			return
		}

		if !exist {
			util.DiscordBotResponder(constants.BotNoResultsMsg, s, i)
		}

		if isLive {
			members, _ = twitchapi.GetLiveTeamMembers(teamName)
		} else {
			members, _ = twitchapi.GetAllTeamMembers(teamName)
		}

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
