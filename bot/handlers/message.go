package handlers

import "github.com/bwmarrin/discordgo"

type messageHandler struct{}

func newMessageHandler() messageHandler {
	return messageHandler{}
}

func (messageHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "hello" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hello, World!")
	}
}
