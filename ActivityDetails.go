package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func getActivityDetails(id string, k Kraftee) ActivityDetails {
	fmt.Println("Getting details of activity with ID: " + id)

	url := "https://www.strava.com/api/v3/activities/" + id

	authHeader := "Bearer " + k.GetStravaAccessToken()

	// Build request; include authHeader
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", authHeader)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	stats := ActivityDetails{}

	err = json.NewDecoder(resp.Body).Decode(&stats)
	if err != nil {
		log.Fatal(err)
	}

	return stats
}

type ActivityDetails struct {
	ResourceState int `json:"resource_state"`
	Athlete       struct {
		ID            int `json:"id"`
		ResourceState int `json:"resource_state"`
	} `json:"athlete"`
	Name               string    `json:"name"`
	Distance           float64   `json:"distance"`
	MovingTime         int       `json:"moving_time"`
	ElapsedTime        int       `json:"elapsed_time"`
	TotalElevationGain float64   `json:"total_elevation_gain"`
	Type               string    `json:"type"`
	ID                 int64     `json:"id"`
	StartDate          time.Time `json:"start_date"`
	StartDateLocal     time.Time `json:"start_date_local"`
	Timezone           string    `json:"timezone"`
	UtcOffset          float64   `json:"utc_offset"`
	StartLatlng        []float64 `json:"start_latlng"`
	AchievementCount   int       `json:"achievement_count"`
	Map                struct {
		ID              string `json:"id"`
		Polyline        string `json:"polyline"`
		ResourceState   int    `json:"resource_state"`
		SummaryPolyline string `json:"summary_polyline"`
	} `json:"map"`
	AverageSpeed               float64 `json:"average_speed"`
	MaxSpeed                   float64 `json:"max_speed"`
	AverageCadence             float64 `json:"average_cadence"`
	HasHeartrate               bool    `json:"has_heartrate"`
	AverageHeartrate           float64 `json:"average_heartrate"`
	MaxHeartrate               float64 `json:"max_heartrate"`
	HeartrateOptOut            bool    `json:"heartrate_opt_out"`
	DisplayHideHeartrateOption bool    `json:"display_hide_heartrate_option"`
	ElevHigh                   float64 `json:"elev_high"`
	ElevLow                    float64 `json:"elev_low"`
	PrCount                    int     `json:"pr_count"`
	TotalPhotoCount            int     `json:"total_photo_count"`
	HasKudoed                  bool    `json:"has_kudoed"`
	SufferScore                float64 `json:"suffer_score"`
	Description                string  `json:"description"`
	Calories                   float64 `json:"calories"`
	Photos                     struct {
		Primary interface{} `json:"primary"`
		Count   int         `json:"count"`
	} `json:"photos"`
}

func (ad ActivityDetails) krafteeWhoRecordedActivity() Kraftee {
	return krafteesByStravaId[fmt.Sprint(ad.Athlete.ID)]
}

func (ad ActivityDetails) isWeeklyWorkoutChallenge() bool {
	// Check if this activity is a weekly workout challenge
	re := `wwc\s*$` // Any string ending in wwc (ignoring trailing whitespace)
	isWWCPost, err := regexp.Match(re, []byte(ad.Name))
	if err != nil {
		log.Fatal("Regexp error: ", err)
	}

	return isWWCPost
}

func (ad ActivityDetails) getDiscordPostWithMatchingId() *discordgo.Message {
	dg := getDiscord()
	defer dg.Close()

	/**
	For each of the last 100 messages, check if it contains "ID: <activityID>". If one is found
	with a matching URL, update it.

	This is desired even if the Strava webhook type is "create" because Strava's webhook accidentally
	fires duplicate events, often.
	*/
	messagesList := dg.lastOneHundredMessages()
	re := "ID: " + fmt.Sprint(ad.ID)
	for i, msg := range messagesList {
		matched, err := regexp.Match(re, []byte(msg.Content))
		if err != nil {
			log.Fatal("Regexp error: ", err)
		}
		if matched {
			fmt.Println("Found a matching discord post with id: " + msg.ID + " which is " + fmt.Sprint(i) + " posts from the end of the thread.")
			return msg
		}
	}
	return nil
}

