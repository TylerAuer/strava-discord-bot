package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

var krafteesByStravaId = map[string]Kraftee{
	"20419783": {"Tyler", "Auer", "TYLER", "20419783", ""},
	"80996402": {"Jamie", "Quella", "Q", "80996402", ""},
	"80485980": {"Bryan", "Eckelmann", "BRYAN", "80485980", ""},
	"23248014": {"Fred", "Brasz", "FRED", "23248014", ""},
	"83356822": {"Larry", "Dworkin", "SMOOTH", "83356822", ""},
	"2102360":  {"Brian", "Munroe", "PHYS", "2102360", ""},
	"81799070": {"Norman", "Nicolson", "NORMAN", "81799070", ""},
	"65753450": {"Zach", "Grossman", "ZACH", "65753450", ""},
	"65626950": {"Alex", "Hogan", "HOGAN", "65626950", ""},
	// "60682578": {"Tom", "Samuelson", "TOM", "60682578", ""},
	// "80341128": {"Conor", "Quinn", "CONOR", "80341128", ""},
	// "82860978": {"Andre", "Martinez", "DRE", "82860978", ""},
}

var emojis = map[string]string{
	"walk":           "🚶🚶🚶",
	"run":            "🏃‍♂️🏃‍♂️🏃‍♂️",
	"ride":           "🚴‍♂️🚴‍♂️🚴‍♂️",
	"swim":           "🏊‍♂️🏊‍♂️🏊‍♂️",
	"weighttraining": "🏋️💪🏋️💪",
}

func main() {
	fmt.Println("Starting")
	godotenv.Load()
	fmt.Println("Loaded env vars")

	// Decide what to execute based on where things are running.
	// Production is set to false locally and true on AWS Lambda
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

	purpose := os.Getenv("PURPOSE")

	if purpose == "NEW_POSTS" {
		switch httpMethod {
		case "GET":
			fmt.Printf("Handling GET request which should be subscription validation from Strava")
			return handleStravaSubscriptionChallenge(req.QueryStringParameters)
		case "POST":
			fmt.Printf("Handling POST request which should be webhook event from Strava")
			handleStravaWebhook(req.Body)
		}
	} else if purpose == "WEEKLY_UPDATES" {
		handleWeeklyUpdatePost()
	}

	return defaultResponse, nil
}

func handleLocal() {
	defer duration(track("handleLocal"))
	// handleWeeklyUpdatePost()
	handleStravaWebhook(`{
    "aspect_type": "create",
    "event_time": 1618623328,
    "object_id": 5139372673,
    "object_type": "activity",
    "owner_id": 2102360,
    "subscription_id": 188592,
    "updates": {}
}`)
}
