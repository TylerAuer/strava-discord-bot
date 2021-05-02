package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/dustin/go-humanize"
)

func buildActivityPost(a ActivityDetails, k Kraftee) string {
	id := fmt.Sprint(a.ID)
	url := "https://www.strava.com/activities/" + id

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

	movTime := "Time:               " + secToHMS(a.MovingTime)

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
		k.First + " posted: *" + title + "* " + emoji() + "\n" +
		msg() +
		"\n" +
		"**Activity Summary**\n" +
		// "*Where you stood on the leaderboard when this activity was first posted*\n" +
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
		"\n" + url +
		"\n\n" +
		buildLeaderboardStatus(a, k)
}

func buildLeaderboardStatus(a ActivityDetails, k Kraftee) string {
	startOfWeek := getStartOfWeekInUnixTime()
	kStats, _ := getAllKrafteeStatsSince(startOfWeek)

	lbs := "**Leaderboard Update**\n"
	lbs += "*Where you stood on the leaderboard when this activity was first posted*\n"
	lbs += "```"

	sort.Slice(kStats, func(i, j int) bool { return kStats[i].AllCount > kStats[j].AllCount })
	rank, err := findKrafteeRankinStatsList(kStats, k)
	if err != nil {
		log.Fatal(err)
	}

	lbs += "Activities:    " + humanize.Ordinal(rank+1) + " with " + fmt.Sprint(kStats[rank].AllCount)

	// Only add relative values if not in first
	if rank > 0 {
		lbs += " ("

		// Only compare to person in front of you if not in first or second
		if rank >= 2 {
			rankInFront := rank - 1
			countBehindPersonInFront := kStats[rankInFront].AllCount - kStats[rank].AllCount
			lbs += fmt.Sprint(countBehindPersonInFront) + " behind " + humanize.Ordinal(rank) + "; "
		}
		countBehindFirst := kStats[0].AllCount - kStats[rank].AllCount
		lbs += fmt.Sprint(countBehindFirst) + " behind " + humanize.Ordinal(1)

		lbs += ")"
	}

	lbs += "\n"

	// Total time standings
	sort.Slice(kStats, func(i, j int) bool { return kStats[i].AllMovingSeconds > kStats[j].AllMovingSeconds })
	rank, err = findKrafteeRankinStatsList(kStats, k)
	if err != nil {
		log.Fatal(err)
	}

	lbs += "Time:          " + humanize.Ordinal(rank+1) + " with " + secToHMS(kStats[rank].AllMovingSeconds)

	// Only add relative values if not in first
	if rank > 0 {
		lbs += " ("
		// Only compare to person in front of you if not in first or second
		if rank >= 2 {
			secsBehindPersonAhead := kStats[rank-1].AllMovingSeconds - kStats[rank].AllMovingSeconds
			lbs += secToHMS(secsBehindPersonAhead) + " behind " + humanize.Ordinal(rank) + "; "
		}
		secsBehindFirst := kStats[0].AllMovingSeconds - kStats[rank].AllMovingSeconds
		lbs += secToHMS(secsBehindFirst) + " behind " + humanize.Ordinal(1)

		lbs += ")"
	}

	if a.Type == "Run" {
		lbs += "\n\n" // Spacer

		sort.Slice(kStats, func(i, j int) bool { return kStats[i].RunMeters > kStats[j].RunMeters })
		rank, err = findKrafteeRankinStatsList(kStats, k)
		if err != nil {
			log.Fatal(err)
		}
		lbs += "Run Distance:  " + humanize.Ordinal(rank+1) + " with " + fmt.Sprintf("%.1f", metersToMiles(kStats[rank].RunMeters)) + " mi."
		if rank > 0 {
			lbs += " ("

			if rank >= 2 {
				metersBehindPersonAhead := kStats[rank-1].RunMeters - kStats[rank].RunMeters
				lbs += fmt.Sprintf("%.1f", metersToMiles(metersBehindPersonAhead)) + " mi. behind " + humanize.Ordinal(rank) + "; "
			}
			secsBehindFirst := kStats[0].RunMeters - kStats[rank].RunMeters
			lbs += fmt.Sprintf("%.1f", metersToMiles(secsBehindFirst)) + " mi. behind " + humanize.Ordinal(1)

			lbs += ")"
		}

		lbs += "\n"

		sort.Slice(kStats, func(i, j int) bool { return kStats[i].RunMovingSeconds > kStats[j].RunMovingSeconds })
		rank, err = findKrafteeRankinStatsList(kStats, k)
		if err != nil {
			log.Fatal(err)
		}
		lbs += "Run Time:      " + humanize.Ordinal(rank+1) + " with " + secToHMS(kStats[rank].AllMovingSeconds)

		if rank > 0 {
			lbs += " ("
			if rank >= 2 {
				secsBehindPersonAhead := kStats[rank-1].RunMovingSeconds - kStats[rank].RunMovingSeconds
				lbs += secToHMS(secsBehindPersonAhead) + " behind " + humanize.Ordinal(rank) + "; "
			}
			secsBehindFirst := kStats[0].RunMovingSeconds - kStats[rank].RunMovingSeconds
			lbs += secToHMS(secsBehindFirst) + " behind " + humanize.Ordinal(1)

			lbs += ")"
		}
	}

	if a.Type == "Ride" {
		lbs += "\n\n" // Spacer
		sort.Slice(kStats, func(i, j int) bool { return kStats[i].RideMeters > kStats[j].RideMeters })
		rank, err = findKrafteeRankinStatsList(kStats, k)
		if err != nil {
			log.Fatal(err)
		}
		lbs += "Ride Distance: " + humanize.Ordinal(rank+1) + " with " + fmt.Sprintf("%.1f", metersToMiles(kStats[rank].RideMeters)) + " mi."
		if rank > 0 {
			lbs += " ("
			if rank >= 2 {
				metersBehindPersonAhead := kStats[rank-1].RideMeters - kStats[rank].RideMeters
				lbs += fmt.Sprintf("%.1f", metersToMiles(metersBehindPersonAhead)) + " mi. behind " + humanize.Ordinal(rank) + "; "
			}
			secsBehindFirst := kStats[0].RideMeters - kStats[rank].RideMeters
			lbs += fmt.Sprintf("%.1f", metersToMiles(secsBehindFirst)) + " mi. behind " + humanize.Ordinal(1)

			lbs += ")"
		}

		lbs += "\n"

		sort.Slice(kStats, func(i, j int) bool { return kStats[i].RideMovingSeconds > kStats[j].RideMovingSeconds })
		rank, err = findKrafteeRankinStatsList(kStats, k)
		if err != nil {
			log.Fatal(err)
		}
		lbs += "Ride Time:     " + humanize.Ordinal(rank+1) + " with " + secToHMS(kStats[rank].AllMovingSeconds)

		if rank > 0 {
			lbs += " ("

			if rank >= 2 {
				secsBehindPersonAhead := kStats[rank-1].RideMovingSeconds - kStats[rank].RideMovingSeconds
				lbs += secToHMS(secsBehindPersonAhead) + " behind " + humanize.Ordinal(rank) + "; "
			}
			secsBehindFirst := kStats[0].RideMovingSeconds - kStats[rank].RideMovingSeconds
			lbs += secToHMS(secsBehindFirst) + " behind " + humanize.Ordinal(1)

			lbs += ")"
		}

	}

	lbs += "```"
	return lbs
}

func findKrafteeRankinStatsList(kStats []Stats, k Kraftee) (int, error) {
	for i, kInList := range kStats {
		if kInList.ID == k.StravaId {
			return i, nil
		}
	}
	return 0, errors.New("did not find kraftee in list of Kraftee stats")
}
