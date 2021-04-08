package main

import "fmt"

type WebhookData struct {
	ObjectType string `json:"object_type"` // "athlete" or "activity"
	ObjectId   string `json:"object_id"`   // id of athlete or activity
	AspectType string `json:"aspect_type"` // "create" "update" "delete"
	OwnerId    int    `json:"owner_id"`    // ID of the athlete who owns the event
	EventTime  int    `json:"event_time"`
}

func handleStravaWebhook(body string) {
	fmt.Println("Strava Post Content:" + body)
}
