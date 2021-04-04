package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Member struct {
	Resource_state int    `json:"resource_state"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	Membership     string `json:"membership"`
	Admin          bool   `json:"admin"`
	Owner          bool   `json:"owner"`
}

func getGroupMemberList(st string) []Member {
	fmt.Println("Getting group member list")

	cid := os.Getenv("STRAVA_CLUB_ID")
	url := "https://www.strava.com/api/v3/clubs/" + cid + "/members/?access_token=" + st

	res, resErr := http.Get(url)
	if resErr != nil {
		log.Fatal(resErr)
	}
	defer res.Body.Close()

	body, bodyErr := ioutil.ReadAll(res.Body)
	if bodyErr != nil {
		log.Fatal(bodyErr)
	}

	ml := []Member{}
	parseErr := json.Unmarshal(body, &ml)
	if parseErr != nil {
		log.Fatal(parseErr)
	}

	fmt.Println("Member list received")
	return ml
}
