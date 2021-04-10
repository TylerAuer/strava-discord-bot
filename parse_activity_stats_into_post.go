package main

import (
	"fmt"
	"strings"
)

/*
ID
<Name> finished a <type>
<Distance> gaining <elevation gain> in <Moving Time> at <min/mile>
*/

func parseActivityStatsIntoPost(a ActivityDetails, k Kraftee) string {
	id := fmt.Sprint(a.ID)
	url := "https://www.strava.com/activities/" + id
	dist := fmt.Sprintf("%.1f", metersToMiles(a.Distance)) + " miles travelled"
	elev := fmt.Sprintf("%.1f", metersToFeet(a.TotalElevationGain)) + "' gained"
	movTime := secondsToHoursMinsSeconds(a.MovingTime) + " moving time"
	paceInSecondsPerMile := float64(a.MovingTime) / metersToMiles(a.Distance)
	pace := secondsToMinSec(paceInSecondsPerMile) + " per mile"
	relativeEffort := func() string {
		if a.SufferScore == 0 {
			return ""
		} else {
			return "Relative Effort: " + fmt.Sprint(a.SufferScore) + "\n"
		}
	}()

	return "" +
		"*" + k.First + " just logged a " + strings.ToLower(a.Type) + "*\n\n" +
		dist + "\n" +
		elev + "\n" +
		movTime + " @ " + pace + "\n" +
		relativeEffort +
		"\n" + url
}
