package main

import "fmt"

func handleNagCheck() {
	// Load all of the data for x amount of time
	// Check if they have at least 1 activity in that data set
	// If not, send a nag message

	dg := getActiveDiscordSession()
	defer dg.Close()

	for _, k := range krafteesByStravaId {
		if k.daysBeforeNag != 0 {
			dateToStartCheckFrom := getNDaysAgoInUnixtime(k.daysBeforeNag)
			activities := getActivitiesSince(dateToStartCheckFrom, k)
			if len(activities) == 0 {
				msg := k.First + ", it's been more than "
				msg += fmt.Sprint(k.daysBeforeNag) + " since your last workout. Man up!\n"
				msg += "https://media.giphy.com/media/dZcJvBQL503SW5fErx/giphy.gif"

				postToDiscord(dg, msg)
			}
		}
	}
}
