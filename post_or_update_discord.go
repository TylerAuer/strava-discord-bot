package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func postOrUpdateActivity(activityID string, postContent string, stravaWebhookDeclaredType string, a ActivityDetails, k Kraftee) {
	c := os.Getenv("DISCORD_CHANNEL_ID")

	// Get a connection to Discord, defer closing it
	dg := getActiveDiscordSession()
	defer dg.Close()

	// Collect last 100 messages
	msgs := getDiscordChannelMessages(dg, c)

	re := "ID: " + activityID

	/**
	For each of the last 100 messages, check if it contains "ID: <activity id>". If one is found
	with a matching URL, update it.

	This is desired even if the Strava webhook type is "create" because Strava's webhook accidentally
	fires duplicate events, often.
	*/
	for i, m := range msgs {
		matched, err := regexp.Match(re, []byte(m.Content))
		if err != nil {
			log.Fatal("Regexp error: ", err)
		}
		if matched {
			fmt.Println("Updating post with id: " + m.ID + " which is " + fmt.Sprint(i) + " posts from the end of the thread.")
			dg.ChannelMessageEdit(m.ChannelID, m.ID, postContent)
			return
		}
	}

	/**
	The function only reaches here if it did not find a message with a matching URL. That may be
	because the message is brand new or because the activity is old (not within the last 100 messages)

	We only want to create a new post if it is truly new, so that's why we check if strava's webhook
	indicates that this is a "create" type of update.
	*/
	if stravaWebhookDeclaredType == "create" {
		postToDiscord(dg, postContent)

		// Build and post the leaderboard status for a user's post just once. Otherwise, if they update
		// a post a few days later it will botch the whole thing.
		lbs := buildLeaderboardStatus(a, k)
		postToDiscord(dg, lbs)

	}
}
