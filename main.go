package main

import (
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() // Load env vars from ./.env

	stravaToken := getStravaAccessToken()

}
