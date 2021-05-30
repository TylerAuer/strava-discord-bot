package main

func buildLeaderboardPost(lb Leaderboard) string {
	lbPost := "Leaderboard\n"
	lbPost += "```"
	lbPost += lb.printDurationUpToKraftee(nil)
	lbPost += lb.printActivityCountUpToKraftee(nil)
	lbPost += lb.printRunDistanceUpToKraftee(nil)
	lbPost += lb.printRunDurationUpToKraftee(nil)
	lbPost += lb.printRideDistanceUpToKraftee(nil)
	lbPost += lb.printRideDurationUpToKraftee(nil)
	lbPost += "```"
	return lbPost
}
