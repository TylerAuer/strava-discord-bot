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
	dist := "**Dist:** " + fmt.Sprintf("%.1f", metersToMiles(a.Distance))
	elev := "**Elev:** +" + fmt.Sprintf("%.1f", metersToFeet(a.TotalElevationGain))
	movTime := "**Time:** " + secondsToHoursMinsSeconds(a.MovingTime)
	paceInSecondsPerMile := float64(a.MovingTime) / metersToMiles(a.Distance)
	pace := secondsToMinSec(paceInSecondsPerMile) + " per mile"
	relativeEffort := func() string {
		if a.SufferScore == 0 {
			return ""
		} else {
			return "**RelEff:** " + fmt.Sprint(a.SufferScore) + "\n"
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
