package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getActivityDetails(actId string, st string) {
	fmt.Println("Getting activity details")

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
