package main

import (
	"fmt"
	"strings"
)

func buildActivityPost(a ActivityDetails, k Kraftee) string {
	title := a.Name

	msg := func() string {
		switch a.Type {
		case "WeightTraining":
			return "Get swole!\n"
		default:
			return ""
		}
	}

	emoji := func() string {
		if emojis, ok := emojis[strings.ToLower(a.Type)]; ok {
			return emojis
		}
		fmt.Println("No emoji for: " + strings.ToLower(a.Type))
		return emojis["fallback"]
	}

	dist := func() string {
		if a.Distance > 0 {
			return "Dist:    " + fmt.Sprintf("%.2f", metersToMiles(a.Distance)) + " miles\n"
		} else {
			return ""
		}
	}

	elev := func() string {
		if a.TotalElevationGain > 0 {
			return "Elev:    +" + fmt.Sprintf("%.0f", metersToFeet(a.TotalElevationGain)) + "'\n"
		} else {
			return ""
		}
	}

	movTime := "Time:    " + secToHMS(a.MovingTime)

	pace := func() string {
		if a.Distance > 0 {
			paceInSecondsPerMile := float64(a.MovingTime) / metersToMiles(a.Distance)
			return "Pace:    " + secondsToMinSec(paceInSecondsPerMile) + " per mile\n"
		} else {
			return ""
		}
	}

	relativeEffort := func() string {
		if a.SufferScore == 0 {
			return ""
		}
		return "RE:      " + fmt.Sprint(a.SufferScore) + "\n"
	}()

	cals := func() string {
		if a.Calories == 0 {
			return ""
		}
		return "Cals:    " + fmt.Sprint(a.Calories) + "\n"
	}()

	avgHeartRate := func() string {
		if a.AverageHeartrate == 0 {
			return ""
		}
		return "AVG HR:  " + fmt.Sprint(a.AverageHeartrate) + " bpm\n"
	}()

	return "" +
		k.First + " logged a " + emoji() + "\n" +
		msg() +
		"\n*" + title + "*\n" +
		"\n" +
		"**This Activity**\n" +
		// "*Where you stood on the leaderboard when this activity was first posted*\n" +
		"```\n" +
		dist() +
		movTime + "\n" +
		pace() +
		elev() +
		avgHeartRate +
		relativeEffort +
		cals +
		"```" +
		"\n"
}
