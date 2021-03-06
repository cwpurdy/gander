package csv

import (
	"testing"
)

func TestGetShape(t *testing.T) {
	got, err := GetShape("../../majestic_million.csv", false)

	expected := Shape{1000000, 12}

	if err != nil {
		t.Error(err)
	}
	if *got != expected {
		t.Errorf("Did not get expected results. Got: '%v' Expected '%v'", got, expected)
	}

	// Test for bad path to CSV
	_, err = GetShape("some/bogus/path.csv", false)

	if err == nil {
		t.Error("Function should return error for bad path -> some/bogus/path.csv" )
	}
}

func BenchmarkGetShape(b *testing.B) {
	GetShape("../../majestic_million.csv", false)
}