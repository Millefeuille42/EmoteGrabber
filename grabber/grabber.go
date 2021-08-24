package main

import (
	_ "image/png"
	"net/http"
	"net/url"
	"strings"
)

func isValidFormat(u *url.URL) bool {
	path := u.Path
	resp := strings.HasSuffix(path, ".png")
	return resp
}

func parseUri(uri string, agent discordAgent) string {
	_, err := url.ParseRequestURI(uri)
	if err != nil {
		return ""
	}

	u, err := url.Parse(uri)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return ""
	}
	if u.Host != "cdn.discordapp.com" {
		_, _ = agent.session.ChannelMessageSend(agent.message.ChannelID, "Image comes from invalid host")
		return ""
	}
	if !isValidFormat(u) {
		_, _ = agent.session.ChannelMessageSend(agent.message.ChannelID, "Image is not png")
		return ""
	}
	response, err := http.Get(uri)
	if err != nil {
		_, _ = agent.session.ChannelMessageSend(agent.message.ChannelID, "Unable to get image")
		return ""
	}
	data, err := ReadHTTPResponse(response)
	if err != nil {
		_, _ = agent.session.ChannelMessageSend(agent.message.ChannelID, "Unable to get image")
		return ""
	}
	return string(data)
}
