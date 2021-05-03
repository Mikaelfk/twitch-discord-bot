package main

import (
	"log"
	"os"
	"os/signal"

	"twitch-discord-bot/command"
	"twitch-discord-bot/db"
	"twitch-discord-bot/util"

	"github.com/bwmarrin/discordgo"
)

// Bot variable
var (
	// array with command definitions
	commandDefinitions = []discordgo.ApplicationCommand{}

	// map of command handlers
	commandHandlers = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))
)

// bot session
var session *discordgo.Session

// try to load config
func init() {
	err := util.LoadConfig()
	db.InitDB()
	if err != nil {
		log.Fatalf("Unable to load config: %v", err)
	}
}

// create session
func init() {
	var err error
	session, err = discordgo.New("Bot " + util.Config.DiscordBotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	registerCommands()
}

func registerCommands() {
	// register commands
	command.RegisterPing(&commandDefinitions, commandHandlers)
	command.RegisterPang(&commandDefinitions, commandHandlers)
	command.RegisterChannel(&commandDefinitions, commandHandlers)
	command.RegisterStream(&commandDefinitions, commandHandlers)
<<<<<<< HEAD
<<<<<<< HEAD
	command.RegisterGames(&commandDefinitions, commandHandlers)
=======
	command.RegisterTopCategories(&commandDefinitions, commandHandlers)
>>>>>>> Added the RegisterTopCategories command in the registerCommand function in main.go
=======
	command.RegisterTopCategories(&commandDefinitions, commandHandlers)
>>>>>>> Added the RegisterTopCategories command in the registerCommand function in main.go

	// add a handler for handling commands
	session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// if command is in commandHandlers map, call handler function
		if handlerFunc, ok := commandHandlers[i.Data.Name]; ok {
			handlerFunc(s, i)
		}
	})
}

func main() {
	// just log that bot is running
	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is up (:")
	})

	// open session to discord
	err := session.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	// register slash-commands
	for _, v := range commandDefinitions {
		// try to register command
		_, err := session.ApplicationCommandCreate(session.State.User.ID, util.Config.DiscordServerID, &v)
		// if not log error
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
	}

	// close session when bot is stopeed
	defer session.Close()

	// graceful shutdown when Interrupting
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("bot stopped")
}
