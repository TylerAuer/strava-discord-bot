package main

import (
	"fmt"
)

func getAllKrafteeStatsSince(startEpochTime int64) (Leaderboard, ActivityList) {
	var lb Leaderboard          // Holds each Kraftees stats for comparison
	var activities ActivityList // Holds all activities for group stats computation

	lbChan := make(chan Stats)
	activityListChan := make(chan ActivityList)

	for _, k := range krafteesByStravaId {
		go getOneKrafteesStats(startEpochTime, k, lbChan, activityListChan)
	}

	// Handle incoming channel messages
	for i := 0; i < 2*len(krafteesByStravaId); i++ {
		select {
		case newKrafteeStats := <-lbChan:
			lb = append(lb, newKrafteeStats)
		case newListOfActivities := <-activityListChan:
			activities = append(activities, newListOfActivities...)
		}
	}

	return lb, activities
}

func getOneKrafteesStats(t int64, k Kraftee, lbChan chan Stats, activityListChan chan ActivityList) {
	actList := getActivitiesSince(t, k)
	activityListChan <- actList

	kStats := actList.buildStatsFromActivityList(k.First, k.StravaId)
	lbChan <- kStats

	fmt.Println("Finished " + k.FullName())
}
