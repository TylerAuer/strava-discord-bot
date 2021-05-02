package main

import "fmt"

func buildGroupStatsPost(gs Stats) string {
	post := "All Kraftees\n"
	post += "```"
	post += "Activities: " + fmt.Sprint(gs.AllCount) + "\n"
	post += "Total Time: " + secToHMS(gs.AllMovingSeconds) + "\n"
	post += "Total Heartbeats: " + fmt.Sprint(gs.Heartbeats/1000) + "k \n"
	post += "\n"
	post += "Runs: " + fmt.Sprint(gs.RunCount) + "\n"
	post += "Rides: " + fmt.Sprint(gs.RideCount) + "\n"
	post += "\n"
	post += "Run Time: " + secToHMS(gs.RunMovingSeconds) + "\n"
	post += "Ride Time: " + secToHMS(gs.RideMovingSeconds) + "\n"
	post += "\n"
	post += "Run Distance: " + fmt.Sprintf("%.1f", metersToMiles(gs.RunMeters)) + " mi.\n"
	post += "Ride Distance: " + fmt.Sprintf("%.1f", metersToMiles(gs.RideMeters)) + " mi.\n"
	post += "```"

	return post
}
