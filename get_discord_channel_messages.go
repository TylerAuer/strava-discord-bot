package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func getDiscordChannelMessages(dg *discordgo.Session, channelID string) []*discordgo.Message {
	c := os.Getenv("DISCORD_CHANNEL_ID")

	// Gets the last 100 messages from the channel
	fmt.Println("Getting posts from Discord")

	msgs, err := dg.ChannelMessages(c, 100, "", "", "")
	if err != nil {
		log.Fatal("Error getting last 100 messages,", err)
	}

	return msgs
}
