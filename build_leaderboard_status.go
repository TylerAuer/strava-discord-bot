package main

func buildLeaderboardStatus(a ActivityDetails, k Kraftee) string {
	startOfWeek := getStartOfWeekInUnixTime()
	kStats, _ := getAllKrafteeStatsSince(startOfWeek)

	lbs := "**Leaderboard** @ post time\n"
	lbs += "```\n"
	lbs += kStats.printActivityCountUpToKraftee(k)
	lbs += kStats.printDurationUpToKraftee(k)

	if a.Type == "Run" {
		lbs += kStats.printRunDistanceUpToKraftee(k)
		lbs += kStats.printRunDurationUpToKraftee(k)
	}

	if a.Type == "Ride" {
		lbs += kStats.printRideDistanceUpToKraftee(k)
		lbs += kStats.printRideDurationUpToKraftee(k)
	}

	lbs += "```"
	return lbs
}