func (ad ActivityDetails) paceInSecondsPerMile() string {
	pace := float64(ad.MovingTime) / metersToMiles(ad.Distance)
	return secondsToMinSec(pace)
}

func (ad ActivityDetails) composeActivityPost() string {
	k := ad.krafteeWhoRecordedActivity()
	title := ad.Name

	msg := func() string {
		switch ad.Type {
		case "WeightTraining":
			return "Get swole!\n"
		default:
			return ""
		}
	}

	emoji := func() string {
		if emojis, ok := emojis[strings.ToLower(ad.Type)]; ok {
			return emojis
		}
		fmt.Println("No emoji for: " + strings.ToLower(ad.Type))
		return emojis["fallback"]
	}

	dist := func() string {
		if ad.Distance > 0 {
			return "Dist:    " + fmt.Sprintf("%.2f", metersToMiles(ad.Distance)) + " miles\n"
		} else {
			return ""
		}
	}

	elev := func() string {
		if ad.TotalElevationGain > 0 {
			return "Elev:    +" + fmt.Sprintf("%.0f", metersToFeet(ad.TotalElevationGain)) + "'\n"
		} else {
			return ""
		}
	}

	movTime := "Time:    " + secToHMS(ad.MovingTime)

	pace := func() string {
		if ad.Distance > 0 {
			return "Pace:    " + ad.paceInSecondsPerMile() + " per mile\n"
		} else {
			return ""
		}
	}

	relativeEffort := func() string {
		if ad.SufferScore == 0 {
			return ""
		}
		return "RE:      " + fmt.Sprint(ad.SufferScore) + "\n"
	}()

	cals := func() string {
		if ad.Calories == 0 {
			return ""
		}
		return "Cals:    " + fmt.Sprint(ad.Calories) + "\n"
	}()

	avgHeartRate := func() string {
		if ad.AverageHeartrate == 0 {
			return ""
		}
		return "AVG HR:  " + fmt.Sprint(ad.AverageHeartrate) + " bpm\n"
	}()

	return "" +
		k.First + " logged a " + emoji() + "\n" +
		msg() +
		"\n*" + title + "*\n" +
		"\n" +
		"**This Activity**\n" +
		// "*Where you stood on the leaderboard when this activity was first posted*\n" +
		"```\n" +
		dist() +
		movTime + "\n" +
		pace() +
		elev() +
		avgHeartRate +
		relativeEffort +
		cals +
		"```" +
		"\n"
}

func (ad ActivityDetails) composeChallengePost() string {
	k := ad.krafteeWhoRecordedActivity()
	challenge := getCurrentlyActiveToday()

	var score string

	if challenge.GoalKind == "maxReps" || challenge.GoalKind == "minReps" {
		score = ad.Description
	} else if challenge.GoalKind == "minTime" || challenge.GoalKind == "maxTime" {
		score = secToHMS(ad.MovingTime)
	} else {
		score = "Unable to find value for workout challenge score"
	}

	var msg string
	msg += k.First + " just did the " + challenge.Name + " WWC\n"
	msg += "\n"
	msg += "```"
	msg += "Score: " + score + "\n"
	msg += "\n"
	msg += "## " + challenge.ShortDescription + " ##\n"
	msg += "\n"
	msg += challenge.LongDescription + "\n"
	msg += "```"
	msg += "\n"

	return msg
}

