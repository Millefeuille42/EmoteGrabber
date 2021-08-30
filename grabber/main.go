package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"time"
)

var ownerID string = "268431730967314435" //Please change this when using my bot

// startBot Starts discord bot
func startBot() *discordgo.Session {
	discordBot, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	checkError(err)
	discordBot.AddHandler(messageHandler)
	err = discordBot.Open()
	checkError(err)
	fmt.Println("Discord bot created")
	channel, err := discordBot.UserChannelCreate(ownerID)
	if err != nil {
		return nil
	}
	hostname, _ := os.Hostname()
	_, _ = discordBot.ChannelMessageSend(channel.ID, "Bot up - "+
		time.Now().Format(time.Stamp)+" - "+hostname)

	setUpCloseHandler(discordBot)

	return discordBot
}

func main() {
	err := createDirIfNotExist("./data")
	if err != nil {
		return
	}

	_ = startBot()
	for {
		time.Sleep(time.Second * 3)
	}
}
