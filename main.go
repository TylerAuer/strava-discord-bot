package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

type MyEvent struct {
	Name string `json:"name"`
}

// This will run when the image is running on AWS Lambda
func handleRequest(ctx context.Context, name MyEvent) (string, error) {
	fmt.Println("Executing handleRequest")
	return "I am correctly handling this request", nil
}

// This will run locally
func handleLocal() {
	ta := Kraftee{"Tyler", "Auer", "Ugly Stick", "2007", "TYLER", "20419783", ""}

	kraftees := []Kraftee{ta}

	// TODO: Use go routines here
	for _, k := range kraftees {
		k.StravaAccessToken = getStravaAccessToken(k)
		stats := getAthleteStats(k)

		fmt.Println(stats.YtdRunsTotalsString())
		// postToDiscord(stats.YtdRunsTotalsString())
	}
}

func main() {
	fmt.Println("Starting")
	godotenv.Load()
	fmt.Println("Loaded env vars")

	// Decide what to execute based on where things are running
	production := os.Getenv("PRODUCTION")
	if production != "FALSE" {
		fmt.Println("Passing handleRequest to lambda.Start")
		lambda.Start(handleRequest)
	} else {
		handleLocal()
	}
}
