package command

import (
	"fmt"
	"log"
	"strings"
	"twitch-discord-bot/twitchapi"
	"twitch-discord-bot/util"

	"github.com/bwmarrin/discordgo"
)

var (
	// define name and description for command
	followListCommand = discordgo.ApplicationCommand{
		Name:        "follow-list",
		Description: "gives a list of all streamers a user follows",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "username",
				Description: "A twitch username",
				Required:    true,
			},
		},
	}

	// define commandHandler for this command
	followListCommandHandler = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Get all the streamers in a slice
		username := fmt.Sprintf("%v", i.Data.Options[0].Value)
		messageString := "The user " + username + " follows:"
		util.DiscordBotResponder(messageString, s, i)

		streamers, err := twitchapi.GetFollowList(username)
		if err != nil {
			// Prints th error to the user
			_, err = s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
				Content: err.Error(),
			})
			if err != nil {
				log.Println("unable to send follow-up error")
				return
			}

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
			// Prints all the streamers that did not fit in the first message
			for _, v := range streamersStringArray {
				_, err = s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
					Content: strings.TrimSuffix(v, ", "),
				})
				if err != nil {
					log.Println("unable to send follow-up message")
					return
				}
			}
		}
	}
)

// RegisterFollowList function for registering command for the bot to serve
func RegisterFollowList(commands *[]discordgo.ApplicationCommand, commandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	*commands = append(*commands, followListCommand)
	commandHandlers["follow-list"] = followListCommandHandler
}
