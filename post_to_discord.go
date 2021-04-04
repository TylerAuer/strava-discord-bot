package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func postToDiscord(post string) {

	token := os.Getenv("DISCORD_BOT_TOKEN")
	c := os.Getenv("DISCORD_TEST_CHANNEL_ID")
	// c := os.Getenv("DISCORD_WORKOUT_CHANNEL_ID")

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	defer dg.Close()

	dg.ChannelMessageSend(c, post)

}
