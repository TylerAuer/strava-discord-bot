package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Totals struct {
	Count            int64   `json:"count"`
	Distance         float64 `json:"distance"`
	MovingTime       int64   `json:"moving_time"`
	ElapsedTime      int64   `json:"elapsed_time"`
	ElevationGain    float64 `json:"elevation_gain"`
	AchievementCount int64   `json:"achievement_count"`
}

type Stats struct {
	YtdRunTotals  Totals `json:"ytd_run_totals"`
	YtdRideTotals Totals `json:"ytd_ride_totals"`
	YtdSwimTotals Totals `json:"ytd_swim_totals"`
}

func getAthleteStats(k Kraftee) AthleteStats {
	fmt.Println("Getting athlete stats for " + k.fullName())
	athleteStatsUrl := "https://www.strava.com/api/v3/athletes/" + k.StravaId + "/stats"
	authHeader := "Bearer " + k.StravaAccessToken

	// Build request; include authHeader
	req, err := http.NewRequest("GET", athleteStatsUrl, nil)
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

	stats := AthleteStats{}

	err = json.NewDecoder(resp.Body).Decode(&stats)
	if err != nil {
		log.Fatal(err)
	}

	return stats
}
