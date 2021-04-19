package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"twitch-discord-bot/command"

	"github.com/bwmarrin/discordgo"
)

// Bot parameters
var (
	// guild (server) to register commands in
	GuildID = flag.String("g", "833404821809266708", "Guild ID to register commands in (must have permissions there)")

	// bot token
	BotToken = flag.String("t", "", "Bot access token")

	// array with command definitions
	commandDefinitions = []discordgo.ApplicationCommand{}

	// map of command handlers
	commandHandlers = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))
)

// bot session
var s *discordgo.Session

// read command line flags
func init() {
	flag.Parse()
}

// create session
func init() {
	var err error
	s, err = discordgo.New("Bot " + *BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	registerCommands()
}

func registerCommands() {
	// register commands
	command.RegisterPing(&commandDefinitions, commandHandlers)

	// add a handler for handling commands
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// if command is in commandHandlers map, call handler function
		if handlerFunc, ok := commandHandlers[i.Data.Name]; ok {
			handlerFunc(s, i)
		}
	})
}

func main() {
	// just log that bot is running
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is up!")
	})

	// open session to discord
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	// register slash-commands
	for _, v := range commandDefinitions {
		// try to register command
		_, err := s.ApplicationCommandCreate(s.State.User.ID, *GuildID, &v)
		// if not log error
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
	}

	// close session when bot is stopeed
	defer s.Close()

	// graceful shutdown when Interrupting
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("bot stopped")
}
