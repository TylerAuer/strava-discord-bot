package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

func getStravaAccessToken() string {
	fmt.Println("Requesting access token from Strava")

	client_id := os.Getenv("STRAVA_CLIENT_ID")
	client_secret := os.Getenv("STRAVA_CLIENT_SECRET")
	refresh_token := os.Getenv("STRAVA_REFRESH_TOKEN")

	stravaTokenUrl := "https://www.strava.com/api/v3/oauth/token"

	reqBody := url.Values{
		"client_id":     {client_id},
		"client_secret": {client_secret},
		"refresh_token": {refresh_token},
		"grant_type":    {"refresh_token"},
	}

	res, err := http.PostForm(stravaTokenUrl, reqBody)
	if err != nil {
		log.Fatal(err)
	}

	var result map[string]string

	json.NewDecoder(res.Body).Decode(&result)

	fmt.Println("Strava access token received")

	return result["access_token"]
}
