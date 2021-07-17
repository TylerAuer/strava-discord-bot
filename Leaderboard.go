package main

import (
	"fmt"
	"sort"
)

type Leaderboard []Stats

/*
Methods that generate strings summarizing the leaderboard through the passed Kraftee (that kraftee)
loses ties.
*/
func (l Leaderboard) composeActivityCountUpToKraftee(k *Kraftee) string {
	l.sortByActivityCount(k) // Sort
	rank := l.findRankOfKrafteeOrLastIfAbsent(k)
	str := "## Activities ##\n" // Header
	currentRank := 0            // Matches the index of the list until multiple Kraftees are tied
	currentStat := 0            // Holds person in front's stat to check for ties
	for i, kraftee := range l {
		if kraftee.AllCount <= 0 {
			break // Stop adding to the leaderboard when you reach a Kraftee with no stats
		}
		// Track stat of person in front to check for ties and adjust rank accordingly
		if currentStat != kraftee.AllCount {
			currentRank = i
			currentStat = kraftee.AllCount
		}
		str += medal[currentRank] + " "                           // Rank
		str += padRight(kraftee.Name, NAME_LENGTH)                // Name
		str += padLeft(fmt.Sprint(kraftee.AllCount), STAT_LENGTH) // Stat
		str += "\n"                                               // Line break
		if i == rank {
			break // Stop when reaching the given kraftee
		}
	}
	str += "\n"
	return str
}

func (l Leaderboard) composeDurationUpToKraftee(k *Kraftee) string {
	l.sortByActivityDuration(k) // Sort
	rank := l.findRankOfKrafteeOrLastIfAbsent(k)
	str := "## Time ##\n" // Header
	currentRank := 0      // Matches the index of the list until multiple Kraftees are tied
	currentStat := 0      // Holds person in front's stat to check for ties
	for i, kraftee := range l {
		if kraftee.AllMovingSeconds <= 0 {
			break // Stop adding to the leaderboard when you reach a Kraftee with no stats
		}
		// Track stat of person in front to check for ties and adjust rank accordingly
		if currentStat != kraftee.AllMovingSeconds {
			currentRank = i
			currentStat = kraftee.AllMovingSeconds
		}
		str += medal[currentRank] + " "                                 // Rank
		str += padRight(kraftee.Name, NAME_LENGTH)                      // Name
		str += padLeft(secToHMS(kraftee.AllMovingSeconds), STAT_LENGTH) // Stat
		str += "\n"                                                     // Line break
		if i == rank {
			break // Stop when reaching the given kraftee
		}
	}
	str += "\n"
	return str
}

func (l Leaderboard) composeRunDistanceUpToKraftee(k *Kraftee) string {
	l.sortByRunDistance(k) // Sort
	rank := l.findRankOfKrafteeOrLastIfAbsent(k)
	str := emojis["run"] + " Distance\n" // Header
	currentRank := 0                     // Matches the index of the list until multiple Kraftees are tied
	currentStat := 0.0                   // Holds person in front's stat to check for ties
	for i, kraftee := range l {
		if kraftee.RunMeters <= 0 {
			break // Stop adding to the leaderboard when you reach a Kraftee with no stats
		}
		// Track stat of person in front to check for ties and adjust rank accordingly
		if currentStat != kraftee.RunMeters {
			currentRank = i
			currentStat = kraftee.RunMeters
		}
		str += medal[currentRank] + " "                                                          // Rank
		str += padRight(kraftee.Name, NAME_LENGTH)                                               // Name
		str += padLeft(fmt.Sprintf("%.1f", metersToMiles(kraftee.RunMeters))+" mi", STAT_LENGTH) // Stat
		str += "\n"                                                                              // Line break
		if i == rank {
			break // Stop when reaching the given kraftee
		}
	}
	str += "\n"
	return str
}

