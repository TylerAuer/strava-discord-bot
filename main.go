package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting")
	godotenv.Load()
	fmt.Println("Loaded env vars")

	// Decide what to execute based on where things are running
	production := os.Getenv("PRODUCTION")
	if production != "FALSE" {
		fmt.Println("Passing handleRequest to lambda.Start")
		lambda.Start(handleLambda)
	} else {
		fmt.Println("Invoking handleLocal")
		handleLocal()
	}
}

func handleLambda(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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

func handleLocal() {
	ta := Kraftee{"Tyler", "Auer", "Ugly Stick", "2007", "TYLER", "20419783", ""}
	ta.StravaAccessToken = getStravaAccessToken(ta)

	a := getActivityDetails("5088461939", ta)
	p := parseActivityStatsIntoPost(a, ta)

	postToDiscord(p)
}
