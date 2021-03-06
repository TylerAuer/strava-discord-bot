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
	expect = "   π© Tyler"
	got = padLeft("π© Tyler", 10)
	if got != expect {
		t.Errorf("Expected %s, got %s", expect, got)
	}

	// Test with multi emoji emoji
	expect = "   πΆβπ«οΈ"
	got = padLeft("πΆβπ«οΈ", 4)
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
	expect = "π© Tyler   "
	got = padRight("π© Tyler", 10)
	if got != expect {
		t.Errorf("Expected \n%s, got \n%s", expect, got)
	}

	// Test string with multi emoji emoji
	expect = "πΆβπ«οΈ Tyler   "
	got = padRight("πΆβπ«οΈ Tyler", 10)
	if got != expect {
		t.Errorf("Expected \n%s, got \n%s", expect, got)
	}
}

func TestTableCompose(t *testing.T) {
	table := Table{
		TableRow{"1 Tyler", "4 mi", "2:12:03"},
		TableRow{"2 Jim", "4.7 mi", "12:03"},
		TableRow{"3 Jonathan", "14.7 mi", "2:03"},
		TableRow{"4 K", "1 mi", "1:22:03"},
	}
	expected := "1 Tyler         4 mi   2:12:03\n2 Jim         4.7 mi     12:03\n3 Jonathan   14.7 mi      2:03\n4 K             1 mi   1:22:03\n"

	got := table.composeRightAlignedTable(2, false)
	if got != expected {
		t.Errorf("Expected:\n%s, got:\n%s", expected, got)
	}

	table = Table{
		TableRow{"1 Tyler", "4 mi", ""},
		TableRow{"2 Jim", "", "12:03"},
		TableRow{"3 Jonathan", "14.7 mi", "2:03"},
		TableRow{"4 K", "", ""},
	}
	expected = "1 Tyler         4 mi        \n2 Jim                  12:03\n3 Jonathan   14.7 mi    2:03\n4 K                         \n"
	got = table.composeRightAlignedTable(2, false)
	if got != expected {
		t.Errorf("Expected:\n%s, got:\n%s", expected, got)
	}
}
