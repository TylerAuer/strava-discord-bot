package main

import (
	"encoding/json"
	"log"
)

type Cron struct {
	Type string `json:"type"`
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
