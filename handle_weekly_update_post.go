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

	groupStats := listOfEveryActivity.buildStats("All", "")
	groupStatsPost := groupStats.printGroupStats()

	leaderboardPost := listOfKrafteeStats.composeLeaderboardPost()

	msg := "**Weekly Update**\n"
	msg += "*Here's a summary for " + fmt.Sprint(krafteeCount) + " kraftees over the last week*"
	msg += "\n\n" + groupStatsPost
	msg += "\n\n" + leaderboardPost

	dg := getDiscord()
	dg.Close()

	dg.post(msg)
}
