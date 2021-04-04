package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() // Load env vars from ./.env

	st := getStravaAccessToken()

	la := getLatestGroupActivity(st)

	str := fmt.Sprintf("%#v", la)

	postToDiscord("Progress boys! Check this out!")
	postToDiscord(str)
	postToDiscord("It's running on my local machine, but it's doing most of the hard work already!")
}
