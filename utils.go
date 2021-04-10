package main

import (
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

	var sString string
	if secs < 10 {
		sString = "0" + strconv.Itoa(secs)
	} else {
		sString = strconv.Itoa(secs)
	}

	var mString string
	if mins < 10 {
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
