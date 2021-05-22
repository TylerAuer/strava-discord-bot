package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
)

func buildLeaderboardStatus(a ActivityDetails, k Kraftee) string {
	startOfWeek := getStartOfWeekInUnixTime()
	kStats, _ := getAllKrafteeStatsSince(startOfWeek)

	lbs := "\n**Leaderboard** @ post time\n"
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
		lbs += fmt.Sprint(i+1) + " " + kraftee.Name + " (" + fmt.Sprint(kraftee.AllCount) + ")\n"
	}
	lbs += "\n"

	lbs += "## Time ##\n"
	sort.Slice(kStats, func(i, j int) bool { return kStats[i].AllMovingSeconds > kStats[j].AllMovingSeconds })
	rank, err = findKrafteeRankinStatsList(kStats, k)
	if err != nil {
		log.Fatal(err)
	}
	for i, kraftee := range kStats {
		if i > rank {
			break // Only show kraftees with data
		}
		lbs += fmt.Sprint(i+1) + " " + kraftee.Name + " (" + secToHMS(kraftee.AllMovingSeconds) + ")\n"
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
			lbs += fmt.Sprint(i+1) + " " + kraftee.Name + " (" + fmt.Sprintf("%.1f", metersToMiles(kraftee.RunMeters)) + " mi.)\n"
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
			lbs += fmt.Sprint(i+1) + " " + kraftee.Name + " (" + secToHMS(kraftee.RunMovingSeconds) + ")\n"
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
			lbs += fmt.Sprint(i+1) + " " + kraftee.Name + " (" + fmt.Sprintf("%.1f", metersToMiles(kraftee.RideMeters)) + " mi.)\n"
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
			lbs += fmt.Sprint(i+1) + " " + kraftee.Name + " (" + secToHMS(kraftee.RideMovingSeconds) + ")\n"
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
