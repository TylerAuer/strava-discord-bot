package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

const NAME_LENGTH = 10
const STAT_LENGTH = 7

var krafteesByStravaId = map[string]Kraftee{
	"20419783": {"Tyler", "Auer", "TYLER", "20419783", "", 1},
	"80996402": {"Jamie", "Quella", "Q", "80996402", "", 4},
	"80485980": {"Bryan", "Eckelmann", "BRYAN", "80485980", "", 0},
	"23248014": {"Fred", "Brasz", "FRED", "23248014", "", 0},
	"83356822": {"Larry", "Dworkin", "SMOOTH", "83356822", "", 0},
	"2102360":  {"Brian", "Munroe", "PHYS", "2102360", "", 0},
	"81799070": {"Norman", "Nicolson", "NORMAN", "81799070", "", 0},
	"65753450": {"Zach", "Grossman", "ZACH", "65753450", "", 0},
	"65626950": {"Alex", "Hogan", "HOGAN", "65626950", "", 0},
	"80341128": {"Conor", "Quinn", "CONOR", "80341128", "", 0},
	"83956179": {"Owen", "Simpson", "FICUS", "83956179", "", 0},
	// "60682578": {"Tom", "Samuelson", "TOM", "60682578", "", 0},
	// "82860978": {"Andre", "Martinez", "DRE", "82860978", "", 0},
}

var emojis = map[string]string{
	"walk":           "🚶🚶🚶",
	"hike":           "🥾🥾🥾",
	"run":            "🏃🏃🏃",
	"ride":           "🚴🚴🚴",
	"swim":           "🏊🏊🏊",
	"weighttraining": "🏋️💪🏋️💪",
	"fallback":       "🥵🥵🥵",
}
var medal = map[int]string{
	0: "🥇",
	1: "🥈",
	2: "🥉",
	3: "4️⃣",
	4: "5️⃣",
	5: "6️⃣",
	6: "7️⃣",
	7: "8️⃣",
	8: "9️⃣",
	9: "🔟",
}

func main() {
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

func handleLambda(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	defaultResponse := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Not much to see here!",
	}

	httpMethod := event.HTTPMethod
	fmt.Println("HTTP Method: " + httpMethod)

	// Each purpose is a different AWS Lambda function running this container
	// The env var PURPOSE customizes the behavior.
	// This allows for code reuse and simplicity
	purpose := os.Getenv("PURPOSE")

	if purpose == "NEW_POSTS" {
		switch httpMethod {
		case "GET":
			fmt.Println("Handling GET request which should be subscription validation from Strava")
			return handleStravaSubscriptionChallenge(event.QueryStringParameters)
		case "POST":
			fmt.Println("Handling POST request which should be webhook event from Strava")
			handleStravaWebhook(event.Body)
		}
	} else if purpose == "WEEKLY_UPDATES" {
		fmt.Println("Running the weekly update post")
		handleWeeklyUpdatePost()
	} else if purpose == "NAG" {
		fmt.Println("Running a nag check")
		handleNagCheck()
	} else if purpose == "JESSICA_DAILY_UPDATES" {
		fmt.Println("Sending Jessica a daily update")
		handleJessicaDailyUpdate()
	} else if purpose == "CRON" {
		fmt.Println("Handling a cron task")
		handleCron(event.Body).executeCronJobBasedOnType()
	}

	return defaultResponse, nil
}

func handleLocal() {
	defer duration(track("handleLocal"))

	// getCurrentChallenge()

	// handleJessicaDailyUpdate()

	// handleNagCheck()

	// handleStravaWebhook(`{
	// 	"aspect_type": "create",
	// 	"event_time": 1619767037,
	// 	"object_id": 5226015088,
	// 	"object_type": "activity",
	// 	"owner_id": 65626950,
	// 	"subscription_id": 188592,
	// 	"updates": {}
	// 	}`)

	// Tyler
	// handleStravaWebhook(`{
	// 	"aspect_type": "create",
	// 	"event_time": 1619767037,
	// 	"object_id": 5198828416,
	// 	"object_type": "activity",
	// 	"owner_id": 20419783,
	// 	"subscription_id": 188592,
	// 	"updates": {}
	// 	}`)

	// Bryan
	// 	handleStravaWebhook(`{
	//     "aspect_type": "create",
	//     "event_time": 1622380204,
	//     "object_id": 5383270682,
	//     "object_type": "activity",
	//     "owner_id": 80485980,
	//     "subscription_id": 188592,
	//     "updates": {}
	// }`)

	// Connor
	// handleStravaWebhook(`{
	//     "aspect_type": "create",
	//     "event_time": 1622391955,
	//     "object_id": 5384643767,
	//     "object_type": "activity",
	//     "owner_id": 80341128,
	//     "subscription_id": 188592,
	//     "updates": {}
	// }`)

	// Quella
	// handleStravaWebhook(`{
	//     "aspect_type": "create",
	//     "event_time": 1622391955,
	//     "object_id": 5397738686,
	//     "object_type": "activity",
	//     "owner_id": 80996402,
	//     "subscription_id": 188592,
	//     "updates": {}
	// }`)

	// Tyler's wwc example
	// handleStravaWebhook(`{
	//   "aspect_type": "create",
	//   "event_time": 1624338393,
	//   "object_id": 5509090759,
	//   "object_type": "activity",
	//   "owner_id": 20419783,
	//   "subscription_id": 188592,
	//   "updates": {}
	// }`)

	// CRON
	handleCron(`{"type": "jessica_daily_update"}`).executeCronJobBasedOnType()

}
