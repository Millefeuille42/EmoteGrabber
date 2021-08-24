package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/vincent-petithory/dataurl"
	"math/rand"
	"strings"
)

// discordAgent Contains discord's session and message structs, and the guild's command channel
type discordAgent struct {
	session *discordgo.Session
	message *discordgo.MessageCreate
	channel string
}

func getFromLink(agent discordAgent) {
	content := parseUri(agent.message.Content, agent)
	if content == "" {
		return
	}
	data := dataurl.New([]byte(content), "image/png")
	fmt.Println(data.String())
	id := rand.Int()
	profile, err := loadProfile(agent, "")
	if err != nil {
		return
	}
	_, err = agent.session.GuildEmojiCreate(profile.Guild, fmt.Sprintf("%d", id), data.String(), nil)
	if err != nil {
		_, _ = agent.session.ChannelMessageSend(agent.message.ChannelID, "Cannot create emote")
		return
	}
	_, _ = agent.session.ChannelMessageSend(agent.message.ChannelID, "Success")
}

func messageRouter(agent discordAgent) {
	if strings.HasPrefix(agent.message.Content, "/register") {
		createProfile(agent)
		return
	}
	if agent.message.Attachments != nil && len(agent.message.Attachments) > 0 {
		fmt.Println("PROUT")
		return
	}
	getFromLink(agent)
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
