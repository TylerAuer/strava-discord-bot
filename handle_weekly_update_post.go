package main

import "time"

func handleWeeklyUpdatePost() {
	// Find 7 days ago in unix epoch time
	secsToLookBack := int64(7 * 24 * 60 * 60)
	epochTime := time.Now().Unix()
	startInEpochTime := epochTime - secsToLookBack

	listOfKrafteeStats, listOfEveryActivity := getAllKrafteeStats(startInEpochTime)

	groupStats := compileStatsFromActivities("All", "", listOfEveryActivity)
	groupStatsPost := buildGroupStatsPost(groupStats)

	leaderboardPost := buildLeaderboardPost(listOfKrafteeStats)

	post := "**Weekly Update Post**"
	post += "\n\n" + groupStatsPost
	post += "\n\n" + leaderboardPost
	postToDiscord(post)
}
