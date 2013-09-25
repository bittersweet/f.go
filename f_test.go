package main

import (
	"testing"
)

func TestColorizeMatch(t *testing.T) {
	input := "test"
	output := "\033[31mtest\033[0m"
	result := colorizeMatch(input)
	if input == output {
		t.Fatalf("Failed, got %v, expected %v", result, output)
	}
}
