package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"bot/config"
	"bot/handlers"
)

func main() {
	dg, err := discordgo.New("Bot " + config.DiscordToken())
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	defer dg.Close()
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages)

	h := handlers.AllHandlers()
	dg.AddHandler(h.MessageHandler)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "hello" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hello, World!")
	}
}
