package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/now"
)

type WeeklyChallenge struct {
	GoalKind    string // max/minTime, max/minReps
	Title       string
	Description string
}

// These should be indexed by an ID, that way they can be reused and previous bests can be found
func getChallengeByDate(dateKey string) WeeklyChallenge {
	wcs := map[string]WeeklyChallenge{
		// Dates MUST be for the monday of each week.
		"July-5-2021":  {"maxReps", "10 min AMRAP: 5 pushups, 10 squats, 20 crunches", "As many rounds as possible; squats must reach >=90 degree knee bend"},
		"June-28-2021": {"minTime", "50 burpess for time", "Full body to ground at bottom. Body fully upright for jump"},
		"June-21-2021": {"maxReps", "AMRAP pushups", "Max pushups without taking your full weight off arms"},
	}
	return wcs[dateKey]
}

// Should not be used in posts because if an old post is updated it might access the wrong challenge
func getChallengeActiveToday() WeeklyChallenge {
	now.WeekStartDay = time.Monday
	monday := now.BeginningOfWeek()
	dateKey := fmt.Sprint(monday.Month()) + "-" + fmt.Sprint(monday.Day()) + "-" + fmt.Sprint(monday.Year())
	return getChallengeByDate(dateKey)
}

func handleWeeklyWorkoutChallengeStravaWebhook(k Kraftee, ad ActivityDetails, webhook WebhookData) {
	dg := getActiveDiscordSession()
	defer dg.Close()

	challenge := getChallengeActiveToday()

	var score string

	if challenge.GoalKind == "maxReps" || challenge.GoalKind == "minReps" {
		score = ad.Description
	} else if challenge.GoalKind == "minTime" || challenge.GoalKind == "maxTime" {
		score = secToHMS(ad.MovingTime)
	} else {
		score = "Unable to find value for workout challenge score"
	}

	var msg string
	msg += k.First + " just did the Weekly Workout Challenge\n"
	msg += "\n"
	msg += "```"
	msg += "Score: " + score + "\n"
	msg += "\n"
	msg += "##" + challenge.Title + " ##\n"
	msg += "\n"
	msg += challenge.Description + "\n"
	msg += "```"
	msg += "\n"
	msg += "ID: " + fmt.Sprint(ad.ID)

	postToDiscord(dg, msg)
}
