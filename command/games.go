package command

import (
	"github.com/bwmarrin/discordgo"
	"strconv"
	"strings"
	"twitch-discord-bot/constants"
	"twitch-discord-bot/util"
)


type game struct {
	ArtUrl string `json:"box_art_url"`
	Id     string `json:"id"`
	Name   string `json:"name"`
}

type gamesData struct {
	Data []game `json:"data"`
}

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
			Name:        "game-name", //TODO: Doesn't currently allow spaces. Use "-" instead of space or no spaces at all.
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
		if len(i.Data.Options)==2{
			num = int(i.Data.Options[1].IntValue())
			if num > 10 {
				num = 10
			}
		}
		var content = ""


		games, err := findGames(i.Data.Options[0].StringValue(), num)
		if err != nil {
			content = constants.BotUnexpectedErrorMessage
		} else if len(games.Data) <= 0 {
			content = constants.BotNoResults
		} else {

			if len(games.Data)<num{
				// if the length of the twitch api response is shorter than the requested amount, then...
				num = len(games.Data)
			}
			for index:=0; index < num; index++{
				parts := strings.Split(games.Data[index].ArtUrl, "52x72.jpg")
				icon := parts[0]+constants.DiscordBotImgResolution+".jpg"
				content +=  "\n--------------"+
				    "\nName: "+games.Data[index].Name+
					"\nId:   "+games.Data[index].Id+
					"\nIcon: "+icon
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


//findGames finds the 'first' games that is somewhat similar to gameName
func findGames(gameName string, first int) (gamesData, error) {

	//Replaces possible spaces with "-" (dashes) before calling the handleRequest method
	gameName = strings.Join(strings.Split(gameName , " "),"-")
	URL := constants.UrlTwitchGames + gameName + "&first=" + strconv.Itoa(first)

	var data gamesData
	err := util.HandleRequest(URL, "GET", &data)

	return data, err
}

