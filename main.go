package main

import (
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() // Load env vars from ./.env

	st := getStravaAccessToken()

	// getGroupMemberList(st)
	getActivityDetails("5064060938", st)
}
