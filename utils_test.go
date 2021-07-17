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
	expected += "LeftButLonger" + padding + "  Right medium\n"
	expected += "LeftButLong  " + padding + "RightButLonger\n"

	data := TwoDimensionalTable{
		{left: "Left", right: "Right"},
		{left: "LeftButLonger", right: "Right medium"},
		{left: "LeftButLong", right: "RightButLonger"},
	}
	got := data.composeTwoColumnTable()
	if got != expected {
		t.Errorf("Expected:\n\n%s, got:\n\n%s", expected, got)
	}
}

func TestPadLeft(t *testing.T) {
	// Test a string
	expect := "     Tyler"
	got := padLeft("Tyler", 10)
	if got != expect {
		t.Errorf("Expected %s, got %s", expect, got)
	}

	// Test a sting that is the same length as the padding
	expect = "Tyler"
	got = padLeft("Tyler", 5)
	if got != expect {
		t.Errorf("Expected %s, got %s", expect, got)
	}

	// Test string with emoji
	expect = "   💩 Tyler"
	got = padLeft("💩 Tyler", 10)
	if got != expect {
		t.Errorf("Expected %s, got %s", expect, got)
	}
}

func TestPadRight(t *testing.T) {
	// Test a string
	expect := "Tyler     "
	got := padRight("Tyler", 10)
	if got != expect {
		t.Errorf("Expected %s, got %s", expect, got)
	}

	// Test a sting that is the same length as the padding
	expect = "Tyler"
	got = padRight("Tyler", 5)
	if got != expect {
		t.Errorf("Expected %s, got %s", expect, got)
	}

	// Test string with emoji
	expect = "💩 Tyler   "
	got = padRight("💩 Tyler", 10)
	if got != expect {
		t.Errorf("Expected \n%s|||, got \n%s|||", expect, got)
	}
}
