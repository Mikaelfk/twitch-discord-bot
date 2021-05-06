package util

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// DiscordBotResponder sends the content string to discord chat as a response so that you do not need to write these lines of code in every command
func DiscordBotResponder(content string, s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Content: content,
		},
	})

	if err != nil {
		log.Println("unable to send message")
	}
}
