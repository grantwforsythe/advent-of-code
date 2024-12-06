// Day 04 is a search problem
package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	rows := parseInput("input.txt")
	log.Printf("Part 1: %d", solvePart1(rows))
}

func parseInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Unable to open input.txt")
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	// PERF: Memory inefficient as the slice constantly needs to reallocate memory
	rows := make([]string, 0)

	for {
		row, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Failed to read line from input.txt")
		}

		rows = append(rows, row)
	}

	return rows
}

func solvePart1(rows []string) int {
	// Looking at the data we can discern that it is symmetrical and is 140x140
	// The upperBound is 140 - 1 because arrays are zero-indexed
	upperBound := len(rows) - 1

	count := 0
	for colPos := range len(rows) {
		for rowPos := range len(rows[colPos]) {
			// Pivot on X
			if rows[colPos][rowPos] != 'X' {
				continue
			}

			// NOTE: The first value in a slice index is inclusive whereas the last value is exclusive

			// Check that not out of bounds and walk right
			if rowPos+4 <= upperBound && rows[colPos][rowPos+1:rowPos+4] == "MAS" {
				count += 1
			}

			// Check that not out of bounds and walk left
			if rowPos-3 >= 0 && rows[colPos][rowPos-3:rowPos] == "SAM" {
				count += 1
			}

			// Check that not out of bounds and walk down
			if colPos+3 <= upperBound && (rows[colPos+1][rowPos] == 'M' &&
				rows[colPos+2][rowPos] == 'A' &&
				rows[colPos+3][rowPos] == 'S') {
				count += 1
			}

			// Check that not out of bounds and walk up
			if colPos-3 >= 0 && (rows[colPos-1][rowPos] == 'M' &&
				rows[colPos-2][rowPos] == 'A' &&
				rows[colPos-3][rowPos] == 'S') {
				count += 1
			}

			// Check that not out of bounds and walk north east
			if colPos-3 >= 0 && rowPos+3 <= upperBound && (rows[colPos-1][rowPos+1] == 'M' &&
				rows[colPos-2][rowPos+2] == 'A' &&
				rows[colPos-3][rowPos+3] == 'S') {
				count += 1
			}

			// Check that not out of bounds and walk south east
			if colPos+3 <= upperBound && rowPos+3 <= upperBound &&
				(rows[colPos+1][rowPos+1] == 'M' &&
					rows[colPos+2][rowPos+2] == 'A' &&
					rows[colPos+3][rowPos+3] == 'S') {
				count += 1
			}

			// Check that not out of bounds and walk south west
			if colPos+3 <= upperBound && rowPos-3 >= 0 &&
				(rows[colPos+1][rowPos-1] == 'M' &&
					rows[colPos+2][rowPos-2] == 'A' &&
					rows[colPos+3][rowPos-3] == 'S') {
				count += 1
			}

			// Check that not out of bounds and walk north west
			if colPos-3 >= 0 && rowPos-3 >= 0 && (rows[colPos-1][rowPos-1] == 'M' &&
				rows[colPos-2][rowPos-2] == 'A' &&
				rows[colPos-3][rowPos-3] == 'S') {
				count += 1
			}
		}
	}

	return count
}
