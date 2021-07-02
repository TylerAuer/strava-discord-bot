package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/jinzhu/now"
)

func handleJessicaDailyUpdate() {
	jessica := Kraftee{"Jessica", "Auer", "JESSICA", "31271517", "", 0}
	goal := 100.0

	today := time.Now().Day()
	daysLeftInMonth := now.EndOfMonth().Day() - today + 1

	jessicaActiviesSince := getActivitiesSince(now.BeginningOfMonth().Unix(), jessica)

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
