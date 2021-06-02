package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func getActiveDiscordSession() *discordgo.Session {
	fmt.Println("Connecting to Discord")
	token := os.Getenv("DISCORD_BOT_TOKEN")

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Error creating Discord session,", err)
	}

	return dg
}

// Accepts a Discord Go session and makes a post to the channel matching an ID in an env var
func postToDiscord(dg *discordgo.Session, post string) {
	fmt.Println("Making post to Discord")
	defer dg.Close()
	c := os.Getenv("DISCORD_CHANNEL_ID")
	dg.ChannelMessageSend(c, post)
}

func updateDiscordPost(dg *discordgo.Session, messageId string, newPostContent string) {
	c := os.Getenv("DISCORD_CHANNEL_ID")
	dg.ChannelMessageEdit(c, messageId, newPostContent)
}
