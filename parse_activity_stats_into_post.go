package main

import (
	"fmt"
	"strings"
)

func parseActivityStatsIntoPost(a ActivityDetails, k Kraftee) string {
	emojis := map[string]string{
		"walk":            "🚶🚶🚶",
		"ride":            "🚴‍♂️🚴‍♂️🚴‍♂️",
		"swim":            "🏊‍♂️🏊‍♂️🏊‍♂️",
		"weight training": "🏋️💪🏋️💪",
	}

	id := fmt.Sprint(a.ID)
	url := "https://www.strava.com/activities/" + id

	dist := "Distance:        " + fmt.Sprintf("%.2f", metersToMiles(a.Distance)) + " miles"

	elev := "Elevation:       " + fmt.Sprintf("%.0f", metersToFeet(a.TotalElevationGain)) + "' gained"

	movTime := "Time:            " + secondsToHoursMinsSeconds(a.MovingTime)

	paceInSecondsPerMile := float64(a.MovingTime) / metersToMiles(a.Distance)
	pace := "Pace:            " + secondsToMinSec(paceInSecondsPerMile) + " per mile"

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

	achievementCount := func() string {
		if a.AchievementCount == 0 {
			return ""
		}
		return "Achievement Count: " + fmt.Sprint(a.AchievementCount) + "\n"
	}()

	return "" +
		k.First + " just logged a " + strings.ToLower(a.Type) + emojis[strings.ToLower(a.Type)] + "\n\n" +
		"```" +
		dist + "\n" +
		movTime + "\n" +
		pace + "\n" +
		elev + "\n" +
		relativeEffort +
		cals +
		achievementCount +
		"```" +
		"\n" + url
}
