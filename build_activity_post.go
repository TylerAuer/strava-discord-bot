package main

import (
	"fmt"
	"strings"
)

func buildActivityPost(a ActivityDetails, k Kraftee) string {
	id := fmt.Sprint(a.ID)
	url := "https://www.strava.com/activities/" + id

	fmt.Println(a.Type)

	msg := func() string {
		switch a.Type {
		case "WeightTraining":
			return "Get swole!\n"
		default:
			return ""
		}
	}

	dist := func() string {
		if a.Distance > 0 {
			return "Distance:           " + fmt.Sprintf("%.2f", metersToMiles(a.Distance)) + " miles\n"
		} else {
			return ""
		}
	}

	elev := func() string {
		if a.TotalElevationGain > 0 {
			return "Elevation Gain:     " + fmt.Sprintf("%.0f", metersToFeet(a.TotalElevationGain)) + "'\n"
		} else {
			return ""
		}
	}

	movTime := "Time:               " + secondsToHoursMinsSeconds(a.MovingTime)

	pace := func() string {
		if a.Distance > 0 {
			paceInSecondsPerMile := float64(a.MovingTime) / metersToMiles(a.Distance)
			return "Pace:               " + secondsToMinSec(paceInSecondsPerMile) + " per mile\n"
		} else {
			return ""
		}
	}

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
		msg() +
		"```" +
		dist() +
		movTime + "\n" +
		pace() +
		elev() +
		avgHeartRate +
		relativeEffort +
		cals +
		achievementCount +
		"```" +
		"\n" + url
}
