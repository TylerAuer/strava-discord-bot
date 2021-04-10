package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func postToDiscord(post string) {
	fmt.Println("Making post to Discord")

	token := os.Getenv("DISCORD_BOT_TOKEN")
	c := os.Getenv("DISCORD_CHANNEL_ID")

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	defer dg.Close()

	dg.ChannelMessageSend(c, post)

}
