package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/now"
)

type Challenge struct {
	Name             string
	ShortDescription string
	LongDescription  string
	GoalKind         string // max/minTime, max/minReps
}

func getChallengeById(id string) *Challenge {
	challenges := map[string]Challenge{
		// Dates MUST be for the monday of each week.
		"Poker": {
			"Poker",
			"10 min AMRAP: 5 pushups, 10 squats, 20 crunches",
			"As many rounds as possible; squats must reach >=90 degree knee bend. Reps in description.",
			"maxReps",
		},
		"Odd Quad": {
			"Odd Quad",
			"50 burpess for time",
			"Full body to ground at bottom. Body fully upright for jump. Tracks moving time of activity",
			"minTime",
		},
		"Pain Server": {
			"Pain Server",
			"AMRAP pushups",
			"Max pushups without taking your full weight off arms. Reps in description",
			"maxReps",
		},
		"Goodrich": {
			"Goodrich",
			"1 foot balancing",
			"Balance on your left foot for as long as possible. Repeat on right. Tracks the moving time of activity.",
			"maxTime",
		},
	}
	challenge, ok := challenges[id]
	if ok {
		return &challenge
	}
	log.Fatal(id + " not found in challenges map")
	return nil
}

// These should be indexed by an ID, that way they can be reused and previous bests can be found
func getChallengeMondayDate(dateKey string) *Challenge {
	dateMap := map[string]string{
		// Dates MUST be for the monday of each week.
		"July-12-2021": "Poker",
		"July-5-2021":  "Odd Quad",
		"June-28-2021": "Pain Server",
		"June-21-2021": "Goodrich",
	}

	challengeKey, ok := dateMap[dateKey]
	if ok {
		return getChallengeById(challengeKey)
	}
	log.Fatal(dateKey + " not found in dateMap")
	return nil
}

func getMonthDayYearStringOfCurrentWeek() string {
	now.WeekStartDay = time.Monday
	monday := now.BeginningOfWeek()
	return fmt.Sprint(monday.Month()) + "-" + fmt.Sprint(monday.Day()) + "-" + fmt.Sprint(monday.Year())
}

// Should not be used in posts because if an old post is updated it might access the wrong challenge
func getChallengeActiveToday() *Challenge {
	dateKey := getMonthDayYearStringOfCurrentWeek()
	return getChallengeMondayDate(dateKey)
}
