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
