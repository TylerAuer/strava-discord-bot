package main

import (
	"strconv"
)

func metersToMiles(m int) float64 {
	return float64(m) * 0.000621371
}

func metersToFeet(m int) float64 {
	return float64(m) * 3.28084
}

func secondsToHours(s int) float64 {
	return float64(s) / 60.0 / 60.0
}

func secondsToHoursMinsSeconds(s int) string {
	// floatHours := float64(s) / 60.0 / 60.0

	// intHours := int(floatHours)
	// floatMins := math.Remainder(floatHours, float64(1)) * 60
	// intMins := int(floatMins)
	// intSeconds := int(math.Remainder(floatMins, float64(1)) * 60)
	// fmt.Println(floatMins)

	secs := s % 60
	mins := s / 60 % 60
	hours := s / 60 / 60 % 60

	return strconv.Itoa(hours) + ":" + strconv.Itoa(mins) + ":" + strconv.Itoa(secs)

}
