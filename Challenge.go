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
			"Full body to ground at bottom. Body fully upright for jump. Tracks moving time of activity.",
			"minTime",
		},
		"Pain Server": {
			"Pain Server",
			"AMRAP pushups",
			"Max pushups without taking your full weight off arms. Reps in description.",
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
func findChallengeByMondayDate(dateKey string) *Challenge {
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

func formatDateKeyWithOffsetToPacifc(t time.Time) string {
	// Our group considers day changes to occur at midnight pacific. However, most dates are stored
	// in UTC. So, when determining the date key for activity lookups, this should offset to the LA
	// time zone.
	pst, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		panic(err)
	}
	pstTime := t.In(pst)
	return fmt.Sprint(pstTime.Month()) + "-" + fmt.Sprint(pstTime.Day()) + "-" + fmt.Sprint(pstTime.Year())
}

func getMonthDayYearStringOfCurrentWeek() string {
	now.WeekStartDay = time.Monday
	monday := now.BeginningOfWeek()
	return formatDateKeyWithOffsetToPacifc(monday)
}

func getMondayMonthDayYearStringOfDate(t time.Time) string {
	now.WeekStartDay = time.Monday
	monday := now.With(t).BeginningOfWeek()
	return formatDateKeyWithOffsetToPacifc(monday)
}

// Should not be used in posts because if an old post is updated it might access the wrong challenge
func getCurrentlyActiveToday() *Challenge {
	dateKey := getMonthDayYearStringOfCurrentWeek()
	return findChallengeByMondayDate(dateKey)
}

	return getChallengeMondayDate(dateKey)

func announceWeeklyChallenge() {
	d := getDiscord()
	defer d.Close()

	c := getCurrentlyActiveToday()
	msg := c.composeChallengeAnnouncement()
	d.post(msg)
}
