package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/vincent-petithory/dataurl"
	"grabber/grab"
	"net/url"
	"strings"
)

func grabRouter(agent discordAgent) {
	imageData := ""
	format := ""

	if agent.message.Attachments != nil && len(agent.message.Attachments) > 0 {
		fmt.Println("PROUT")
		return
	} else {
		u, err := url.Parse(agent.message.Content)
		if err != nil {
			return
		}
		fMat := strings.Trim(u.Path[len(u.Path)-4:], ".")
		data, err := grab.GetFromLink(agent.message.Content)
		if err != nil {
			return
		}
		format = fMat
		imageData = data
	}
	if format == "jpeg" {
		format = "jpg"
	}
	createEmote(imageData, format, agent)
}

func createEmote(imageData, format string, agent discordAgent) {
	data := dataurl.New([]byte(imageData), "image/"+format)
	id := utils.UUID()
	id = id[:7]
	profile, err := loadProfile(agent, "")
	if err != nil {
		return
	}
	_, err = agent.session.GuildEmojiCreate(profile.Guild, id, data.String(), nil)
	if err != nil {
		_, _ = agent.session.ChannelMessageSend(agent.message.ChannelID, "Cannot create emote")
		return
	}
	_, _ = agent.session.ChannelMessageSend(agent.message.ChannelID, "Success")
}
