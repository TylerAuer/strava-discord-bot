package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

// Accepts a Discord Go session and makes a post to the channel matching an ID in an env var
func postToDiscord(dg *discordgo.Session, post string) {
	fmt.Println("Making post to Discord")
	defer dg.Close()
	c := os.Getenv("DISCORD_CHANNEL_ID")
	dg.ChannelMessageSend(c, post)
}
