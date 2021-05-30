package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
)

type Leaderboard []Stats

/*
Methods that generate strings summarizing the leaderboard through the passed Kraftee (that kraftee)
loses ties.
*/
func (l Leaderboard) printActivityCountUpToKraftee(k Kraftee) string {
	l.sortByActivityCount(k) // Sort
	rank, err := l.findRankOfKraftee(k)
	if err != nil {
		log.Fatal(err)
	}
	str := "## Activities ##\n" // Header
	for i, kraftee := range l {
		str += fmt.Sprint(i+1) + " "                              // Rank
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

func (l Leaderboard) printDurationUpToKraftee(k Kraftee) string {
	l.sortByActivityDuration(k) // Sort
	rank, err := l.findRankOfKraftee(k)
	if err != nil {
		log.Fatal(err)
	}
	str := "## Time ##\n" // Header
	for i, kraftee := range l {
		str += fmt.Sprint(i+1) + " "                                    // Rank
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

func (l Leaderboard) printRunDistanceUpToKraftee(k Kraftee) string {
	l.sortByRunDistance(k) // Sort
	rank, err := l.findRankOfKraftee(k)
	if err != nil {
		log.Fatal(err)
	}
	str := emojis["run"] + " Distance\n" // Header
	for i, kraftee := range l {
		str += fmt.Sprint(i+1) + " "                                                             // Rank
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

func (l Leaderboard) printRunDurationUpToKraftee(k Kraftee) string {
	l.sortByRunTime(k) // Sort
	rank, err := l.findRankOfKraftee(k)
	if err != nil {
		log.Fatal(err)
	}
	str := emojis["run"] + " Time\n" // Header
	for i, kraftee := range l {
		str += fmt.Sprint(i+1) + " "                                    // Rank
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

func (l Leaderboard) printRideDistanceUpToKraftee(k Kraftee) string {
	l.sortByRideDistance(k) // Sort
	rank, err := l.findRankOfKraftee(k)
	if err != nil {
		log.Fatal(err)
	}
	str := emojis["ride"] + " Distance\n" // Header
	for i, kraftee := range l {
		str += fmt.Sprint(i+1) + " "                                                              // Rank
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

func (l Leaderboard) printRideDurationUpToKraftee(k Kraftee) string {
	l.sortByRideTime(k) // Sort
	rank, err := l.findRankOfKraftee(k)
	if err != nil {
		log.Fatal(err)
	}
	str := emojis["ride"] + " Time\n" // Header
	for i, kraftee := range l {
		str += fmt.Sprint(i+1) + " "                                     // Rank
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

/*
These methods sort the leaderboard. The passed Kraftee loses all ties.
*/
func (l Leaderboard) sortByActivityCount(k Kraftee) {
	sort.Slice(l, func(i, j int) bool {
		if l[i].AllCount == l[j].AllCount {
			return l[i].ID != k.StravaId
		} else {
			return l[i].AllCount > l[j].AllCount
		}
	})
}

func (l Leaderboard) sortByActivityDuration(k Kraftee) {
	sort.Slice(l, func(i, j int) bool {
		if l[i].AllMovingSeconds == l[j].AllMovingSeconds {
			return l[i].ID != k.StravaId
		} else {
			return l[i].AllMovingSeconds > l[j].AllMovingSeconds
		}
	})
}

func (l Leaderboard) sortByRunDistance(k Kraftee) {
	sort.Slice(l, func(i, j int) bool {
		if l[i].RunMeters == l[j].RunMeters {
			return l[i].ID != k.StravaId
		} else {
			return l[i].RunMeters > l[j].RunMeters
		}
	})
}

func (l Leaderboard) sortByRunTime(k Kraftee) {
	sort.Slice(l, func(i, j int) bool {
		if l[i].RunMovingSeconds == l[j].RunMovingSeconds {
			return l[i].ID != k.StravaId
		} else {
			return l[i].RunMovingSeconds > l[j].RunMovingSeconds
		}
	})
}

func (l Leaderboard) sortByRideDistance(k Kraftee) {
	sort.Slice(l, func(i, j int) bool {
		if l[i].RideMeters == l[j].RideMeters {
			return l[i].ID != k.StravaId
		} else {
			return l[i].RideMeters > l[j].RideMeters
		}
	})
}

func (l Leaderboard) sortByRideTime(k Kraftee) {
	sort.Slice(l, func(i, j int) bool {
		if l[i].RideMovingSeconds == l[j].RideMovingSeconds {
			return l[i].ID != k.StravaId
		} else {
			return l[i].RideMovingSeconds > l[j].RideMovingSeconds
		}
	})
}

func (l Leaderboard) findRankOfKraftee(k Kraftee) (int, error) {
	for i, kInList := range l {
		if kInList.ID == k.StravaId {
			return i, nil
		}
	}
	return 0, errors.New("did not find kraftee in list of Kraftee stats")
}
