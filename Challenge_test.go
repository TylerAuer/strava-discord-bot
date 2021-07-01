package main

import (
	"testing"
	"time"
)

func TestFormatDateWithOffsetIntoMonthDayYearKey(t *testing.T) {
	// Noon yields same day
	got := formatDateKeyWithOffsetToPacifc(time.Date(2021, time.June, 30, 12, 0, 0, 0, time.UTC))
	expected := "June-30-2021"
	if got != expected {
		t.Errorf("Got %v but wanted %v", got, expected)
	}

	// Handles Pacific time to UTC offset during PDT. Offset is -7 hrs for the date used.
	got = formatDateKeyWithOffsetToPacifc(time.Date(2021, time.June, 30, 6, 59, 59, 999, time.UTC))
	expected = "June-29-2021"
	if got != expected {
		t.Errorf("Got %v but wanted %v", got, expected)
	}
	got = formatDateKeyWithOffsetToPacifc(time.Date(2021, time.June, 30, 7, 0, 0, 0, time.UTC))
	expected = "June-30-2021"
	if got != expected {
		t.Errorf("Got %v but wanted %v", got, expected)
	}

	// Handles Pacific time to UTC offset during PST. Offset is -8 hrs for the date used.
	got = formatDateKeyWithOffsetToPacifc(time.Date(2021, time.January, 30, 7, 59, 59, 999, time.UTC))
	expected = "January-29-2021"
	if got != expected {
		t.Errorf("Got %v but wanted %v", got, expected)
	}
	got = formatDateKeyWithOffsetToPacifc(time.Date(2021, time.January, 30, 8, 0, 0, 0, time.UTC))
	expected = "January-30-2021"
	if got != expected {
		t.Errorf("Got %v but wanted %v", got, expected)
	}
}
