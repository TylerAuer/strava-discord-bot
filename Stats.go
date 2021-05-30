package main

import "fmt"

type Stats struct {
	Name string
	ID   string

	AllCount         int
	AllMovingSeconds int
	Heartbeats       int
	MaxHeartRate     int
	Calories         int

	RunCount         int
	RunMovingSeconds int
	RunMeters        float64
	RunElevationGain float64

	RideCount         int
	RideMovingSeconds int
	RideMeters        float64
	RideElevationGain float64

	WalkOrHikeCount         int
	WalkOrHikeMovingSeconds int
	WalkOrHikeMeters        float64
	WalkOrHikeElevationGain float64
}

func (gs Stats) printGroupStats() string {
	post := "All Kraftees\n"
	post += "```"

	post += "Activities: " + fmt.Sprint(gs.AllCount) + "\n"
	post += "Total Time: " + secToHMS(gs.AllMovingSeconds) + "\n"
	totalElevGain := gs.RunElevationGain + gs.RideElevationGain + gs.WalkOrHikeElevationGain
	post += "Total Elevation: " + fmt.Sprintf("%.0f", metersToFeet(totalElevGain)) + " ft\n"
	post += "Total Heartbeats: " + fmt.Sprint(gs.Heartbeats/1000) + "k \n"

	post += "\n"

	post += "Runs: " + fmt.Sprint(gs.RunCount) + "\n"
	post += "Rides: " + fmt.Sprint(gs.RideCount) + "\n"

	post += "\n"

	post += "Run Time: " + secToHMS(gs.RunMovingSeconds) + "\n"
	post += "Ride Time: " + secToHMS(gs.RideMovingSeconds) + "\n"

	post += "\n"

	post += "Run Distance: " + fmt.Sprintf("%.1f", metersToMiles(gs.RunMeters)) + " mi\n"
	post += "Ride Distance: " + fmt.Sprintf("%.1f", metersToMiles(gs.RideMeters)) + " mi\n"

	post += "```"

	return post
}

func (s Stats) printKrafteeStats(k Kraftee) string {
	post := "**" + k.First + "'s Weekly Stats** @ post time\n"
	post += "```"
	post += "Activities: " + fmt.Sprint(s.AllCount) + "\n"
	post += "Total Time: " + secToHMS(s.AllMovingSeconds) + "\n"
	totalElevGain := s.RunElevationGain + s.RideElevationGain + s.WalkOrHikeElevationGain
	post += "Total Elevation: " + fmt.Sprintf("%.0f", metersToFeet(totalElevGain)) + " ft\n"
	post += "Total Heartbeats: " + fmt.Sprint(s.Heartbeats/1000) + "k \n"
	if s.RunCount > 0 || s.RideCount > 0 || s.WalkOrHikeCount > 0 {
		post += "\n"
	}
	if s.RunCount > 0 {
		post += "Runs: " + fmt.Sprint(s.RunCount) + "\n"
		post += "Run Time: " + secToHMS(s.RunMovingSeconds) + "\n"
		post += "Run Distance: " + fmt.Sprintf("%.1f", metersToMiles(s.RunMeters)) + " mi\n"
		post += "\n"
	}
	if s.RideCount > 0 {
		post += "Rides: " + fmt.Sprint(s.RideCount) + "\n"
		post += "Ride Time: " + secToHMS(s.RideMovingSeconds) + "\n"
		post += "Ride Distance: " + fmt.Sprintf("%.1f", metersToMiles(s.RideMeters)) + " mi\n"
		post += "\n"
	}
	if s.WalkOrHikeCount > 0 {
		post += "Walks: " + fmt.Sprint(s.WalkOrHikeCount) + "\n"
		post += "Walk Time: " + secToHMS(s.WalkOrHikeCount) + "\n"
		post += "Walk Distance: " + fmt.Sprintf("%.1f", metersToMiles(s.WalkOrHikeMeters)) + " mi\n"
	}

	post += "```"

	return post
}
