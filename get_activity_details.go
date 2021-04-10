package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func getActivityDetails(id string, k Kraftee) ActivityDetails {
	fmt.Println("Getting details of activity with ID: " + id)

	url := "https://www.strava.com/api/v3/activities/" + id

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

	stats := ActivityDetails{}

	err = json.NewDecoder(resp.Body).Decode(&stats)
	if err != nil {
		log.Fatal(err)
	}

	return stats
}

type ActivityDetails struct {
	ResourceState int `json:"resource_state"`
	Athlete       struct {
		ID            int `json:"id"`
		ResourceState int `json:"resource_state"`
	} `json:"athlete"`
	Name               string    `json:"name"`
	Distance           float64   `json:"distance"`
	MovingTime         int       `json:"moving_time"`
	ElapsedTime        int       `json:"elapsed_time"`
	TotalElevationGain float64   `json:"total_elevation_gain"`
	Type               string    `json:"type"`
	ID                 int64     `json:"id"`
	StartDate          time.Time `json:"start_date"`
	StartDateLocal     time.Time `json:"start_date_local"`
	Timezone           string    `json:"timezone"`
	UtcOffset          float64   `json:"utc_offset"`
	StartLatlng        []float64 `json:"start_latlng"`
	AchievementCount   int       `json:"achievement_count"`
	Map                struct {
		ID              string `json:"id"`
		Polyline        string `json:"polyline"`
		ResourceState   int    `json:"resource_state"`
		SummaryPolyline string `json:"summary_polyline"`
	} `json:"map"`
	AverageSpeed               float64     `json:"average_speed"`
	MaxSpeed                   float64     `json:"max_speed"`
	AverageCadence             float64     `json:"average_cadence"`
	HasHeartrate               bool        `json:"has_heartrate"`
	AverageHeartrate           float64     `json:"average_heartrate"`
	MaxHeartrate               float64     `json:"max_heartrate"`
	HeartrateOptOut            bool        `json:"heartrate_opt_out"`
	DisplayHideHeartrateOption bool        `json:"display_hide_heartrate_option"`
	ElevHigh                   float64     `json:"elev_high"`
	ElevLow                    float64     `json:"elev_low"`
	PrCount                    int         `json:"pr_count"`
	TotalPhotoCount            int         `json:"total_photo_count"`
	HasKudoed                  bool        `json:"has_kudoed"`
	SufferScore                float64     `json:"suffer_score"`
	Description                interface{} `json:"description"`
	Calories                   float64     `json:"calories"`
	Photos                     struct {
		Primary interface{} `json:"primary"`
		Count   int         `json:"count"`
	} `json:"photos"`
}
