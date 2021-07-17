package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

func fetchStravaAccessToken(k Kraftee) string {
	fmt.Println("Requesting Strava Access Token for " + k.FullName())
	client_id := os.Getenv("STRAVA_CLIENT_ID")
	client_secret := os.Getenv("STRAVA_CLIENT_SECRET")
	refresh_token := os.Getenv("STRAVA_REFRESH_" + k.RefreshTokenEnvName)

	stravaTokenUrl := "https://www.strava.com/api/v3/oauth/token"

	reqBody := url.Values{
		"client_id":     {client_id},
		"client_secret": {client_secret},
		"grant_type":    {"refresh_token"},
		"refresh_token": {refresh_token},
	}

	res, err := http.PostForm(stravaTokenUrl, reqBody)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var result map[string]string

	json.NewDecoder(res.Body).Decode(&result)

	if result["access_token"] == "" {
		log.Fatal("Error getting access token for " + k.FullName())
	}
	fmt.Println("Received access token for " + k.FullName() + " of: " + result["access_token"])

	return result["access_token"]
}

func fetchAllKrafteeStatsSince(startEpochTime int64) (Leaderboard, ActivityList) {
	var lb Leaderboard          // Holds each Kraftees stats for comparison
	var activities ActivityList // Holds all activities for group stats computation

	lbChan := make(chan Stats)
	activityListChan := make(chan ActivityList)

	for _, k := range krafteesByStravaId {
		go fetchOneKrafteeStats(startEpochTime, k, lbChan, activityListChan)
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

func fetchOneKrafteeStats(t int64, k Kraftee, lbChan chan Stats, activityListChan chan ActivityList) {
	actList := fetchKrafteeActivitiesSince(t, k)
	activityListChan <- actList

	kStats := actList.buildStats(k.First, k.StravaId)
	lbChan <- kStats

	fmt.Println("Finished " + k.FullName())
}

func fetchKrafteeActivitiesSince(s int64, k Kraftee) ActivityList {
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
