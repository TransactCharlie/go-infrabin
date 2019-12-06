package main

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	expected := "120"
	os.Setenv("MAX_DELAY", expected)
	value := GetEnv("MAX_DELAY", "12")
	if value != expected {
		t.Errorf("GetEnv returned unexpected value: got %v want %v",
			value, expected)
	}

	defaultValue := "42"
	value = GetEnv("MISSING", defaultValue)
	if value != defaultValue {
		t.Errorf("GetEnv returned unexpected value: got %v want %v",
			value, defaultValue)
	}
}

func TestMin(t *testing.T) {
	var x, y int = 5, 3
	min := Min(x, y)
	if min != y {
		t.Errorf("Min returned unexpected value: got %v want %v",
			min, y)
	}
}
