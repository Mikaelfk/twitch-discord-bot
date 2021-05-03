package command

import (
	"fmt"
	"twitch-discord-bot/twitchAPI"

	"github.com/bwmarrin/discordgo"
)

//
// Example command
//

var (
	// define name and description for command
	followListCommand = discordgo.ApplicationCommand{
		Name:        "follow-list",
		Description: "gives a list of all streamers a user follows",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "string-option",
				Description: "String option",
				Required:    true,
			},
		},
	}

	// define commandHandler for this command
	followListCommandHandler = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		username := fmt.Sprintf("%v", i.Data.Options[0].Value)
		streamers := twitchAPI.GetFollowList(username)
		streamersString := ""
		index := 0
		streamersStringArray := []string{}
		for _, v := range streamers {
			streamersString += v + ", "
			if index > 60 {
				streamersStringArray = append(streamersStringArray, streamersString)
				streamersString = ""
				index = -1
			}
			index++
		}
		streamersStringArray = append(streamersStringArray, streamersString)

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionApplicationCommandResponseData{
				Content: "The user " + username + " follows: \n" + streamersStringArray[0],
			},
		})
		for j, v := range streamersStringArray {
			if j > 0 {
				s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
					Content: v,
				})
			}
		}
	}
)

// function for registering command for the bot to serve
func RegisterFollowList(commands *[]discordgo.ApplicationCommand, commandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	*commands = append(*commands, followListCommand)
	commandHandlers["follow-list"] = followListCommandHandler
}
