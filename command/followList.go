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
				Name:        "Username",
				Description: "String option",
				Required:    true,
			},
		},
	}

	// define commandHandler for this command
	followListCommandHandler = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		username := fmt.Sprintf("%v", i.Data.Options[0].Value)
		// Get all the streamers in a slice
		streamers, err := twitchAPI.GetFollowList(username)
		if err != nil {
			// Prints th error to the user
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionApplicationCommandResponseData{
					Content: err.Error(),
				},
			})
		} else {
			streamersString := ""
			index := 0
			streamersStringArray := []string{}
			// Loops through all the streamers, and adds them to strings, there is a max number of 60 streamers per string.
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
			// Prints all the streamers that did not fit in the first message
			for j, v := range streamersStringArray {
				if j > 0 {
					s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
						Content: v,
					})
				}
			}
		}
	}
)

// function for registering command for the bot to serve
func RegisterFollowList(commands *[]discordgo.ApplicationCommand, commandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	*commands = append(*commands, followListCommand)
	commandHandlers["follow-list"] = followListCommandHandler
}
