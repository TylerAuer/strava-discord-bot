package main

import (
	"fmt"
	"sort"
)

func buildLeaderboardPost(lb Leaderboard) string {
	lbPost := "Leaderboard\n"
	lbPost += "```"

	AllMovingSeconds := "Total Time\n"
	sort.Slice(lb, func(i, j int) bool { return lb[i].AllMovingSeconds > lb[j].AllMovingSeconds })
	for i, k := range lb[:9] {
		if k.AllMovingSeconds > 0 {
			AllMovingSeconds += medal[i] + " " + secToHMS(k.AllMovingSeconds) + " " + k.Name + "\n"
		}
	}
	lbPost += AllMovingSeconds

	ActivityCount := "\nActivity Count\n"
	sort.Slice(lb, func(i, j int) bool { return lb[i].AllCount > lb[j].AllCount })
	for i, k := range lb[:9] {
		if k.AllCount > 0 {
			ActivityCount += medal[i] + " " + fmt.Sprint(k.AllCount) + " " + k.Name + "\n"
		}
	}
	lbPost += ActivityCount

	RunDistance := "\nRun Distance " + emojis["run"] + "\n"
	sort.Slice(lb, func(i, j int) bool { return lb[i].RunMeters > lb[j].RunMeters })
	for i, k := range lb[:9] {
		if k.RunMeters > 0 {
			RunDistance += medal[i] + " " + fmt.Sprintf("%.2f", metersToMiles(k.RunMeters)) + " mi. " + k.Name + "\n"
		}
	}
	lbPost += RunDistance

	RunTime := "\nRun Time " + emojis["run"] + "\n"
	sort.Slice(lb, func(i, j int) bool { return lb[i].RunMovingSeconds > lb[j].RunMovingSeconds })
	for i, k := range lb[:9] {
		if k.RunMovingSeconds > 0 {
			RunTime += medal[i] + " " + secToHMS(k.RunMovingSeconds) + " " + k.Name + "\n"
		}
	}
	lbPost += RunTime

	RunElev := "\nRun Elevation Gain " + emojis["run"] + "\n"
	sort.Slice(lb, func(i, j int) bool { return lb[i].RunElevationGain > lb[j].RunElevationGain })
	for i, k := range lb[:9] {
		if k.RunElevationGain > 0 {
			RunElev += medal[i] + " " + fmt.Sprintf("%.0f", metersToFeet(k.RunElevationGain)) + "' " + k.Name + "\n"
		}
	}
	lbPost += RunElev

	RideDistance := "\nRide Distance " + emojis["ride"] + "\n"
	sort.Slice(lb, func(i, j int) bool { return lb[i].RideMeters > lb[j].RideMeters })
	for i, k := range lb[:9] {
		if k.RideMeters > 0 {
			RideDistance += medal[i] + " " + fmt.Sprintf("%.2f", metersToMiles(k.RideMeters)) + " mi. " + k.Name + "\n"
		}
	}
	lbPost += RideDistance

	RideTime := "\nRide Time " + emojis["ride"] + "\n"
	sort.Slice(lb, func(i, j int) bool { return lb[i].RideMovingSeconds > lb[j].RideMovingSeconds })
	for i, k := range lb[:9] {
		if k.RideMovingSeconds > 0 {
			RideTime += medal[i] + " " + secToHMS(k.RideMovingSeconds) + " " + k.Name + "\n"
		}
	}
	lbPost += RideTime

	RideElev := "\nRide Elevation Gain " + emojis["ride"] + "\n"
	sort.Slice(lb, func(i, j int) bool { return lb[i].RideElevationGain > lb[j].RideElevationGain })
	for i, k := range lb[:9] {
		if k.RideElevationGain > 0 {
			RideElev += medal[i] + " " + fmt.Sprintf("%.0f", metersToFeet(k.RideElevationGain)) + "' " + k.Name + "\n"
		}
	}
	lbPost += RideElev
	lbPost += "```"

	return lbPost
}
