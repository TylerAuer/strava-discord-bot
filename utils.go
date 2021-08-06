package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/rivo/uniseg"
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

func getStartOfWeekInUnixTime() int64 {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset).Unix()
}

func getNDaysAgoInUnixtime(days int) int64 {
	return time.Now().AddDate(0, 0, -1*days).Unix()
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
		if uniseg.GraphemeClusterCount(paddedString) >= length {
			return paddedString
		}
		paddedString = " " + paddedString
	}
}

func padRight(s string, length int) string {
	paddedString := s
	for {
		if uniseg.GraphemeClusterCount(paddedString) >= length {
			return paddedString
		}
		paddedString = paddedString + " "
	}
}

type TableRow []string
type Table []TableRow

// Builds a custom-aligned table where
// the first column is left aligned and the rest are right aligned
func (t Table) composeRightAlignedTable(gutterSize int, isAllActivitiesTable bool) string {
	var gutter string
	for i := 0; i <= gutterSize; i++ {
		gutter += " "
	}

	// Check that all table rows are the same length
	for _, row := range t {
		if len(row) != len(t[0]) {
			log.Fatal("Table rows are not of equal length")
		}
	}

	// Get the maximum length of the columns
	var colMaxLengths []int
	for range t[0] {
		// Populate colMaxLengths with the 0s to avoid out-of-bounds errors
		colMaxLengths = append(colMaxLengths, 0)
	}
	for _, row := range t {
		for colIndex, cellContents := range row {
			cellLength := uniseg.GraphemeClusterCount(cellContents)
			if cellLength > colMaxLengths[colIndex] {
				colMaxLengths[colIndex] = cellLength
			}
		}
	}

	// Compose the tableString so the first column is left-aligned and the rest are right aligned
	var table string
	for _, row := range t {
		colCount := len(row)
		for j, cell := range row {
			if j == 0 {
				// First column
				table += padRight(cell, colMaxLengths[j]) + gutter
			} else if j == colCount-1 {
				// Last column
				table += padLeft(cell, colMaxLengths[j])
			} else {
				// Not the first or the last column (pad and add gutter)
				colWidth := colMaxLengths[j]
				// Discords emojis are 2 chars wide in the fixed width font
				// So to avoid issues in the all activites table, we need to make empty columns
				// an extra character wide
				if cell == "" && isAllActivitiesTable {
					colWidth++
				}
				table += padLeft(cell, colWidth) + gutter
			}
		}
		table += "\n"
	}

	return table
}

func getEmojiSequence(name string, length int) string {
	var emojis = map[string]string{
		"walk":           "🚶",
		"hike":           "🥾",
		"run":            "🏃",
		"ride":           "🚴",
		"swim":           "🏊",
		"weighttraining": "🏋️",
		"ebikeride":      "🛵",
	}
	var str string
	var e string

	if emoji, ok := emojis[name]; ok {
		e = emoji
	} else {
		e = "🥵"
	}

	// Make <length> copies of emoji
	for i := 0; i < length; i++ {
		str += e
	}
	return str
}

func getRankEmoji(zeroIndexedRank int) string {
	var medals = map[int]string{
		0: "🥇",
		1: "🥈",
		2: "🥉",
		3: "4️⃣",
		4: "5️⃣",
		5: "6️⃣",
		6: "7️⃣",
		7: "8️⃣",
		8: "9️⃣",
		9: "🔟",
	}

	if medal, ok := medals[zeroIndexedRank]; ok {
		return medal
	} else {
		return "🙁"
	}
}
