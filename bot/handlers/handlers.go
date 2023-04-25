package handlers

import "github.com/bwmarrin/discordgo"

type Handler struct {
	MessageHandler func(*discordgo.Session, *discordgo.MessageCreate)
}

func AllHandlers() Handler {
	return Handler{
		newMessageHandler().Handle,
	}
}