func (l Leaderboard) composeRunDurationUpToKraftee(k *Kraftee) string {
	l.sortByRunTime(k) // Sort
	rank := l.findRankOfKrafteeOrLastIfAbsent(k)
	str := emojis["run"] + " Time\n" // Header
	currentRank := 0                 // Matches the index of the list until multiple Kraftees are tied
	currentStat := 0                 // Holds person in front's stat to check for ties
	for i, kraftee := range l {
		if kraftee.RunMovingSeconds <= 0 {
			break // Stop adding to the leaderboard when you reach a Kraftee with no stats
		}
		// Track stat of person in front to check for ties and adjust rank accordingly
		if currentStat != kraftee.RunMovingSeconds {
			currentRank = i
			currentStat = kraftee.RunMovingSeconds
		}
		str += medal[currentRank] + " "                                 // Rank
		str += padRight(kraftee.Name, NAME_LENGTH)                      // Name
		str += padLeft(secToHMS(kraftee.RunMovingSeconds), STAT_LENGTH) // Stat
		str += "\n"                                                     // Line break
		if i == rank {
			break // Stop when reaching the given kraftee
		}
	}
	str += "\n"
	return str
}

func (l Leaderboard) composeRideDistanceUpToKraftee(k *Kraftee) string {
	l.sortByRideDistance(k) // Sort
	rank := l.findRankOfKrafteeOrLastIfAbsent(k)
	str := emojis["ride"] + " Distance\n" // Header
	currentRank := 0                      // Matches the index of the list until multiple Kraftees are tied
	currentStat := 0.0                    // Holds person in front's stat to check for ties
	for i, kraftee := range l {
		if kraftee.RideMeters <= 0 {
			break // Stop adding to the leaderboard when you reach a Kraftee with no stats
		}
		// Track stat of person in front to check for ties and adjust rank accordingly
		if currentStat != kraftee.RideMeters {
			currentRank = i
			currentStat = kraftee.RideMeters
		}
		str += medal[currentRank] + " "                                                           // Rank
		str += padRight(kraftee.Name, NAME_LENGTH)                                                // Name
		str += padLeft(fmt.Sprintf("%.1f", metersToMiles(kraftee.RideMeters))+" mi", STAT_LENGTH) // Stat
		str += "\n"                                                                               // Line break
		if i == rank {
			break // Stop when reaching the given kraftee
		}
	}
	str += "\n"
	return str
}

func (l Leaderboard) composeRideDurationUpToKraftee(k *Kraftee) string {
	l.sortByRideTime(k) // Sort
	rank := l.findRankOfKrafteeOrLastIfAbsent(k)
	str := emojis["ride"] + " Time\n" // Header
	currentRank := 0                  // Matches the index of the list until multiple Kraftees are tied
	currentStat := 0                  // Holds person in front's stat to check for ties
	for i, kraftee := range l {
		if kraftee.RideMovingSeconds <= 0 {
			break // Stop adding to the leaderboard when you reach a Kraftee with no stats
		}
		// Track stat of person in front to check for ties and adjust rank accordingly
		if currentStat != kraftee.RideMovingSeconds {
			currentRank = i
			currentStat = kraftee.RideMovingSeconds
		}
		str += medal[currentRank] + " "                                  // Rank
		str += padRight(kraftee.Name, NAME_LENGTH)                       // Name
		str += padLeft(secToHMS(kraftee.RideMovingSeconds), STAT_LENGTH) // Stat
		str += "\n"                                                      // Line break
		if i == rank {
			break // Stop when reaching the given kraftee
		}
	}
	str += "\n"
	return str
}

func (l Leaderboard) composeWalkOrHikeDistanceUpToKraftee(k *Kraftee) string {
	l.sortByWalkorHikeDistance(k) // Sort
	rank := l.findRankOfKrafteeOrLastIfAbsent(k)
	str := emojis["walk"] + " Distance\n" // Header
	currentRank := 0                      // Matches the index of the list until multiple Kraftees are tied
	currentStat := 0.0                    // Holds person in front's stat to check for ties
	for i, kraftee := range l {
		if kraftee.WalkOrHikeMeters <= 0 {
			break // Stop adding to the leaderboard when you reach a Kraftee with no stats
		}
		// Track stat of person in front to check for ties and adjust rank accordingly
		if currentStat != kraftee.WalkOrHikeMeters {
			currentRank = i
			currentStat = kraftee.WalkOrHikeMeters
		}
		str += medal[currentRank] + " "                                                                 // Rank
		str += padRight(kraftee.Name, NAME_LENGTH)                                                      // Name
		str += padLeft(fmt.Sprintf("%.1f", metersToMiles(kraftee.WalkOrHikeMeters))+" mi", STAT_LENGTH) // Stat
		str += "\n"                                                                                     // Line break
		if i == rank {
			break // Stop when reaching the given kraftee
		}
	}
	str += "\n"
	return str
}

