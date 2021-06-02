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
func (l Leaderboard) printActivityCountUpToKraftee(k *Kraftee) string {
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

func (l Leaderboard) printDurationUpToKraftee(k *Kraftee) string {
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

func (l Leaderboard) printRunDistanceUpToKraftee(k *Kraftee) string {
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

func (l Leaderboard) printRunDurationUpToKraftee(k *Kraftee) string {
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

func (l Leaderboard) printRideDistanceUpToKraftee(k *Kraftee) string {
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

func (l Leaderboard) printRideDurationUpToKraftee(k *Kraftee) string {
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

func (l Leaderboard) printWalkOrHikeDistanceUpToKraftee(k *Kraftee) string {
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

func (l Leaderboard) printWalkOrHikeDurationUpToKraftee(k *Kraftee) string {
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
