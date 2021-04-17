package main

import (
	"fmt"
	"sort"
)

func buildLeaderboardPost(sList []Stats) string {
	medal := map[int]string{
		0: "🥇",
		1: "🥈",
		2: "🥉",
	}

	DiscordPost := "Leaderboard\n"
	DiscordPost += "```"

	AllMovingSeconds := "Total Time\n"
	sort.Slice(sList, func(i, j int) bool { return sList[i].AllMovingSeconds > sList[j].AllMovingSeconds })
	for i, k := range sList[:3] {
		if k.AllMovingSeconds > 0 {
			AllMovingSeconds += medal[i] + " " + secondsToHoursMinsSeconds(k.AllMovingSeconds) + " " + k.Name + "\n"
		}
	}
	DiscordPost += AllMovingSeconds

	ActivityCount := "\nActivity Count\n"
	sort.Slice(sList, func(i, j int) bool { return sList[i].AllCount > sList[j].AllCount })
	for i, k := range sList[:3] {
		if k.AllCount > 0 {
			ActivityCount += medal[i] + " " + fmt.Sprint(k.AllCount) + " " + k.Name + "\n"
		}
	}
	DiscordPost += ActivityCount

	RunDistance := "\nRun Distance " + emojis["run"] + "\n"
	sort.Slice(sList, func(i, j int) bool { return sList[i].RunMeters > sList[j].RunMeters })
	for i, k := range sList[:3] {
		if k.RunMeters > 0 {
			RunDistance += medal[i] + " " + fmt.Sprintf("%.2f", metersToMiles(k.RunMeters)) + " mi. " + k.Name + "\n"
		}
	}
	DiscordPost += RunDistance

	RunTime := "\nRun Time " + emojis["run"] + "\n"
	sort.Slice(sList, func(i, j int) bool { return sList[i].RunMovingSeconds > sList[j].RunMovingSeconds })
	for i, k := range sList[:3] {
		if k.RunMovingSeconds > 0 {
			RunTime += medal[i] + " " + secondsToHoursMinsSeconds(k.RunMovingSeconds) + " " + k.Name + "\n"
		}
	}
	DiscordPost += RunTime

	RunElev := "\nRun Elevation Gain " + emojis["run"] + "\n"
	sort.Slice(sList, func(i, j int) bool { return sList[i].RunElevationGain > sList[j].RunElevationGain })
	for i, k := range sList[:3] {
		if k.RunElevationGain > 0 {
			RunElev += medal[i] + " " + fmt.Sprintf("%.0f", metersToFeet(k.RunElevationGain)) + "' " + k.Name + "\n"
		}
	}
	DiscordPost += RunElev

	RideDistance := "\nRide Distance " + emojis["ride"] + "\n"
	sort.Slice(sList, func(i, j int) bool { return sList[i].RideMeters > sList[j].RideMeters })
	for i, k := range sList[:3] {
		if k.RideMeters > 0 {
			RideDistance += medal[i] + " " + fmt.Sprintf("%.2f", metersToMiles(k.RideMeters)) + " mi. " + k.Name + "\n"
		}
	}
	DiscordPost += RideDistance

	RideTime := "\nRide Time " + emojis["ride"] + "\n"
	sort.Slice(sList, func(i, j int) bool { return sList[i].RideMovingSeconds > sList[j].RideMovingSeconds })
	for i, k := range sList[:3] {
		if k.RideMovingSeconds > 0 {
			RideTime += medal[i] + " " + secondsToHoursMinsSeconds(k.RideMovingSeconds) + " " + k.Name + "\n"
		}
	}
	DiscordPost += RideTime

	RideElev := "\nRide Elevation Gain " + emojis["ride"] + "\n"
	sort.Slice(sList, func(i, j int) bool { return sList[i].RideElevationGain > sList[j].RideElevationGain })
	for i, k := range sList[:3] {
		if k.RideElevationGain > 0 {
			RideElev += medal[i] + " " + fmt.Sprintf("%.0f", metersToFeet(k.RideElevationGain)) + "' " + k.Name + "\n"
		}
	}
	DiscordPost += RideElev
	DiscordPost += "```"

	return DiscordPost
}
