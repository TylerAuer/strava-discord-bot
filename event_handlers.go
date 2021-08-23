package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/jinzhu/now"
)

type WebhookData struct {
	ObjectType string `json:"object_type"` // "athlete" or "activity"
	ObjectId   int    `json:"object_id"`   // id of athlete or activity
	AspectType string `json:"aspect_type"` // "create" "update" "delete"
	OwnerId    int    `json:"owner_id"`    // ID of the athlete who owns the event
	EventTime  int    `json:"event_time"`
}

type Cron struct {
	Type string `json:"type"`
}

type StravaValidateSubscriptionBody struct {
	HubChallenge string `json:"hub.challenge"`
}

func handleStravaWebhook(body string) {
	fmt.Println("Strava Post Content:" + body)

	var b WebhookData
	err := json.Unmarshal([]byte(body), &b)
	if err != nil {
		log.Fatal(err)
	}

	if b.ObjectType == "activity" && (b.AspectType == "create" || b.AspectType == "update") {
		fmt.Println("Handling new activity with ID: " + fmt.Sprint(b.ObjectId))

		k := krafteesByStravaId[fmt.Sprint(b.OwnerId)]
		idStr := fmt.Sprint(b.ObjectId)

		ad := getActivityDetails(idStr, k)
		isWeeklyWorkoutChallengeActivity := ad.isWeeklyWorkoutChallenge()

		// Upsert activity details to MongoDB
		initMongo().Upsert(ad)

		isCreateType := b.AspectType == "create"

		if isWeeklyWorkoutChallengeActivity {
			ad.postOrUpdateWeeklyWorkoutChallengePost(isCreateType)
		} else {
			ad.postOrUpdateActivityPost(isCreateType)
		}
	} else {
		fmt.Println("webhook was none of the following 1) activity 2) create aspect_type 3) update aspect_type")
	}
}

