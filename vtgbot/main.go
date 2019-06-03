package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/DaMonkE/VTGbot-monke/lib/characters"
	"github.com/DaMonkE/VTGbot-monke/lib/help"
	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}

	if m.Content == "/help" {
		s.ChannelMessageSend(m.ChannelID, help.Message())
	}

	if strings.HasPrefix(m.Content, "/create") {
		nameStr := strings.TrimSpace(strings.TrimPrefix(m.Content, "/create"))
		if nameStr == "" {
			nameStr = m.Author.Username
		}
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%v : Creating sheet for %v...", m.Author.Mention(), nameStr))
		sheetMsg := characters.MakeBlankSheet(nameStr)
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%v : %v", m.Author.Mention(), sheetMsg))
	}

	if strings.HasPrefix(m.Content, "/set") {
		setStr := strings.Fields(strings.TrimPrefix(m.Content, "/set"))
		name := m.Author.Username
		var attr string
		var num int
		if len(setStr) < 2 || len(setStr) > 3 {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%v : Error, see /help for usage", m.Author.Mention()))
			return
		}
		if len(setStr) == 2 {
			attr = setStr[0]
			num, _ = strconv.Atoi(setStr[1])
		}
		if len(setStr) == 3 {
			name = setStr[0]
			attr = setStr[1]
			num, _ = strconv.Atoi(setStr[2])
		}
		err := characters.SetDots(name, strings.Title(attr), int64(num))
		if err != nil {
			fmt.Println(err)
			s.ChannelMessageSend(m.ChannelID, "Error Occured, check console")
			return
		}
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%v : Set %v for %v to %v", m.Author.Mention(), attr, name, num))
	}

	if strings.HasPrefix(m.Content, "/get") {
		getStr := strings.Fields(strings.TrimPrefix(m.Content, "/get"))
		name := m.Author.Username
		var attr string
		var num int
		if len(getStr) == 1 {
			attr = getStr[0]
		} else if len(getStr) == 2 {
			name = getStr[0]
			attr = getStr[1]
		} else {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%v : Error, see /help for usage", m.Author.Mention()))
			return
		}
		num, err := characters.GetDots(name, strings.Title(attr))
		if err != nil {
			fmt.Println(err)
			s.ChannelMessageSend(m.ChannelID, "Error Occured, check console")
			return
		}
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%v : %v has %v dots in %v", m.Author.Mention(), name, num, attr))
	}
}
