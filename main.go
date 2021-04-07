package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

// This will run when the image is running on AWS Lambda
func handleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Invoking handleRequest")

	defaultResponse := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Not much to see here!",
	}

	httpMethod := req.HTTPMethod
	fmt.Println("HTTP Method: " + httpMethod)

	switch httpMethod {
	case "GET":
		fmt.Printf("Handling GET request which should be subscription validation from Strava")
		return handleStravaSubscriptionChallenge(req.QueryStringParameters)
	case "POST":
		fmt.Printf("Handling POST request which should be webhook event from Strava")
		handleStravaWebhook(req.Body)
	}

	return defaultResponse, nil
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
