// Package command provides commands for the bot
package command

import (
	"fmt"
	"log"
	"twitch-discord-bot/db"
	"twitch-discord-bot/util"

	"github.com/bwmarrin/discordgo"
)

var (
	// define name and description for command
	unSubscribeCommmand = discordgo.ApplicationCommand{
		Name:        "unsubscribe",
		Description: "Will remove the going-live notifications in the current channel for a streamer",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "streamer-name",
				Description: "Name of the streamer to stop getting notifications from",
				Required:    true,
			},
		},
	}

	// define commandHandler for this command
	unSubscribeCommandHandler = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		util.DiscordBotResponder("Trying to remove subscription", s, i)
		channelID := i.ChannelID

		userID, err := util.GetUserID(fmt.Sprintf("%v", i.Data.Options[0].Value))
		if err != nil {
			log.Println("Unable to find a streamer named " + fmt.Sprintf("%v", i.Data.Options[0].Value))
			_, err = s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
				Content: "Unable to remove subscription. Could not find streamer :(",
			})
			if err != nil {
				log.Println("unable to send follow-up message")
			}
			return
		}

		err = db.DeleteSubscription(userID, channelID)
		if err != nil {
			log.Println("Unable to remove subscription for streamer ", fmt.Sprintf("%v", i.Data.Options[0].Value))
			_, err = s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
				Content: "Unable to remove subscription for streamer " + fmt.Sprintf("%v", i.Data.Options[0].Value) + ". Is there an active subscription in this channel?",
			})
			if err != nil {
				log.Println("unable to send follow-up message")
			}
		}

		log.Println("Removed subscription for streamer", fmt.Sprintf("%v", i.Data.Options[0].Value))
		_, err = s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
			Content: "Removed subscription for streamer" + fmt.Sprintf("%v", i.Data.Options[0].Value),
		})
		if err != nil {
			log.Println("unable to send follow-up message")
		}
	}
)

// RegisterSubscribe function for registering command for the bot to serve
func RegisterUnSubscribe(commands *[]discordgo.ApplicationCommand, commandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	*commands = append(*commands, unSubscribeCommmand)
	commandHandlers["unsubscribe"] = unSubscribeCommandHandler
}
