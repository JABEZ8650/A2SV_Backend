package main

import "testing"

func averageCalculaterTest(t *testing.T) {
	grades := map[string]float64{
		"math": 75,
		"eng":  65,
		"phy":  70,
	}

	got := average(grades)
	want := 70.0

	if got != want {
		t.Errorf("Expected value is %v, but got %v", want, got)
	}
}
