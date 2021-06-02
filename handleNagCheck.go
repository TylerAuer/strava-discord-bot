package main

import (
	"math/rand"
	"strings"
	"time"
)

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
	msg += "You are pitiful and lazy. You have not logged a workout in a while. Man up!\n"
	msg += getRandomGifOfInspiration()

	postToDiscord(dg, msg)

}

func checkIfLazy(k Kraftee, lazyChan chan string) {
	// Only check Kraftees who have opted in
	if k.daysBeforeNag == 0 {
		lazyChan <- ""
		return
	}

	dateToStartCheckFrom := getNDaysAgoInUnixtime(k.daysBeforeNag)
	activities := getActivitiesSince(dateToStartCheckFrom, k)

	if len(activities) == 0 {
		lazyChan <- k.First
	} else {
		lazyChan <- ""
	}

}

func getRandomGifOfInspiration() string {
	// Ensure a random seed
	rand.Seed(time.Now().UnixNano())

	gifsOfInspiration := []string{
		"https://media.giphy.com/media/6XA99Q0nPSXyU/giphy.gif",
		"https://media.giphy.com/media/mcH0upG1TeEak/giphy.gif",
		"https://media.giphy.com/media/rfskmSvktqSoo/giphy.gif",
		"https://media.giphy.com/media/IVUT6ZrmwluToUnb1F/giphy.gif",
		"https://media.giphy.com/media/pVkmGyqYRt4qY/giphy.gif",
		"https://media.giphy.com/media/w9DgGBsFfU9OM/giphy.gif",
		"https://media.giphy.com/media/kaCdrXPD9hkZzsEEJn/giphy.gif",
		"https://media.giphy.com/media/Ob7p7lDT99cd2/giphy.gif",
		"https://media.giphy.com/media/d5mI2F3MxCTJu/giphy.gif",
	}
	randomIndex := rand.Intn(len(gifsOfInspiration))

	return gifsOfInspiration[randomIndex]
}
