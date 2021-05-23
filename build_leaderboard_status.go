package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
)

var DEF_NAME_LEN = 10
var DEF_STAT_LEN = 7

func buildLeaderboardStatus(a ActivityDetails, k Kraftee) string {
	startOfWeek := getStartOfWeekInUnixTime()
	kStats, _ := getAllKrafteeStatsSince(startOfWeek)

	lbs := "**Leaderboard** @ post time\n"
	lbs += "```\n"

	lbs += "## Activities ##\n"
	sort.Slice(kStats, func(i, j int) bool { return kStats[i].AllCount > kStats[j].AllCount })
	rank, err := findKrafteeRankinStatsList(kStats, k)
	if err != nil {
		log.Fatal(err)
	}
	for i, kraftee := range kStats {
		if i > rank {
			break // Only show kraftees with data
		}
		lbs += fmt.Sprint(i+1) + " " + kraftee.PadName(DEF_NAME_LEN) + padLeft(fmt.Sprint(kraftee.AllCount), DEF_STAT_LEN) + "\n"
	}
	lbs += "\n"

	lbs += "## Time ##\n"
	sort.Slice(kStats, func(i, j int) bool { return kStats[i].AllMovingSeconds > kStats[j].AllMovingSeconds })
	rank, err = findKrafteeRankinStatsList(kStats, k)
	if err != nil {
		log.Fatal(err)
	}
	for i, kraftee := range kStats {
		if i <= rank {
			lbs += fmt.Sprint(i+1) + " " + kraftee.PadName(DEF_NAME_LEN) + padLeft(secToHMS(kraftee.AllMovingSeconds), DEF_STAT_LEN) + "\n"
		} else {
			break // Only show kraftees with data
		}
	}
	lbs += "\n"

	if a.Type == "Run" {
		lbs += emojis["run"] + " Distance\n"
		sort.Slice(kStats, func(i, j int) bool { return kStats[i].RunMeters > kStats[j].RunMeters })
		rank, err = findKrafteeRankinStatsList(kStats, k)
		if err != nil {
			log.Fatal(err)
		}
		for i, kraftee := range kStats {
			if i > rank {
				break // Only show kraftees with data
			}
			lbs += fmt.Sprint(i+1) + " " + kraftee.PadName(DEF_NAME_LEN) + padLeft(fmt.Sprintf("%.1f", metersToMiles(kraftee.RunMeters))+" mi", DEF_STAT_LEN) + "\n"
		}
		lbs += "\n"

		lbs += emojis["run"] + " Time\n"
		sort.Slice(kStats, func(i, j int) bool { return kStats[i].RunMovingSeconds > kStats[j].RunMovingSeconds })
		rank, err = findKrafteeRankinStatsList(kStats, k)
		if err != nil {
			log.Fatal(err)
		}
		for i, kraftee := range kStats {
			if i > rank {
				break // Only show kraftees with data
			}
			lbs += fmt.Sprint(i+1) + " " + kraftee.PadName(DEF_NAME_LEN) + padLeft(secToHMS(kraftee.RunMovingSeconds), DEF_STAT_LEN) + "\n"
		}
	}

	if a.Type == "Ride" {
		lbs += emojis["ride"] + " Distance\n"
		sort.Slice(kStats, func(i, j int) bool { return kStats[i].RideMeters > kStats[j].RideMeters })
		rank, err = findKrafteeRankinStatsList(kStats, k)
		if err != nil {
			log.Fatal(err)
		}
		for i, kraftee := range kStats {
			if i > rank {
				break // Only show kraftees with data
			}
			lbs += fmt.Sprint(i+1) + " " + kraftee.PadName(DEF_NAME_LEN) + padLeft(fmt.Sprintf("%.1f", metersToMiles(kraftee.RideMeters))+" mi", DEF_STAT_LEN) + "\n"
		}
		lbs += "\n"

		lbs += emojis["ride"] + " Time\n"
		sort.Slice(kStats, func(i, j int) bool { return kStats[i].RideMovingSeconds > kStats[j].RideMovingSeconds })
		rank, err = findKrafteeRankinStatsList(kStats, k)
		if err != nil {
			log.Fatal(err)
		}
		for i, kraftee := range kStats {
			if i > rank {
				break // Only show kraftees with data
			}
			lbs += fmt.Sprint(i+1) + " " + kraftee.PadName(DEF_NAME_LEN) + padLeft(secToHMS(kraftee.RideMovingSeconds), DEF_STAT_LEN) + "\n"
		}
	}

	lbs += "```"
	return lbs
}

func findKrafteeRankinStatsList(kStats []Stats, k Kraftee) (int, error) {
	for i, kInList := range kStats {
		if kInList.ID == k.StravaId {
			return i, nil
		}
	}
	return 0, errors.New("did not find kraftee in list of Kraftee stats")
}

func padLeft(s string, length int) string {
	paddedString := s
	for {
		if len(paddedString) >= length {
			return paddedString
		}
		paddedString = " " + paddedString
	}
}
