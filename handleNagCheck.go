package main

import (
	"fmt"
	"math/rand"
)

func handleNagCheck() {

	dg := getActiveDiscordSession()
	defer dg.Close()

	for _, k := range krafteesByStravaId {
		if k.daysBeforeNag != 0 {
			dateToStartCheckFrom := getNDaysAgoInUnixtime(k.daysBeforeNag)
			activities := getActivitiesSince(dateToStartCheckFrom, k)
			if len(activities) == 0 {
				msg := k.First + ", it's been more than "
				msg += fmt.Sprint(k.daysBeforeNag) + " days since your last workout. Man up!\n"
				msg += getRandomGifOfInspiration()

				postToDiscord(dg, msg)
			}
		}
	}
}

func getRandomGifOfInspiration() string {
	gifsOfInspiration := []string{
		"https://media.giphy.com/media/dZcJvBQL503SW5fErx/giphy.gif",
		"https://media.giphy.com/media/6XA99Q0nPSXyU/giphy.gif",
		"https://media.giphy.com/media/mcH0upG1TeEak/giphy.gif",
		"https://media.giphy.com/media/rfskmSvktqSoo/giphy.gif",
		"https://media.giphy.com/media/IVUT6ZrmwluToUnb1F/giphy.gif",
		"https://media.giphy.com/media/pVkmGyqYRt4qY/giphy.gif",
		"https://media.giphy.com/media/w9DgGBsFfU9OM/giphy.gif",
		"https://media.giphy.com/media/kaCdrXPD9hkZzsEEJn/giphy.gif",
		"https://media.giphy.com/media/Ob7p7lDT99cd2/giphy.gif",
		"https://media.giphy.com/media/d5mI2F3MxCTJu/giphy.gif",
	}

	randomIndex := rand.Intn(len(gifsOfInspiration))

	return gifsOfInspiration[randomIndex]

}

// func getAllKrafteeStatsSince(startEpochTime int64) (Leaderboard, ActivityList) {
// 	var lb Leaderboard          // Holds each Kraftees stats for comparison
// 	var activities ActivityList // Holds all activities for group stats computation

// 	lbChan := make(chan Stats)
// 	activityListChan := make(chan ActivityList)

// 	for _, k := range krafteesByStravaId {
// 		go getOneKrafteesStats(startEpochTime, k, lbChan, activityListChan)
// 	}

// 	// Handle incoming channel messages
// 	for i := 0; i < 2*len(krafteesByStravaId); i++ {
// 		select {
// 		case newKrafteeStats := <-lbChan:
// 			lb = append(lb, newKrafteeStats)
// 		case newListOfActivities := <-activityListChan:
// 			activities = append(activities, newListOfActivities...)
// 		}
// 	}

// 	return lb, activities
// }

// func getOneKrafteesStats(t int64, k Kraftee, lbChan chan Stats, activityListChan chan ActivityList) {
// 	actList := getActivitiesSince(t, k)
// 	activityListChan <- actList

// 	kStats := actList.buildStats(k.First, k.StravaId)
// 	lbChan <- kStats

// 	fmt.Println("Finished " + k.FullName())
// }
