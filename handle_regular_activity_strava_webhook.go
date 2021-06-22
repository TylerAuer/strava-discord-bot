package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func handleRegularActivityStravaWebhook(k Kraftee, ad ActivityDetails, webhook WebhookData) {
	activityId := fmt.Sprint(webhook.ObjectId)

	postString := buildActivityPost(ad, k)

	// postOrUpdateActivity(idStr, postString, webhook.AspectType, ad, k)

	c := os.Getenv("DISCORD_CHANNEL_ID")

	// Get a connection to Discord, defer closing it
	dg := getActiveDiscordSession()
	defer dg.Close()

	// Collect last 100 messages
	messagesList := getDiscordChannelMessages(dg, c)
	re := "ID: " + activityId

	/**
	For each of the last 100 messages, check if it contains "ID: <activityID>". If one is found
	with a matching URL, update it.

	This is desired even if the Strava webhook type is "create" because Strava's webhook accidentally
	fires duplicate events, often.
	*/
	for i, msg := range messagesList {
		matched, err := regexp.Match(re, []byte(msg.Content))
		if err != nil {
			log.Fatal("Regexp error: ", err)
		}
		if matched {
			fmt.Println("Updating post with id: " + msg.ID + " which is " + fmt.Sprint(i) + " posts from the end of the thread.")
			regexForLeaderboard := regexp.MustCompile(`[*]*Leaderboard[*]* @ post time[\w|\W]*`)
			oldLeaderboard := regexForLeaderboard.Find([]byte(msg.Content))
			updatedPost := postString + "\n" + string(oldLeaderboard)

			updateDiscordPost(dg, msg.ID, updatedPost)
			return
		}
	}

	/**
	The function only reaches here if it did not find a message with a matching URL. That may be
	because the message is brand new or because the activity is old (not within the last 100 messages)

	We only want to create a new post if it is truly new, so that's why we check if strava's webhook
	indicates that this is a "create" type of update.
	*/
	if webhook.AspectType == "create" {
		// Build and post the leaderboard status for a user's post just once. Otherwise, if they update
		// a post a few days later it will botch the whole thing.
		lbs := buildLeaderboardStatus(ad, k)
		postToDiscord(dg, postString+lbs+"\nID: "+activityId)
	}
}
