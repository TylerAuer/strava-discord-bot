package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

func metersToMiles(m float64) float64 {
	return m * 0.000621371
}

func metersToFeet(m float64) float64 {
	return m * 3.28084
}

func secondsToHours(s int) float64 {
	return float64(s) / 60.0 / 60.0
}

func secToHMS(s int) string {
	// floatHours := float64(s) / 60.0 / 60.0

	// intHours := int(floatHours)
	// floatMins := math.Remainder(floatHours, float64(1)) * 60
	// intMins := int(floatMins)
	// intSeconds := int(math.Remainder(floatMins, float64(1)) * 60)
	// fmt.Println(floatMins)

	secs := s % 60
	mins := s / 60 % 60
	hours := s / 60 / 60 % 60

	var sString string
	if secs < 10 {
		sString = "0" + strconv.Itoa(secs)
	} else {
		sString = strconv.Itoa(secs)
	}

	var mString string
	if mins < 10 && hours > 0 {
		mString = "0" + strconv.Itoa(mins)
	} else {
		mString = strconv.Itoa(mins)
	}

	if hours == 0 {
		return mString + ":" + sString
	}

	return strconv.Itoa(hours) + ":" + mString + ":" + sString

}

func secondsToMinSec(p float64) string {
	mins := int(p / 60.0)                    // mins with decimal truncated
	fractionalMins := p/60.0 - float64(mins) // fractional minute left after truncating
	secs := int(fractionalMins * 60.0)

	var secString string
	if secs < 10 {
		secString = "0" + strconv.Itoa(secs)
	} else {
		secString = strconv.Itoa(secs)
	}

	return strconv.Itoa(mins) + ":" + secString
}

func prettyPrintStruct(s interface{}) {
	j, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("%s\n", string(j))
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

func getStartOfWeekInUnixTime() (startOfWeekInUnixTime int64) {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStartInUnix := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset).Unix()

	return weekStartInUnix
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func padLeft(s string, length int) string {
	paddedString := s
	for {
		if len(paddedString) >= length {
			return paddedString
		}
		paddedString = " " + paddedString
	}
}

func padRight(s string, length int) string {
	paddedString := s
	for {
		if len(paddedString) >= length {
			return paddedString
		}
		paddedString = paddedString + " "
	}
}
