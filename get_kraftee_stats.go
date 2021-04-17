package main

import (
	"fmt"
)

func getAllKrafteeStats(startEpochTime int64) ([]Stats, []ActivityDetails) {
	var kStatsList []Stats           // Holds each Kraftees stats for comparison
	var allActList []ActivityDetails // Holds all activities for group stats computation

	kChan := make(chan Stats)
	allChan := make(chan []ActivityDetails)

	for _, k := range krafteesByStravaId {
		go getOneKrafteesStats(startEpochTime, k, kChan, allChan)
	}

	// Handle incoming channel messages

	for i := 0; i < 2*len(krafteesByStravaId); i++ {
		select {
		case newKrafteeStats := <-kChan:
			fmt.Println("Received kChan message")
			kStatsList = append(kStatsList, newKrafteeStats)
		case newListOfActivities := <-allChan:
			fmt.Println("Received allChan message")
			allActList = append(allActList, newListOfActivities...)
		}
	}

	return kStatsList, allActList
}

func getOneKrafteesStats(t int64, k Kraftee, kChan chan Stats, allChan chan []ActivityDetails) {
	kActList := getActivitiesSince(t, k)
	allChan <- kActList

	kStats := buildStatsFromActivityList(k.FullName(), k.StravaId, kActList)
	kChan <- kStats

	fmt.Println("Finished " + k.FullName())
}
