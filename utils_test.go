package main

import (
	"testing"
)

func TestMetersToMiles(t *testing.T) {
	got := metersToMiles(1000)
	expect := 0.621371
	if got != expect {
		t.Errorf("Converting 1000m to miles; got %f but wanted %f", got, expect)
	}
}

func TestComposeTwoColumnTable(t *testing.T) {
	padding := "  "

	var expected string
	expected += "Left         " + padding + "         Right\n"
	expected += "LeftButLonger" + padding + "         Right\n"
	expected += "Left         " + padding + "RightButLonger\n"

	d := []TwoDimensionalTableData{
		{left: "Left", right: "Right"},
		{left: "LeftButLonger", right: "Right"},
		{left: "Left", right: "RightButLonger"},
	}
	got := composeTwoColumnTable(d)
	if got != expected {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}
