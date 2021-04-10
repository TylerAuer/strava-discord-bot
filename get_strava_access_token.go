package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

func getStravaAccessToken(k Kraftee) string {
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
