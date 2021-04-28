package command

import (
	"fmt"

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
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionApplicationCommandResponseData{
				Content: "pong",
			},
		})
		s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
			Content: "hey",
		})
		fmt.Println(i.Data.Options[0].Value)
	}
)

// function for registering command for the bot to serve
func RegisterFollowList(commands *[]discordgo.ApplicationCommand, commandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	*commands = append(*commands, followListCommand)
	commandHandlers["follow-list"] = followListCommandHandler
}