func handleWeeklyUpdatePost() {
	// Find 7 days ago in unix epoch time
	secsToLookBack := int64(7 * 24 * 60 * 60)
	epochTime := time.Now().Unix()
	startInEpochTime := epochTime - secsToLookBack

	krafteeCount := len(krafteesByStravaId)

	listOfKrafteeStats, listOfEveryActivity := fetchAllKrafteeStatsSince(startInEpochTime)

	groupStats := listOfEveryActivity.buildStats("All", "")
	groupStatsPost := groupStats.printGroupStats()

	leaderboardPost := listOfKrafteeStats.composeLeaderboardPost()

	msg := "**Weekly Update**\n"
	msg += "*Here's a summary for " + fmt.Sprint(krafteeCount) + " kraftees over the last week*"
	msg += "\n\n" + groupStatsPost
	msg += "\n\n" + leaderboardPost

	dg := getDiscord()
	dg.Close()

	dg.post(msg)
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

func handleNagCheck() {
	dg := getActiveDiscordSession()
	defer dg.Close()

	var lazyKraftees []string
	lazyChan := make(chan string)

	// Initiate checks to see if a Kraftee is lazy
	for _, k := range krafteesByStravaId {
		go checkIfLazy(k, lazyChan)
	}

	// Wait for a response from every Kraftee request
	for range krafteesByStravaId {
		newLazyKraftee := <-lazyChan
		if newLazyKraftee != "" {
			lazyKraftees = append(lazyKraftees, newLazyKraftee)
		}
	}

	// Nothing to post if no one is lazy
	if len(lazyKraftees) == 0 {
		return
	}

	msg := strings.Join(lazyKraftees, ", ") + "\n"
	msg += "You are pitiful and lazy. You have not logged a workout in a while. Let's go!\n"
	msg += getRandomGifOfInspiration()

	postToDiscord(dg, msg)
}

func checkIfLazy(k Kraftee, lazyChan chan string) {
	// Only check Kraftees who have opted in
	if k.daysBeforeNag == 0 {
		fmt.Println(k.SafeFirstName(), "hasn't opted into nagging")
		lazyChan <- ""
		return
	}

	dateToStartCheckFrom := getNDaysAgoInUnixtime(k.daysBeforeNag)
	activities := fetchKrafteeActivitiesSince(dateToStartCheckFrom, k)

	if len(activities) == 0 {
		fmt.Println(k.SafeFirstName(), "is lazy")
		lazyChan <- k.SafeFirstName()
	} else {
		fmt.Println(k.SafeFirstName(), "is not lazy")
		lazyChan <- ""
	}

}

func getRandomGifOfInspiration() string {
	// Ensure a random seed
	rand.Seed(time.Now().UnixNano())

	gifsOfInspiration := []string{
		"https://media.giphy.com/media/6XA99Q0nPSXyU/giphy.gif",      // Breaking Bad
		"https://media.giphy.com/media/mcH0upG1TeEak/giphy.gif",      // Old School
		"https://media.giphy.com/media/rfskmSvktqSoo/giphy.gif",      // TSwift
		"https://media.giphy.com/media/IVUT6ZrmwluToUnb1F/giphy.gif", // Daniel Jones
		"https://media.giphy.com/media/pVkmGyqYRt4qY/giphy.gif",      // Cat on sofa
		"https://media.giphy.com/media/w9DgGBsFfU9OM/giphy.gif",      // Edna Mode
		"https://media.giphy.com/media/kaCdrXPD9hkZzsEEJn/giphy.gif", // Groundhog day
		"https://media.giphy.com/media/Ob7p7lDT99cd2/giphy.gif",      // GOT Shame
		"https://media.giphy.com/media/d5mI2F3MxCTJu/giphy.gif",      // Sloth
		"https://media.giphy.com/media/DOId45m6FAEUw/giphy.gif",      // Major Payne
		"https://media.giphy.com/media/gKO8OzGBZMSxOBaxjp/giphy.gif", // Big hero 6
		"https://media.giphy.com/media/3oEduUsg7Ad7gXH4CA/giphy.gif", // Fat Joe
		"https://media.giphy.com/media/SnplG1QPbcULwdj1wR/giphy.gif", // Parks and recs
	}
	randomIndex := rand.Intn(len(gifsOfInspiration))

	return gifsOfInspiration[randomIndex]
}

func handleJessicaDailyUpdate() {
	jessica := Kraftee{"Jessica", "Auer", "JESSICA", "31271517", "", 0}
	goal := 100.0

	today := time.Now().Day()
	daysLeftInMonth := now.EndOfMonth().Day() - today + 1

	jessicaActiviesSince := fetchKrafteeActivitiesSince(now.BeginningOfMonth().Unix(), jessica)

	jStats := jessicaActiviesSince.buildStats("Jessica", jessica.StravaId)

	if metersToMiles(jStats.WalkOrHikeMeters) > goal {
		sendSmsToJessica("You beat your goal! So proud!")
	}

	msg := ""
	msg += fmt.Sprint(jStats.WalkOrHikeCount) + " walks over "
	msg += fmt.Sprintf("%.1f", metersToMiles(jStats.WalkOrHikeMeters)) + " mi in "
	msg += fmt.Sprintf("%.1f", secondsToHours(jStats.WalkOrHikeMovingSeconds)) + " hrs | "
	msg += fmt.Sprintf("%.1f", (goal-metersToMiles(jStats.WalkOrHikeMeters))/float64(daysLeftInMonth)) + " mi/day for "
	msg += fmt.Sprint(daysLeftInMonth) + " days"

	sendSmsToJessica(msg)
}

func sendSmsToJessica(message string) {
	fmt.Println("Sending sms message: ", message)

	sender := os.Getenv("TWILIO_PHONE_NUMBER")
	recipient := os.Getenv("JESSICA_PHONE")
	sid := os.Getenv("TWILIO_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + sid + "/Messages.json"

	// Pack up the data for our message
	msgData := url.Values{}
	msgData.Set("To", recipient)
	msgData.Set("From", sender)
	msgData.Set("Body", message)
	msgDataReader := *strings.NewReader(msgData.Encode())

	// Create HTTP request client
	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(sid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make HTTP POST request and return message SID
	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status)
	}
}

func handleCron(body string) Cron {
	var b Cron
	err := json.Unmarshal([]byte(body), &b)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func (c Cron) executeCronJobBasedOnType() {
	if c.Type == "nag" {
		handleNagCheck()
	}
	if c.Type == "weekly_update" {
		handleWeeklyUpdatePost()
	}
	if c.Type == "jessica_daily_update" {
		handleJessicaDailyUpdate()
	}
}
