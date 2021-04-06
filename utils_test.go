package main

import (
	"fmt"
	"testing"
)

func TestMetersToMiles(t *testing.T) {
	got := metersToMiles(1000)
	if got != 0.621371 {
		t.Errorf("Converting 1000m to miles; got %q but wanted 0.621371", fmt.Sprintf("%f", got))
	}
}
