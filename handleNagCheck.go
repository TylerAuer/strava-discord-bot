package main

import (
	"fmt"
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
	msg += "You are pitiful and lazy. You have not logged a workout in a while. Let's go!\n"
	msg += getRandomGifOfInspiration()

	postToDiscord(dg, msg)

}

func checkIfLazy(k Kraftee, lazyChan chan string) {
	// Only check Kraftees who have opted in
	if k.daysBeforeNag == 0 {
		fmt.Println(k.First, "hasn't opted into nagging")
		lazyChan <- ""
		return
	}

	dateToStartCheckFrom := getNDaysAgoInUnixtime(k.daysBeforeNag)
	activities := getActivitiesSince(dateToStartCheckFrom, k)

	if len(activities) == 0 {
		fmt.Println(k.First, "is lazy")
		lazyChan <- k.First
	} else {
		fmt.Println(k.First, "is not lazy")
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
