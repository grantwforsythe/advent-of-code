package main

import (
	"testing"
)

func TestSolvePart1(t *testing.T) {
	rows := parseInput("input_test.txt")

	if len(rows) != 10 && len(rows[0]) != 10 {
		t.Fatalf("rows does not have the right dimensions")
	}

	result := solvePart1(rows)
	if result != 18 {
		t.Errorf("got %d; expected 18", result)
	}
}
