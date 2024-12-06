package main

import (
	"testing"
)

type test struct {
	filename string
	expected int
}

func TestSolvePart1(t *testing.T) {
	tests := []test{
		{"input_test.txt", 18},
		{"test1.txt", 9},
		{"test2.txt", 1},
		{"test3.txt", 4},
	}

	for _, test := range tests {
		rows := parseInput(test.filename)

		result := solvePart1(rows)
		if result != test.expected {
			t.Errorf("got %d; expected %d", result, test.expected)
		}
	}

}
