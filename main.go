package main

import (
	"log"
	"os"
	"os/signal"

	"twitch-discord-bot/command"
	"twitch-discord-bot/util"

	"github.com/bwmarrin/discordgo"
)

// Bot variable
var (
	// store config ciles
	Config util.Configuration

	// array with command definitions
	commandDefinitions = []discordgo.ApplicationCommand{}

	// map of command handlers
	commandHandlers = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))
)

// bot session
var session *discordgo.Session

// try to load config
func init() {
	err := util.LoadConfig(&Config)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

// create session
func init() {
	var err error
	session, err = discordgo.New("Bot " + Config.DiscordToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	registerCommands()
}

func registerCommands() {
	// register commands
	command.RegisterPing(&commandDefinitions, commandHandlers)

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
		_, err := session.ApplicationCommandCreate(session.State.User.ID, Config.ServerID, &v)
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
