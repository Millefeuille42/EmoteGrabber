package main

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

// discordAgent Contains discord's session and message structs, and the guild's command channel
type discordAgent struct {
	session *discordgo.Session
	message *discordgo.MessageCreate
	channel string
}

func messageRouter(agent discordAgent) {
	if strings.HasPrefix(agent.message.Content, "/register") {
		createProfile(agent)
		return
	}
	grabRouter(agent)
}

// messageHandler Discord bot message handler
func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	botID, _ := session.User("@me")
	agent := discordAgent{
		session: session,
		message: message,
	}

	if message.Author.ID == botID.ID {
		return
	}
	messageRouter(agent)
}