func (l Leaderboard) composeWalkOrHikeDurationUpToKraftee(k *Kraftee) string {
	l.sortByWalkOrHikeTime(k) // Sort
	rank := l.findRankOfKrafteeOrLastIfAbsent(k)
	str := emojis["walk"] + " Time\n" // Header
	currentRank := 0                  // Matches the index of the list until multiple Kraftees are tied
	currentStat := 0                  // Holds person in front's stat to check for ties
	for i, kraftee := range l {
		if kraftee.WalkOrHikeMovingSeconds <= 0 {
			break // Stop adding to the leaderboard when you reach a Kraftee with no stats
		}
		// Track stat of person in front to check for ties and adjust rank accordingly
		if currentStat != kraftee.WalkOrHikeMovingSeconds {
			currentRank = i
			currentStat = kraftee.WalkOrHikeMovingSeconds
		}
		str += medal[currentRank] + " "                                        // Rank
		str += padRight(kraftee.Name, NAME_LENGTH)                             // Name
		str += padLeft(secToHMS(kraftee.WalkOrHikeMovingSeconds), STAT_LENGTH) // Stat
		str += "\n"                                                            // Line break
		if i == rank {
			break // Stop when reaching the given kraftee
		}
	}
	str += "\n"
	return str
}

func (l Leaderboard) composeCombinedActivityLeaderboard(k *Kraftee) string {
	l.sortByActivityDuration(k)
	var table Table
	for i, kraftee := range l {
		if kraftee.AllMovingSeconds <= 0 {
			break
		}
		name := medal[i] + " " + kraftee.Name
		time := secToHMS(kraftee.AllMovingSeconds)
		count := fmt.Sprint(kraftee.AllCount) + "x"

		table = append(table, TableRow{name, time, count})
	}
	return "### All Activities ###\n" + table.composeAlignedTable(3) + "\n"
}

func (l Leaderboard) composeCombinedRunLeaderboard(k *Kraftee) string {
	l.sortByRunTime(k)
	var table Table
	for i, kraftee := range l {
		if kraftee.RunMovingSeconds <= 0 {
			break
		}
		name := medal[i] + " " + kraftee.Name
		distance := fmt.Sprintf("%.1f", metersToMiles(kraftee.RunMeters)) + " mi"
		time := secToHMS(kraftee.RunMovingSeconds)
		elev := "+" + fmt.Sprintf("%.0f", metersToFeet(kraftee.RunElevationGain)) + "'"

		table = append(table, TableRow{name, distance, time, elev})
	}
	title := emojis["run"] + " Run Leaderboard" + emojis["run"] + "\n"
	return title + table.composeAlignedTable(3) + "\n"
}

func (l Leaderboard) composeCombinedRideLeaderboard(k *Kraftee) string {
	l.sortByRideTime(k)
	var table Table
	for i, kraftee := range l {
		if kraftee.RideMovingSeconds <= 0 {
			break
		}
		name := medal[i] + " " + kraftee.Name
		distance := fmt.Sprintf("%.1f", metersToMiles(kraftee.RideMeters)) + " mi"
		time := secToHMS(kraftee.RideMovingSeconds)
		elev := "+" + fmt.Sprintf("%.0f", metersToFeet(kraftee.RideElevationGain)) + "'"

		table = append(table, TableRow{name, distance, time, elev})
	}
	title := emojis["ride"] + " Ride Leaderboard" + emojis["ride"] + "\n"
	return title + table.composeAlignedTable(3) + "\n"
}

func (l Leaderboard) composeCombinedRunAndWalkLeaderboard(k *Kraftee) string {
	l.sortByWalkOrHikeTime(k)
	var table Table
	for i, kraftee := range l {
		if kraftee.WalkOrHikeMovingSeconds <= 0 {
			break
		}
		name := medal[i] + " " + kraftee.Name
		distance := fmt.Sprintf("%.1f", metersToMiles(kraftee.WalkOrHikeMeters)) + " mi"
		time := secToHMS(kraftee.WalkOrHikeMovingSeconds)
		elev := "+" + fmt.Sprintf("%.0f", metersToFeet(kraftee.WalkOrHikeElevationGain)) + "'"

		table = append(table, TableRow{name, distance, time, elev})
	}
	title := emojis["walk"] + " Walk & Hike Leaderboard" + emojis["hike"] + "\n"
	return title + table.composeAlignedTable(3) + "\n"
}

