// Package command provides commands for the bot
package command

import (
	"fmt"
	"log"
	"twitch-discord-bot/db"
	"twitch-discord-bot/twitchapi"
	"twitch-discord-bot/util"

	"github.com/bwmarrin/discordgo"
)

//
// Example command
//

var (
	// define name and description for command
	subscribeCommmand = discordgo.ApplicationCommand{
		Name:        "subscribe",
		Description: "Will set up going-live notifications in the current channel for a streamer",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "streamer-name",
				Description: "Name of the streamer to get notifications from",
				Required:    true,
			},
		},
	}

	// define commandHandler for this command
	subscribeCommandHandler = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		util.DiscordBotResponder("Trying to create subscription...", s, i)
		channelID := i.ChannelID

		userID, err := util.GetUserID(fmt.Sprintf("%v", i.Data.Options[0].Value))
		if err != nil {
			log.Println("Unable to find a streamer named " + fmt.Sprintf("%v", i.Data.Options[0].Value))
			_, err = s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
				Content: "Unable to create subscription. Could not find streamer :(",
			})
			if err != nil {
				log.Println("unable to send follow-up message")
			}
			return
		}

		callBackFunction := func(success bool) {
			log.Println("At least callback works please")

			if success {
				log.Println("Added subscription for streamer " + fmt.Sprintf("%v", i.Data.Options[0].Value))
				_, err = s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
					Content: "Created subscription :)",
				})
				if err != nil {
					log.Println("unable to send follow-up message")
				}

				err = db.AddSubscription(userID, channelID)
				if err != nil {
					log.Println("Unable to save subscription in firebase :O")
					log.Println("If you see this, beware of descrepency in webhooks and notifications")
				}
				return
			} else {
				log.Println("Internal error, unable to create subcsription for streamer " + fmt.Sprintf("%v", i.Data.Options[0].Value))
				_, err = s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
					Content: "Internal error, unable to create subcsription for streamer",
				})
				if err != nil {
					log.Println("unable to send follow-up message")
				}
			}
		}

		twitchapi.CreateSubscription(userID, "stream.online", callBackFunction)
	}
)

// RegisterSubscribe function for registering command for the bot to serve
func RegisterSubscribe(commands *[]discordgo.ApplicationCommand, commandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	*commands = append(*commands, subscribeCommmand)
	commandHandlers["subscribe"] = subscribeCommandHandler
}
