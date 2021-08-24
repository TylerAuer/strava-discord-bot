package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

type Discord struct {
	*discordgo.Session
}

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

func getDiscord() Discord {
	fmt.Println("Connecting to Discord")
	token := os.Getenv("DISCORD_BOT_TOKEN")

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Error creating Discord session,", err)
	}

	return Discord{dg}
}

// Accepts a Discord Go session and makes a post to the channel matching an ID in an env var
func postToDiscord(dg *discordgo.Session, post string) {
	fmt.Println("Making post to Discord")
	defer dg.Close()
	c := os.Getenv("DISCORD_CHANNEL_ID")
	dg.ChannelMessageSend(c, post)
}

func (d Discord) getChannelId() string {
	return os.Getenv("DISCORD_CHANNEL_ID")
}

func (d Discord) post(post string) {
	c := d.getChannelId()
	d.ChannelMessageSend(c, post)
}

func (d Discord) updatePost(postToUpdate *discordgo.Message, replacementContent string) {
	d.ChannelMessageEdit(postToUpdate.ChannelID, postToUpdate.ID, replacementContent)
}

func (d Discord) deletePost(postToDelete *discordgo.Message) {
	c := os.Getenv("DISCORD_CHANNEL_ID")

	err := d.ChannelMessageDelete(c, postToDelete.ID)
	if err != nil {
		log.Fatal("Error deleting message,", err)
	}
}

func (d Discord) lastOneHundredMessages() []*discordgo.Message {
	c := os.Getenv("DISCORD_CHANNEL_ID")

	msgs, err := d.ChannelMessages(c, 100, "", "", "")
	if err != nil {
		log.Fatal("Error getting last 100 messages,", err)
	}

	return msgs
}
