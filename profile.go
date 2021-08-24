package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type userProfile struct {
	Login	string
	Id		string
	Guild	string
}

func loadProfile(agent discordAgent, id string) (userProfile, error) {
	data := userProfile{}

	if id == "" {
		id = agent.message.Author.ID
	}
	fileData, err := ioutil.ReadFile(fmt.Sprintf("./data/%s.json", id))
	if err != nil {
		logErrorToChan(agent, err)
		return userProfile{}, err
	}

	err = json.Unmarshal(fileData, &data)
	if err != nil {
		logErrorToChan(agent, err)
		return userProfile{}, err
	}

	return data, nil
}

func writeProfile(agent discordAgent, profile userProfile) error {
	_, err := createFileIfNotExist("./data/" + agent.message.Author.ID + ".json")
	if err != nil {
		logErrorToChan(agent, err)
		return err
	}
	data, err := json.MarshalIndent(profile, "", "\t")
	if err != nil {
		logErrorToChan(agent, err)
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("./data/%s.json", profile.Id), data, 0677)
	if err != nil {
		logErrorToChan(agent, err)
		return err
	}
	return nil
}

func createProfile(agent discordAgent) {
	args := strings.Split(agent.message.Content, " ")
	if len(args) <= 1 {
		_, _ = agent.session.ChannelMessageSend(agent.message.ChannelID, "Please provide arguments")
		return
	}

	_, err := agent.session.Guild(args[1])
	if err != nil {
		_, _ = agent.session.ChannelMessageSend(agent.message.ChannelID, "This server doesn't exist or I can't see it")
		return
	}

	profile := userProfile {
		Login: agent.message.Author.Username,
		Id:    agent.message.Author.ID,
		Guild: args[1],
	}

	if writeProfile(agent, profile) == nil {
		_, _ = agent.session.ChannelMessageSend(agent.message.ChannelID, "You are now registered")
	}
}