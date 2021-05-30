package main

func buildLeaderboardStatus(a ActivityDetails, k Kraftee) string {
	startOfWeek := getStartOfWeekInUnixTime()
	leaderboard, _ := getAllKrafteeStatsSince(startOfWeek)

	postString := "**Leaderboard** @ post time\n"
	postString += "```\n"
	postString += leaderboard.printActivityCountUpToKraftee(&k)
	postString += leaderboard.printDurationUpToKraftee(&k)

	if a.Type == "Run" {
		postString += leaderboard.printRunDistanceUpToKraftee(&k)
		postString += leaderboard.printRunDurationUpToKraftee(&k)
	}

	if a.Type == "Ride" {
		postString += leaderboard.printRideDistanceUpToKraftee(&k)
		postString += leaderboard.printRideDurationUpToKraftee(&k)
	}

	postString += "```"
	return postString
}