/*
These methods sort the leaderboard. The passed Kraftee loses all ties.
*/
func (l Leaderboard) sortByActivityCount(k *Kraftee) {
	sort.Slice(l, func(i, j int) bool {
		if k != nil && l[i].AllCount == l[j].AllCount {
			return l[i].ID != k.StravaId
		} else {
			return l[i].AllCount > l[j].AllCount
		}
	})
}

func (l Leaderboard) sortByActivityDuration(k *Kraftee) {
	sort.Slice(l, func(i, j int) bool {
		if k != nil && l[i].AllMovingSeconds == l[j].AllMovingSeconds {
			return l[i].ID != k.StravaId
		} else {
			return l[i].AllMovingSeconds > l[j].AllMovingSeconds
		}
	})
}

func (l Leaderboard) sortByRunDistance(k *Kraftee) {
	sort.Slice(l, func(i, j int) bool {
		if k != nil && l[i].RunMeters == l[j].RunMeters {
			return l[i].ID != k.StravaId
		} else {
			return l[i].RunMeters > l[j].RunMeters
		}
	})
}

func (l Leaderboard) sortByRunTime(k *Kraftee) {
	sort.Slice(l, func(i, j int) bool {
		if k != nil && l[i].RunMovingSeconds == l[j].RunMovingSeconds {
			return l[i].ID != k.StravaId
		} else {
			return l[i].RunMovingSeconds > l[j].RunMovingSeconds
		}
	})
}

func (l Leaderboard) sortByRideDistance(k *Kraftee) {
	sort.Slice(l, func(i, j int) bool {
		if k != nil && l[i].RideMeters == l[j].RideMeters {
			return l[i].ID != k.StravaId
		} else {
			return l[i].RideMeters > l[j].RideMeters
		}
	})
}

func (l Leaderboard) sortByRideTime(k *Kraftee) {
	sort.Slice(l, func(i, j int) bool {
		if k != nil && l[i].RideMovingSeconds == l[j].RideMovingSeconds {
			return l[i].ID != k.StravaId
		} else {
			return l[i].RideMovingSeconds > l[j].RideMovingSeconds
		}
	})
}

func (l Leaderboard) sortByWalkorHikeDistance(k *Kraftee) {
	sort.Slice(l, func(i, j int) bool {
		if k != nil && l[i].WalkOrHikeMeters == l[j].WalkOrHikeMeters {
			return l[i].ID != k.StravaId
		} else {
			return l[i].WalkOrHikeMeters > l[j].WalkOrHikeMeters
		}
	})
}

func (l Leaderboard) sortByWalkOrHikeTime(k *Kraftee) {
	sort.Slice(l, func(i, j int) bool {
		if k != nil && l[i].WalkOrHikeMovingSeconds == l[j].WalkOrHikeMovingSeconds {
			return l[i].ID != k.StravaId
		} else {
			return l[i].WalkOrHikeMovingSeconds > l[j].WalkOrHikeMovingSeconds
		}
	})
}

// Returns the rank of the passed kraftee or the last index of the list if the kraftee
// is not found.
func (l Leaderboard) findRankOfKrafteeOrLastIfAbsent(k *Kraftee) int {
	if k == nil {
		return len(l) - 1 // Return last element if no kraftee is passed
	} else {
		// Return rank of kraftee k
		for i, kInList := range l {
			if kInList.ID == k.StravaId {
				return i
			}
		}
		// Kraftee was passed but not found so return last element
		return len(l) - 1
	}
}

func (l Leaderboard) composeLeaderboardPost() string {
	lbPost := "Leaderboard\n"
	lbPost += "```"
	lbPost += l.composeDurationUpToKraftee(nil)
	lbPost += l.composeActivityCountUpToKraftee(nil)
	lbPost += l.composeRunDistanceUpToKraftee(nil)
	lbPost += l.composeRunDurationUpToKraftee(nil)
	lbPost += l.composeRideDistanceUpToKraftee(nil)
	lbPost += l.composeRideDurationUpToKraftee(nil)
	lbPost += l.composeWalkOrHikeDistanceUpToKraftee(nil)
	lbPost += l.composeWalkOrHikeDurationUpToKraftee(nil)
	lbPost += "```"
	return lbPost
}
