package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

type StravaValidateSubscriptionBody struct {
	HubChallenge string `json:"hub.challenge"`
}

func handleStravaSubscriptionChallenge(q map[string]string) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Returning hub.challenge to Strava: " + q["hub.challenge"])

	body := StravaValidateSubscriptionBody{
		HubChallenge: q["hub.challenge"],
	}

	b, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	r := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(b),
	}

	fmt.Println(r)
	return r, nil
}
