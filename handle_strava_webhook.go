package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type WebhookData struct {
	ObjectType string `json:"object_type"` // "athlete" or "activity"
	ObjectId   int    `json:"object_id"`   // id of athlete or activity
	AspectType string `json:"aspect_type"` // "create" "update" "delete"
	OwnerId    int    `json:"owner_id"`    // ID of the athlete who owns the event
	EventTime  int    `json:"event_time"`
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
		a := getActivityDetails(idStr, k)
		p := buildActivityPost(a, k)

		postOrUpdateActivity(idStr, p, b.AspectType, a, k)
	} else {
		fmt.Println("webhook was none of the following 1) activity 2) create aspect_type 3) update aspect_type")
	}
}
