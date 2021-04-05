package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

type MyEvent struct {
	Name string `json:"name"`
}

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	getStravaActivitiesAndPostToDiscord()

	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func getStravaActivitiesAndPostToDiscord() {
	godotenv.Load() // Load env vars from ./.env

	st := getStravaAccessToken()

	la := getLatestGroupActivities(st, 5)

	str := fmt.Sprintf("%#v", la)

	postToDiscord(str)
}
