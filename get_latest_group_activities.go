package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Athlete struct {
	Firstname     string `json:"firstname"`
	Lastname      string `json:"lastname"`
	ResourceState int64  `json:"resource_state"`
}

type ActivityBasic struct {
	ResourceState int64   `json:"resource_state"`
	Athlete       Athlete `json:"athlete"`
	Name          string  `json:"name"`
	Distance      float64 `json:"distance"`
	Type          string  `json:"type"`
	MovingTime    int64   `json:"moving_time"`
	ElapsedTime   int64   `json:"elapsed_time"`
	ElevationGain float64 `json:"total_elevation_gain"`
}

func getLatestGroupActivities(st string, count int16) []ActivityBasic {
	fmt.Println("Getting most recent group activity")

	cid := os.Getenv("STRAVA_CLUB_ID")
	url := "https://www.strava.com/api/v3/clubs/" + cid + "/activities?per_page=" + fmt.Sprint(count) + "&access_token=" + st

	res, resErr := http.Get(url)
	if resErr != nil {
		log.Fatal(resErr)
	}
	defer res.Body.Close()

	body, bodyErr := ioutil.ReadAll(res.Body)
	if bodyErr != nil {
		log.Fatal(bodyErr)
	}

	la := []ActivityBasic{}
	parseErr := json.Unmarshal(body, &la)
	if parseErr != nil {
		log.Fatal(parseErr)
	}

	fmt.Println("Retrieved most recent" + fmt.Sprint(count) + "group activities")
	return la
}
