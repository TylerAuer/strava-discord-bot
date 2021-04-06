package main

import (
	"github.com/joho/godotenv"
)

type MyEvent struct {
	Name string `json:"name"`
}

func main() {
	// lambda.Start(HandleRequest)
	godotenv.Load() // Load env vars from ./.env

	ta := Kraftee{"Tyler", "Auer", "Ugly Stick", "2007", "TYLER", "20419783", ""}

	kraftees := []Kraftee{ta}

	// TODO: Use go routines here
	for _, k := range kraftees {
		k.StravaAccessToken = getStravaAccessToken(k)
		stats := getAthleteStats(k)

		postToDiscord(stats.YtdRunsTotalsString())
	}
}