func (ad ActivityDetails) composeLeaderboardStatusPost() string {
	k := ad.krafteeWhoRecordedActivity()

	startOfWeek := getStartOfWeekInUnixTime()
	lb, _ := fetchAllKrafteeStatsSince(startOfWeek)

	postString := "**Leaderboard** @ post time\n"
	postString += "```\n"
	postString += lb.composeActivityCountAndTimeCombinedOnActivity(&k)

	if ad.Type == "Run" {
		// postString += lb.printRunDistanceUpToKraftee(&k)
		// postString += lb.composeRunDurationUpToKraftee(&k)
		postString += lb.composeRunDistanceAndDurationCombinedOnActivity(&k)
	}

	if ad.Type == "Ride" {
		postString += lb.composeRideDistanceUpToKraftee(&k)
		postString += lb.composeRideDurationUpToKraftee(&k)
	}

	if ad.Type == "Walk" || ad.Type == "Hike" {
		postString += lb.composeWalkOrHikeDistanceUpToKraftee(&k)
		postString += lb.composeWalkOrHikeDurationUpToKraftee(&k)
	}
	postString += "```"

	postString += "\n"

	return postString
}

func (ad ActivityDetails) composePostIdentifier() string {
	return "ID: " + fmt.Sprint(ad.ID)
}

func (ad ActivityDetails) postOrUpdateActivityPost(canMakeNewPost bool) {
	dg := getDiscord()
	defer dg.Close()

	postToUpdate := ad.getDiscordPostWithMatchingId()

	if postToUpdate != nil {
		// Found a post with the matching ID. So need to grab the previous leaderboard with regex
		// and use that with the post
		fmt.Println("Updating activity post")

		// The leaderboard must lock at activity creation time, so we don't want to regenerate it
		// Instead we want to capture the old one with regex
		regexForLeaderboard := regexp.MustCompile("[*]*Leaderboard[*]* @ post time[\\w|\\W]*`{3}\n{1}")
		oldLeaderboard := string(regexForLeaderboard.Find([]byte(postToUpdate.Content)))

		// If no leaderboard was found -- because post was a WWC before -- generate a leaderboard
		if oldLeaderboard == "" {
			oldLeaderboard = ad.composeLeaderboardStatusPost()
		}

		post := ad.composeActivityPost() + oldLeaderboard + ad.composePostIdentifier()
		dg.updatePost(postToUpdate, post)
	} else if canMakeNewPost {
		// No post with matching ID found, so make a new post if allowed
		fmt.Println("Making a new activity post")

		post := ad.composeActivityPost() + ad.composeLeaderboardStatusPost() + ad.composePostIdentifier()
		dg.post(post)
	} else {
		fmt.Println("Old post wasn't found and canMakeNewPost == false. This is likely because Strava sent a duplicate `create` event ")
	}
}

func (ad ActivityDetails) postOrUpdateWeeklyWorkoutChallengePost(canMakeNewPost bool) {
	dg := getDiscord()
	defer dg.Close()

	postToUpdate := ad.getDiscordPostWithMatchingId()

	if postToUpdate != nil {
		// Found a post with the matching ID. So need to grab the previous leaderboard with regex
		// and use that with the post
		fmt.Println("Updating challenge post")

		// The leaderboard must lock at activity creation time, so we don't want to regenerate it
		// Instead we want to capture the old one with regex
		// regexForLeaderboard := regexp.MustCompile("[*]*Leaderboard[*]* @ post time[\\w|\\W]*`{3}\n{1}")
		// oldLeaderboard := string(regexForLeaderboard.Find([]byte(postToUpdate.Content)))

		post := ad.composeChallengePost() + ad.composePostIdentifier()
		dg.updatePost(postToUpdate, post)
	} else if canMakeNewPost {
		// No post with matching ID found, so make a new post if allowed
		fmt.Println("Making a new challenge post")

		post := ad.composeChallengePost() + ad.composePostIdentifier()
		dg.post(post)
	} else {
		fmt.Println("Old post wasn't found and canMakeNewPost == false. This is likely because Strava sent a duplicate `create` event ")
	}
}
