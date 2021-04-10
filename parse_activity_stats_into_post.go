package main

import (
	"fmt"
	"strings"
)

func parseActivityStatsIntoPost(a ActivityDetails, k Kraftee) string {
	emojis := map[string]string{
		"walk":            "🚶🚶🚶",
		"run":             "🏃‍♂️🏃‍♂️🏃‍♂️",
		"ride":            "🚴‍♂️🚴‍♂️🚴‍♂️",
		"swim":            "🏊‍♂️🏊‍♂️🏊‍♂️",
		"weight training": "🏋️💪🏋️💪",
	}

	id := fmt.Sprint(a.ID)
	url := "https://www.strava.com/activities/" + id

	dist := "Distance:           " + fmt.Sprintf("%.2f", metersToMiles(a.Distance)) + " miles"

	elev := "Elevation Gain:     " + fmt.Sprintf("%.0f", metersToFeet(a.TotalElevationGain)) + "'"

	movTime := "Time:               " + secondsToHoursMinsSeconds(a.MovingTime)

	paceInSecondsPerMile := float64(a.MovingTime) / metersToMiles(a.Distance)
	pace := "Pace:               " + secondsToMinSec(paceInSecondsPerMile) + " per mile"

	relativeEffort := func() string {
		if a.SufferScore == 0 {
			return ""
		}
		return "Relative Effort:    " + fmt.Sprint(a.SufferScore) + "\n"
	}()

	cals := func() string {
		if a.Calories == 0 {
			return ""
		}
		return "Calories:           " + fmt.Sprint(a.Calories) + "\n"
	}()

	achievementCount := func() string {
		if a.AchievementCount == 0 {
			return ""
		}
		return "Achievement Count:  " + fmt.Sprint(a.AchievementCount) + "\n"
	}()

	avgHeartRate := func() string {
		if a.AverageHeartrate == 0 {
			return ""
		}
		return "Average Heart Rate: " + fmt.Sprint(a.AverageHeartrate) + "\n"
	}()

	return "" +
		k.First + " just logged a " + strings.ToLower(a.Type) + emojis[strings.ToLower(a.Type)] + "\n\n" +
		"```" +
		dist + "\n" +
		movTime + "\n" +
		pace + "\n" +
		elev + "\n" +
		avgHeartRate +
		relativeEffort +
		cals +
		achievementCount +
		"```" +
		"\n" + url
}
