package main

import (
	"fmt"
	"time"
)

func handleWeeklyUpdatePost() {
	// Find 7 days ago in unix epoch time
	secsToLookBack := int64(7 * 24 * 60 * 60)
	epochTime := time.Now().Unix()
	startInEpochTime := epochTime - secsToLookBack

	krafteeCount := len(krafteesByStravaId)

	listOfKrafteeStats, listOfEveryActivity := getAllKrafteeStatsSince(startInEpochTime)

	groupStats := buildStatsFromActivityList("All", "", listOfEveryActivity)
	groupStatsPost := buildGroupStatsPost(groupStats)

	leaderboardPost := buildLeaderboardPost(listOfKrafteeStats)

	post := "**Weekly Update**\n"
	post += "*Here's a summary for " + fmt.Sprint(krafteeCount) + " kraftees over the last week*"
	post += "\n\n" + groupStatsPost
	post += "\n\n" + leaderboardPost

	dg := getActiveDiscordSession()

	postToDiscord(dg, post)
}
