package main

import (
	"fmt"
	"strings"
)

func parseActivityStatsIntoPost(a ActivityDetails, k Kraftee) string {
	id := fmt.Sprint(a.ID)
	url := "https://www.strava.com/activities/" + id
	dist := "Distance:        " + fmt.Sprintf("%.1f", metersToMiles(a.Distance)) + " miles"
	elev := "Elevation:       " + fmt.Sprintf("%.1f", metersToFeet(a.TotalElevationGain)) + " feet gained"
	movTime := "Time:            " + secondsToHoursMinsSeconds(a.MovingTime)
	paceInSecondsPerMile := float64(a.MovingTime) / metersToMiles(a.Distance)
	pace := secondsToMinSec(paceInSecondsPerMile) + " per mile"
	relativeEffort := func() string {
		if a.SufferScore == 0 {
			return ""
		}
		return "Relative Effort: " + fmt.Sprint(a.SufferScore) + "\n"
	}()
	cals := func() string {
		if a.Calories == 0 {
			return ""
		}
		return "Calories:        " + fmt.Sprint(a.Calories) + "\n"
	}()

	return "" +
		"*" + k.First + " just logged a " + strings.ToLower(a.Type) + "*\n\n" +
		"**STATS**" +
		"```" +
		dist + "\n" +
		movTime + " @ " + pace + "\n" +
		elev + "\n" +
		relativeEffort +
		cals +
		"```" +
		"\n" + url
}
