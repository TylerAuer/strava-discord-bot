package main

import (
	"math"
	"strconv"
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

	if hours == 0 {
		return strconv.Itoa(mins) + ":" + strconv.Itoa(secs)
	}

	return strconv.Itoa(hours) + ":" + strconv.Itoa(mins) + ":" + strconv.Itoa(secs)

}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func secondsToMinSec(p float64) string {
	mins := int(p / 60.0)                    // mins with decimal truncated
	fractionalMins := p/60.0 - float64(mins) // fractional minute left after truncating
	secs := int(fractionalMins * 60.0)

	return strconv.Itoa(mins) + ":" + strconv.Itoa(secs)
}
