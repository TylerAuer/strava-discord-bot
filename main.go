package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() // Load env vars from ./.env

	st := getStravaAccessToken()

	la := getLatestGroupActivities(st, 5)

	str := fmt.Sprintf("%#v", la)

	postToDiscord(str)
}
