package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func getActivitiesSince(s int64, k Kraftee) ActivityList {
	fmt.Println("Getting activity history for " + k.FullName())
	// Grab the stats for Kraftee

	url := "https://www.strava.com/api/v3/athlete/activities?page=1&per_page=200&after=" + fmt.Sprint(s)

	authHeader := "Bearer " + k.GetStravaAccessToken()

	// Build request; include authHeader
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", authHeader)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	activityList := ActivityList{}

	err = json.NewDecoder(resp.Body).Decode(&activityList)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Got " + fmt.Sprint(len(activityList)) + " activities for " + k.FullName())

	return activityList
}
