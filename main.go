package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

var krafteesByStravaId = map[string]Kraftee{
	"20419783": {"Tyler", "Auer", "TYLER", "20419783", ""},
	"80996402": {"Jamie", "Quella", "Q", "80996402", ""},
	"80485980": {"Bryan", "Eckelmann", "BRYAN", "80485980", ""},
	"23248014": {"Fred", "Brasz", "FRED", "23248014", ""},
	// "2102360":  {"Brian", "Munroe", "PHYS", "2102360", ""},
	// "60682578": {"Tom", "Samuelson", "TOM", "60682578", ""},
	// "65626950": {"Alex", "Hogan", "HOGAN", "65626950", ""},
	// "65753450": {"Zach", "Grossman", "ZACH", "65753450", ""},
	// "80341128": {"Conor", "Quinn", "CONOR", "80341128", ""},
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
	secsToLookBack := int64(7 * 24 * 60 * 60)
	epochTime := time.Now().Unix()
	startInEpochTime := epochTime - secsToLookBack

	a := getActivitiesSince(startInEpochTime, krafteesByStravaId["20419783"])
	s := compileStatsFromActivities(a)

	prettyPrintStruct(s)

}
