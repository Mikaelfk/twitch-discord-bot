package command

import (
	"strings"
	"twitch-discord-bot/constants"
	"twitch-discord-bot/twitchapi"
	"twitch-discord-bot/util"

	"github.com/bwmarrin/discordgo"
)

// Some constants
const gamesCommandWord = "games"
const numGamesDefault = 1 // By default, only show 1 game

var (
	// define name and description for command
	gamesCommand = discordgo.ApplicationCommand{
		Name:        gamesCommandWord,
		Description: "Will show games registered by Twitch",
		Options: []*discordgo.ApplicationCommandOption{{

			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "game-name",
			Description: "Gets game IDs by name.",
			Required:    true,
		},
			{

				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "num-games",
				Description: "Specify how many results you want. (up to 10)",
				Required:    false,
			},
		},
	}

	// define commandHandler for this command
	gamesCommandHandler = func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		num := numGamesDefault
		if len(i.Data.Options) == 2 {
			num = int(i.Data.Options[1].IntValue())
			if num > 10 {
				num = 10
			}
		}
		var content = ""

		games, err := twitchapi.FindGames(i.Data.Options[0].StringValue(), num)
		if err != nil {
			content = constants.BotUnexpectedErrorMsg
		} else if len(games.Data) == 0 {
			content = constants.BotNoResultsMsg
		} else {

			if len(games.Data) < num {
				// if the length of the twitch api response is shorter than the requested amount, then...
				num = len(games.Data)
			}
			for index := 0; index < num; index++ {
				parts := strings.Split(games.Data[index].ArtURL, "52x72.jpg")
				icon := parts[0] + constants.DiscordBotImgResolution + ".jpg"
				content += "\n--------------" +
					"\nName: " + games.Data[index].Name +
					"\nId:   " + games.Data[index].ID +
					"\nIcon: " + icon
			}
		}
		util.DiscordBotResponder(content, s, i)
	}
)

// RegisterGames function for registering command for the bot to serve
func RegisterGames(commands *[]discordgo.ApplicationCommand, commandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	*commands = append(*commands, gamesCommand)
	commandHandlers[gamesCommandWord] = gamesCommandHandler
}
